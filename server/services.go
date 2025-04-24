package main

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation int8

// const (
// 	OperationAdd    Operation = 0
// 	OperationUpdate Operation = 1
// 	OperationDelete Operation = 2
// )

// type Log struct {
// 	Operation Operation
// 	Log       string
// 	NewValue  string
// 	OldValue  string
// 	date      time.Time
// 	// User            int32
// }

type Log struct {
	Date string `json:"date"`
	Text string `json:"text"`
}

type status int8

const (
	pending status = iota
	due
	overdue // vencido
	paid
)

func (s status) String() string {
	switch s {
	case pending:
		return "pending"
	case due:
		return "due"
	case overdue:
		return "overdue"
	case paid:
		return "paid"
	}
	return "unknown"
}

type Rent struct {
	Id     int64 `json:"id"`
	active bool
	// Done   bool   `json:"done"`   // remove
	// Status status `json:"status"` // remove
	Date string `json:"date"`
	// Dt     string `json:"dt"`
	Name string `json:"name"`
}

type RentReminder struct {
	Id int64 `json:"id"`
	// Rent     Rent   `json:"rent"`
	// PrevReminders []RentReminder `json:"prev_reminders"`
	RentId   int64  `json:"rent_id,omitempty"`
	RentName string `json:"rent_name,omitempty"`
	Date     string `json:"date"`
	Status   status `json:"status"`
}

// REMOVE
type Tenant struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Active    bool               `json:"active"`
	Name      string             `json:"name"`
	Cpf       string             `json:"cpf"`
	Rg        string             `json:"rg"`
	BirthDate int64              `json:"birth_date" bson:"birth_date"`
	RentId    primitive.ObjectID `json:"rent_id" bson:"rent_id"`
	// PropertyId primitive.ObjectID `json:"property_id" bson:"property_id"`
}

type Property struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Active  bool               `json:"active"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	RentId  primitive.ObjectID `json:"rent_id" bson:"rent_id"`
	Tenant  interface{}        `json:"tenant" bson:"tenant"`
	// TenantId primitive.ObjectID `json:"tenant_id" bson:"tenant_id"`
}

// type Rent struct {
// 	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// 	Active     bool               `json:"active"`
// 	TenantId   primitive.ObjectID `json:"tenant_id"`
// 	PropertyId primitive.ObjectID `json:"property_id"`
// }

type Config struct {
	documentType string
	collection   string
	dataBase     string
}

type DocsConfig struct {
	dataBase string
	configs  []Config
}

// Returns database and collection name of each document type
func (c *DocsConfig) getDocConfig(doc interface{}) (Config, error) {

	getTypeConfig := func(typeName string) (Config, error) {
		for _, c := range c.configs {
			if c.documentType == typeName {
				return c, nil
			}
		}
		return Config{}, errors.New("Tipo de documento não definidoAAAAA")
	}

	config := Config{}
	var err error = nil

	switch doc.(type) {
	case Tenant, *Tenant, *[]Tenant:
		config, err = getTypeConfig("Tenant")
	case Property, *Property, *[]Property:
		config, err = getTypeConfig("Property")
	// case Rent, *Rent, *[]Rent:
	// 	config, err = getTypeConfig("Rent")
	default:
		err = errors.New("Tipo de documento não definido")
	}

	config.dataBase = c.dataBase

	return config, err
}

func (c *DocsConfig) addConfig(newConfig Config) {
	c.configs = append(c.configs, newConfig)
}

var configs DocsConfig = DocsConfig{
	dataBase: "srv1140",
	configs: []Config{
		{
			documentType: "Tenant",
			collection:   "tenants",
		},
		{
			documentType: "Property",
			collection:   "properties",
		},
		{
			documentType: "Rent",
			collection:   "rents",
		},
	},
}

func isTenantValid(tenant Tenant) (bool, string) {
	if isValid, msg := isValidCpf(tenant.Cpf); !isValid {
		return false, msg
	}
	// TODO
	return true, ""
}

func dbExec(query string, values ...any) (sql.Result, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer stmt.Close()
	return stmt.Exec(values...)
}

func clearDB() error {
	if _, err := db.Exec("DROP TABLE cue"); err != nil {
		return errors.Wrap(err, "cue table not created")
	}
	return nil
}

func initDb() error {
	qry := `
		CREATE TABLE IF NOT EXISTS rent (
			id 		INTEGER
					NOT NULL
					PRIMARY KEY AUTOINCREMENT,
			active 	BOOLEAN
					DEFAULT TRUE,
			done 	BOOLEAN 
					DEFAULT FALSE,
			--status 	INTEGER
			--		NOT NULL
			--		DEFAULT 0
			--		CHECK(status >= 0 AND status <= 3), -- 0=pending, 1=due, 2=overdue, 3=paid
			name 	TEXT
					NOT NULL,
			date 	TEXT
					DEFAULT (DATE('0', 'unixepoch'))
					CHECK(length(date) == 10),
			created_at NUMERIC
					DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now'))
		)
	`
	// date NUMERIC DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime'))
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "rent table not created")
	}

	qry = `
		CREATE TABLE IF NOT EXISTS reminder (
			id		INTEGER
					NOT NULL
					PRIMARY KEY 
					AUTOINCREMENT,
			active 	BOOLEAN
					DEFAULT TRUE,
			status 	INTEGER
					NOT NULL
					DEFAULT 0
					CHECK(status >= 0 AND status <= 3), -- 0=pending, 1=due, 2=overdue, 3=paid
			date 	TEXT
					DEFAULT (DATE('0', 'unixepoch'))
					CHECK(length(date) == 10),
			id_rent
					INTEGER
					NOT NULL,
			FOREIGN KEY(id_rent) REFERENCES rent(id)
		)
	`
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "reminder table not created")
	}

	// qry = `
	// 	CREATE TABLE IF NOT EXISTS log (
	// 		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	// 		date NUMERIC,
	// 		log TEXT
	// 	)`
	qry = `
		CREATE TABLE IF NOT EXISTS rent_log (
			id		INTEGER 
					NOT NULL 
					PRIMARY KEY
					AUTOINCREMENT,
			-- date 	TEXT
			--		DEFAULT (DATE('0', 'unixepoch'))
			--		CHECK(length(date) == 10),
			date	NUMERIC
					DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now')),
			log 	TEXT,
			id_rent INTEGER
					NOT NULL,
			FOREIGN KEY(id_rent) REFERENCES rent(id)
			-- table	TEXT,
			-- rowId 	INTEGER,
		)
	`
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "rent_log not created")
	}

	qry = `
		CREATE TABLE IF NOT EXISTS scheduler (
			id 			INTEGER 
						NOT NULL 
						PRIMARY KEY 
						AUTOINCREMENT,
			start_exec 	DATETIME,
			end_exec 	DATETIME,
			next_exec 	DATETIME
		)
	`
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "scheduler not created")
	}
	// qry = `
	// 	CREATE TABLE IF NOT EXISTS log_cue_register (
	// 		id integer not null primary key AUTOINCREMENT,
	// 		row_id integer not null,
	// 		old_active BOOLEAN,
	// 		new_active BOOLEAN,
	// 		old_status INTEGER,
	// 		new_status INTEGER,
	// 		old_name TEXT not null,
	// 		new_name TEXT not null,
	// 		old_date TEXT not null,
	// 		new_date TEXT not null,
	// 		change_type TEXT not null,
	// 		created_at NUMERIC not null
	// 	);
	// `
	// if _, err := db.Exec(qry); err != nil {
	// 	return errors.Wrap(err, "trigger not created")
	// }
	// qry = `
	// 	CREATE TRIGGER IF NOT EXISTS log_cue_register_after_update
	// 	AFTER UPDATE ON rent
	// 		WHEN old.name <> new.name
	// 			OR old.active <> new.active
	// 			OR old.date <> new.date
	// 	BEGIN
	// 		insert into log_cue_register (
	// 			row_id,
	// 			old_active,
	// 			new_active,
	// 			old_name,
	// 			new_name,
	// 			old_date,
	// 			new_date,
	// 			change_type,
	// 			created_at
	// 		)
	// 		values (
	// 			old.id,
	// 			old.active,
	// 			new.active,
	// 			old.name,
	// 			new.name,
	// 			old.date,
	// 			new.date,
	// 			'UPDATE',
	// 			strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
	// 		);
	// 	END;
	// `
	// if _, err := db.Exec(qry); err != nil {
	// 	return errors.Wrap(err, "tr4igger not created")
	// }
	return nil
}

func populateDb() error {
	// stmt, err := db.Prepare("INSERT INTO cue (active, name, dt) VALUES (?, ?, ?)")
	stmt, err := db.Prepare("INSERT INTO rent (active, name, date) VALUES (?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer stmt.Close()
	dt := time.Now()
	// dt = dt.AddDate(0, -1, 0)
	dt = time.Date(dt.Year(), dt.Month(), 11, 0, 0, 0, 0, time.UTC)
	// if _, err = stmt.Exec(true, "a1a1a1", time.Now().Add(time.Duration(rand.IntN(10))*time.Minute).Format(sqliteLayout)); err != nil {
	if _, err = stmt.Exec(true, "a1a1a1", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "")
	}
	dt = time.Date(dt.Year(), dt.Month(), 7, 0, 0, 0, 0, time.UTC)
	// if _, err = stmt.Exec(true, "a2a22aa2", time.Now().Add(time.Duration(rand.IntN(10))*time.Minute).Format(sqliteLayout)); err != nil {
	if _, err = stmt.Exec(true, "a2a22aa2", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "")
	}
	// dt = dt.AddDate(0, -1, 0)
	dt = time.Date(dt.Year(), dt.Month(), 10, 0, 0, 0, 0, time.UTC)
	// if _, err = stmt.Exec(true, "bb34b4b234bb234", time.Now().Add(time.Duration(rand.IntN(10))*time.Minute).Format(sqliteLayout)); err != nil {
	if _, err = stmt.Exec(true, "bb34b4b234bb234", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "")
	}
	dt = time.Date(dt.Year(), dt.Month(), 2, 0, 0, 0, 0, time.UTC)
	if _, err = stmt.Exec(true, "sadjlhasklhasjkldh", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "")
	}
	// dt = dt.AddDate(0, -5, 0)
	dt = time.Date(dt.Year(), dt.Month(), rand.IntN(20), 0, 0, 0, 0, time.UTC)
	if _, err = stmt.Exec(true, "ÇASDçasçdasçfkasjfas", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "")
	}

	dt = dt.Add(time.Duration(rand.IntN(10)) * time.Minute)
	// time.Now().Add(time.Duration(rand.IntN(10)) * time.Minute)
	if _, err := db.Exec("insert into rent (name, date) values ('TESssssTESTE', ?)", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "trigger not created")
	}
	// dt = dt.AddDate(0, 0, 1)
	dt = time.Date(dt.Year(), dt.Month(), rand.IntN(20), 0, 0, 0, 0, time.UTC)
	if _, err := db.Exec("insert into rent (name, date) values ('TESTESTE', ?)", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "trigger not created")
	}

	return nil
}

func saveRentLog(rentId int64, msg string) {
	qry := "INSERT INTO rent_log(id_rent, log) VALUES (?, ?);"
	if _, err := db.Exec(qry, rentId, msg); err != nil {
		log.Error().Stack().Err(err).Msg("log error")
	}
}

func createRent(rent Rent) error {
	log.Debug().Interface("rent", rent).Msg("createRent()")
	parsedDate, err := parseIsoDateTime(rent.Date)
	if err != nil {
		return errors.Wrap(err, "")
	}
	qry := "INSERT INTO rent (active, name, date) VALUES (TRUE, ?, ?)"
	result, err := db.Exec(qry, rent.Name, parsedDate.Format(time.DateOnly))
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("createRent()")
	rentId, err := result.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.Info().Msgf("new rent. id: %d", rentId)

	saveRentLog(rentId, fmt.Sprintf("created: %v - %v", rent.Name, parsedDate.Format(time.DateOnly)))
	err = createReminder(rentId)
	if err != nil {
		log.Warn().
			Err(err).
			Msg("createRent() > createReminder()")
	}
	err = processRemindersDates(rentId)
	if err != nil {
		log.Warn().
			Err(err).
			Msg("createRent() > processRemindersDates()")
	}
	return nil
}

func listRent() ([]Rent, error) {
	rents := []Rent{}
	rows, err := db.Query("SELECT id, active, name, date FROM rent WHERE active = TRUE;")
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("listRent()")
	for rows.Next() {
		rent := Rent{}
		err = rows.Scan(&(rent.Id), &(rent.active), &(rent.Name), &(rent.Date))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		rents = append(rents, rent)
	}
	return rents, nil
}

func parseIsoDateTime(dt interface{}) (time.Time, error) {
	v, ok := dt.(string)
	if !ok {
		return time.Time{}, errors.New("datetime is not a string")
	}
	if len(v) < len(time.RFC3339) {
		log.Warn().
			Str("func", "parseIsoDateTime").
			Str("value", v).
			Msg("value length is smaller than the one expected by the format time.RFC3339")
	}
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString
	// expected format: YYYY-MM-DDTHH:mm:ss.sssZ from javascript new Date().toISOString()
	dtParsed, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return time.Time{}, err
	}
	return dtParsed, nil
}

func updateRent(rentId int64, newValues map[string]interface{}) error {
	var queryColumns []string
	var queryValues []any
	var allowedFields []string = []string{"status", "name", "date", "dt"}

	// Track if date is being updated
	var dateChanged bool
	var oldDateStr string
	var newDateStr string

	// detect fields to be used in where clause
	for _, field := range allowedFields {
		if value, ok := newValues[field]; ok {
			queryColumns = append(queryColumns, field+" = ?")
			if field == "date" {
				dateChanged = true
				parsedDate, err := parseIsoDateTime(value)
				if err != nil {
					return errors.Wrap(err, "parseIsoDateTime()")
				}
				newDateStr = parsedDate.Format(time.DateOnly)
				queryValues = append(queryValues, newDateStr)

				// Get the current date before updating
				row := db.QueryRow("SELECT date FROM rent WHERE id = ?", rentId)
				if err := row.Scan(&oldDateStr); err != nil {
					return errors.Wrap(err, "failed to get current rent date")
				}
			} else if field == "dt" {
				// TODO remove field dt
				dtParsed, _ := parseIsoDateTime(value)
				queryValues = append(queryValues, dtParsed.Format(sqliteLayout))
			} else {
				queryValues = append(queryValues, value)
			}
		}
	}

	var logValues []string
	for _, v := range queryValues {
		logValues = append(logValues, v.(string))
	}
	queryValues = append(queryValues, rentId)

	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "db.Begin()")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("updateRent()")

	qry := `
        UPDATE rent 
        SET ` + strings.Join(queryColumns, ", ") + `
        WHERE id = ? 
          AND active = TRUE;
    `
	log.Debug().Stack().Str("query_prepared", qry).Msg("")
	log.Debug().Interface("query_values", queryValues).Msg("")

	stmt, err := tx.Prepare(qry)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer stmt.Close()

	res, err := stmt.Exec(queryValues...)
	if err != nil {
		return errors.Wrap(err, "")
	}

	count, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "")
	}

	if count != 1 {
		log.Info().Msgf("rowsAffected: %v", count)
		return errors.New("an error occurred while updating")
	}

	// If date was changed, update the related reminder
	if dateChanged {
		// Get current month's reminder
		firstDay, _ := getFirstLastDayOfMonth(time.Now())
		currentMonthReminderDate := firstDay.AddDate(0, 0, getDayFromDateString(oldDateStr)-1)

		// Update the reminder's date if it exists
		updateReminderQry := `
            UPDATE reminder
            SET date = ?
            WHERE id_rent = ?
              AND date = ?
              AND active = TRUE
        `

		newDay := getDayFromDateString(newDateStr)
		newReminderDate := firstDay.AddDate(0, 0, newDay-1)

		_, err = tx.Exec(updateReminderQry,
			newReminderDate.Format(time.DateOnly),
			rentId,
			currentMonthReminderDate.Format(time.DateOnly),
		)
		if err != nil {
			return errors.Wrap(err, "failed to update reminder date")
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "tx.Commit()")
	}

	saveRentLog(rentId, fmt.Sprintf("updated: %v", strings.Join(logValues, " - ")))

	// Re-process reminder to update its status
	if err := processRemindersDates(rentId); err != nil {
		log.Warn().Err(err).Msg("failed to process reminders after date update")
	}
	return nil
}

// Helper function to extract day from date string (YYYY-MM-DD)
func getDayFromDateString(dateStr string) int {
	parts := strings.Split(dateStr, "-")
	if len(parts) < 3 {
		return 0
	}
	day, _ := strconv.Atoi(parts[2])
	return day
}

func removeRent(rentId int64) error {
	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("removeRent()")

	qry := `
		UPDATE rent 
		SET active = FALSE
		WHERE id = ?;
	`
	res, err := db.Exec(qry, rentId)
	if err != nil {
		return errors.Wrap(err, "")
	}
	count, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "")
	}
	if count != 1 {
		log.Info().Msgf("rowsAffected: %d - rent not updated.", count)
	} else {
		log.Info().Msgf("rowsAffected: %d - rent removed. id: %d", count, rentId)
		saveRentLog(rentId, "deleted")
	}

	qry = `
		UPDATE reminder
		SET active = FALSE
		WHERE id_rent = ?
		  AND date >= ?;
	`
	firstDayOfMonth, _ := getFirstLastDayOfMonth(time.Now())
	res, err = db.Exec(qry, rentId, firstDayOfMonth.Format(time.DateOnly))
	if err != nil {
		return errors.Wrap(err, "")
	}
	count, err = res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "")
	}
	if count != 1 {
		log.Info().Msgf("rowsAffected: %d - reminder not updated.", count)
	} else {
		log.Info().Msgf("rowsAffected: %d - reminder removed. id_rent: %d", count, rentId)
	}

	return nil
}

func listRentHystory(rentId int64) ([]Log, error) {
	logs := []Log{}
	rows, err := db.Query("SELECT date, log FROM rent_log WHERE id_rent = ?;", rentId)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("listRentHystory()")
	for rows.Next() {
		rentLog := Log{}
		err = rows.Scan(&(rentLog.Date), &(rentLog.Text))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		logs = append(logs, rentLog)
	}
	return logs, nil
}

func getFirstLastDayOfMonth(date time.Time) (firstDay time.Time, lastDay time.Time) {
	firstDay = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay = firstDay.AddDate(0, 1, -1)
	return firstDay, lastDay
}

func createReminder(rentId int64) error {
	log.Debug().
		Int64("rent id", rentId).
		Msg("createReminder()")
	qry := `
		SELECT id, active, name, date 
		FROM rent 
	`
	if rentId > 0 {
		qry += " WHERE id = ?"
	}
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer tx.Commit()
	rows, err := tx.Query(qry, rentId)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer rows.Close()
	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("createReminder()")

	qry = `
		INSERT INTO reminder (id_rent, date)
		SELECT :id, :date
		WHERE NOT EXISTS (
			SELECT 1 FROM reminder
			WHERE id_rent = :id AND date = :date
		);
	`
	stmt, err := tx.Prepare(qry)
	if err != nil {
		log.Error().Stack().Err(err).Msg("createReminder() > at tx.Prepare()")
	}
	defer stmt.Close()

	rent := Rent{}
	for rows.Next() {
		err = rows.Scan(&(rent.Id), &(rent.active), &(rent.Name), &(rent.Date))
		if err != nil {
			return errors.Wrap(err, "")
		}
		parsedDate, err := time.Parse(time.DateOnly, rent.Date)
		if err != nil {
			return errors.Wrap(err, "")
		}
		dueDay := parsedDate.Day()
		firstDayOfMonth, lastDayOfMonth := getFirstLastDayOfMonth(time.Now())
		if dueDay > lastDayOfMonth.Day() {
			dueDay = lastDayOfMonth.Day()
		}
		rentDue := firstDayOfMonth.AddDate(0, 0, dueDay-1)
		log.Debug().
			Str("rentDue", rentDue.Format(time.DateOnly)).
			Str("firstDayOfMonth", firstDayOfMonth.Format(time.DateOnly)).
			Str("lastDayOfMonth", lastDayOfMonth.Format(time.DateOnly)).
			Msg("")

		res, err := stmt.Exec(rent.Id, rentDue.Format(time.DateOnly))
		if err != nil {
			log.Error().Stack().Err(err).Msg("createReminder() > at stmt.Exec()")
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Error().Stack().Err(err).Msg("createReminder() > at res.RowsAffected()")
		}
		if rowsAffected == 0 {
			log.Info().Msgf("cue not created")
		} else {
			lastInsertId, err := res.LastInsertId()
			if err != nil {
				log.Error().Stack().Err(err).Msg("createReminder() > at res.LastInsertId()")
			}
			log.Info().Msgf("rowsAffected: %d - new reminder id: %d", rowsAffected, lastInsertId)
		}
	}
	return nil
}

func listReminderDetail(reminderId int64) ([]RentReminder, error) {
	qry := `
		SELECT
			date,
			id_rent
		FROM reminder
		WHERE id = ?
		LIMIT 1
	`
	row, err := db.Query(qry, reminderId)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer row.Close()
	reminderBase := RentReminder{}
	for row.Next() {
		err = row.Scan(&reminderBase.Date, &reminderBase.RentId)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
	}

	// log.Info().Str("reminder_date", dateStr).Msg("")
	reminderBaseDate, err := time.Parse(time.DateOnly, reminderBase.Date)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// log.Info().Str("reminder_date", reminderBaseDate.String()).Msg("")

	// get last 3 months
	start, end := getFirstLastDayOfMonth(reminderBaseDate.AddDate(0, -1, 0))
	start = start.AddDate(0, -2, 0)
	// qryStart := start.Format(time.DateOnly)
	// qryEnd := end.Format(time.DateOnly)
	qry = `
		SELECT
			reminder.id,
			reminder.status,
			reminder.date
		FROM reminder 
		JOIN rent 
			ON rent.id = reminder.id_rent
		WHERE 
			-- reminder.id = :id
			rent.id = :rent_id
			AND reminder.date BETWEEN :date_start
			AND :date_end
			-- AND reminder.active
	`
	log.Debug().
		Str("start", start.Format(time.DateOnly)).
		Str("end", end.Format(time.DateOnly)).
		Msg("query params")

	rows, err := db.Query(
		qry,
		reminderBase.RentId,
		start.Format(time.DateOnly),
		end.Format(time.DateOnly),
	)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("listReminders()")

	var reminders []RentReminder
	for rows.Next() {
		reminder := RentReminder{}
		// prev_reminders := []RentReminder{}
		// reminder.PrevReminders = prev_reminders
		err = rows.Scan(
			&(reminder.Id),
			&(reminder.Status),
			&(reminder.Date),
		)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		reminders = append(reminders, reminder)
	}
	return reminders, nil
}

func listReminders(rentId int64, monthStart time.Time, monthEnd time.Time) ([]RentReminder, error) {
	monthStart, _ = getFirstLastDayOfMonth(monthStart)
	_, monthEnd = getFirstLastDayOfMonth(monthEnd)
	qryStart := monthStart.Format(time.DateOnly)
	qryEnd := monthEnd.Format(time.DateOnly)

	qry := `
		SELECT
			reminder.id,
			reminder.status,
			reminder.date,
			rent.id,
			rent.name
		FROM reminder 
		JOIN rent 
			ON rent.id = reminder.id_rent
			-- AND rent.id = 5
		WHERE 
			reminder.date BETWEEN ?
			AND ?
			AND reminder.active
	`
	if rentId > 0 {
		qry += "AND rent.id = ? "
	}

	log.Debug().
		Str("start", qryStart).
		Str("end", qryEnd).
		Msg("query params")

	rows, err := db.Query(qry, qryStart, qryEnd, rentId)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("listReminders()")

	var reminders []RentReminder
	for rows.Next() {
		reminder := RentReminder{}
		err = rows.Scan(
			&(reminder.Id),
			&(reminder.Status),
			&(reminder.Date),
			&(reminder.RentId),
			&(reminder.RentName),
		)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		reminders = append(reminders, reminder)
	}
	return reminders, nil
}

func changeReminderStatus(id int64, s status) error {
	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("changeReminderStatus()")

	qry := `
		UPDATE reminder 
		SET status = :status 
		WHERE id = :id 
		  AND status <> :status ;
	`
	if _, err := db.Exec(qry, s, id); err != nil {
		return errors.Wrap(err, "update not concluded. rollback executed")
	}

	log.Info().
		Int64("reminder_id", id).
		// Int64("count", count).
		Str("new_status", s.String()).
		Msg("cue status changed.")

	if s == status(paid) {
		qry = `
			SELECT id_rent 
			FROM reminder 
			WHERE id = ?;
		`
		rows, err := db.Query(qry, id)
		if err != nil {
			log.Error().Stack().Err(err).Msg("")
		}
		rows.Next()
		var rentId int64
		if err = rows.Scan(&rentId); err != nil {
			log.Error().Stack().Err(err).Msg("")
		}
		rows.Close()
		saveRentLog(rentId, "paid")
	}
	return nil
}

func processRemindersDates(rentId int64) error {
	// reminders, err := listReminders(time.Now(), time.Now())
	reminders, err := listReminders(rentId, time.Now().AddDate(0, -10, 0), time.Now())
	// rents, err := listRent()
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.Debug().
		Interface("len_reminders", len(reminders)).
		Msg("processRemindersDates")
	for _, reminder := range reminders {
		rentDate, err := time.Parse(time.DateOnly, reminder.Date)
		if err != nil {
			return errors.Wrap(err, "")
		}
		now := time.Now()
		now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		// 0=pending, 1=due, 2=overdue, 3=paid
		// "info", "warn", "error", "success"
		log.Debug().
			Int64("rent_id", reminder.RentId).
			Int64("id", reminder.Id).
			Interface("rent_status", reminder.Status).
			Str("name", reminder.RentName).
			Str("date", reminder.Date).
			Msg("processRemindersDates")
		if rentDate.Equal(now) && reminder.Status == pending && reminder.Status != due {
			if err = changeReminderStatus(reminder.Id, due); err != nil {
				return errors.Wrap(err, "")
			}
		} else if rentDate.Before(now) && reminder.Status != paid && reminder.Status != overdue {
			if err = changeReminderStatus(reminder.Id, overdue); err != nil {
				return errors.Wrap(err, "")
			}
		} else if rentDate.After(now) && reminder.Status != paid && reminder.Status != pending {
			if err = changeReminderStatus(reminder.Id, pending); err != nil {
				return errors.Wrap(err, "")
			}
		}
	}
	return nil
}

type schedulerHistory struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func listSchedulerHistory() ([]schedulerHistory, error) {
	history := []schedulerHistory{}

	rows, err := db.Query("SELECT start_exec, end_exec FROM scheduler;")
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open_connections", db.Stats().OpenConnections).
		Msg("listSchedulerHistory()")
	for rows.Next() {
		var arg1, arg2 any
		if err = rows.Scan(&arg1, &arg2); err != nil {
			return nil, errors.Wrap(err, "")
		}

		start := ""
		end := ""
		if t1, ok := arg1.(time.Time); ok {
			start = t1.Format(time.RFC3339)
		}
		if t2, ok := arg2.(time.Time); ok {
			end = t2.Format(time.RFC3339)
		}
		history = append(history, schedulerHistory{start, end})
	}

	return history, nil
}

// var cuesStats = make(map[int]bool)

func checkDiff() {
	log.Info().Msg("checkDiff started")
	cues, err := listRent()
	if err != nil {
		log.Error().Stack().Err(err).Msg("checkDiff")
		return
	}
	// if len(cuesStats) == 0 {
	// 	for _, cue := range cues {
	// 		cuesStats[int(cue.Id)] = cue.Done
	// 	}
	// } else {
	for _, cue := range cues {
		// if cue.Done {
		newStat := make(map[string]interface{})
		newStat["status"] = 1
		err := updateRent(cue.Id, newStat)
		if err != nil {
			log.Error().Stack().Err(err).Msg("")
		} else {
			log.Debug().Int("id", int(cue.Id)).Msg("uptated")
		}
		// }
	}
	// }
	// log.Info().Interface("cues", cuesStats).Msg("")
}

// var s1 scheduler = scheduler{}
// func sc() []string {
// 	r := []string{}
// 	r = append(r, fmt.Sprintf("time %v", time.Now().Format(sqliteLayout)))
// 	r = append(r, fmt.Sprintf("running %v", s1.running))
// 	r = append(r, fmt.Sprintf("ticker %v", s1.ticker))
// 	// r = append(r, fmt.Sprintf("ticker.C %v", s1.ticker.C))
// 	// if s1.ticker.C != nil {
// 	// 	r = append(r, fmt.Sprintf("ticker false %v", s1.ticker.C))
// 	// }
// 	return r
// }

// func work(t any) error {
// 	fmt.Println("Tick at", t)
// 	return errors.New("asdasdasd")
// }
// func start() {
// 	s1.start(work)
// 	// s1.ticker = time.NewTicker(1 * time.Second)
// 	// s1.done = make(chan bool)
// 	// go func() {
// 	// 	for {
// 	// 		select {
// 	// 		case <-s1.done:
// 	// 			return
// 	// 		case t := <-s1.ticker.C:
// 	// 			work(t)
// 	// 		}
// 	// 	}
// 	// }()
// 	// time.Sleep(5 * time.Second)
// 	// s1.ticker.Stop()
// 	// s1.done <- true
// 	fmt.Println("Ticker stopped")

// }
//
//	func stop() {
//		s1.stop()
//		// s1.ticker.Stop()
//		// // s1.done <- true
//	}
var lastSchedulerExecution time.Time

func executeScheduler(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				{
					log.Info().Msg("scheduler start")

					qry := `INSERT INTO scheduler (start_exec) VALUES (?);`
					result, err := dbExec(qry, time.Now().Format(sqliteLayout))
					if err != nil {
						log.Error().Stack().Err(err).Msg("")
					}
					scheduler_id, err := result.LastInsertId()
					if err != nil {
						log.Error().Stack().Err(err).Msg("")
					}

					if lastSchedulerExecution.IsZero() {
						lastSchedulerExecution = time.Now()
					}
					if time.Now().Month() != lastSchedulerExecution.Month() {
						lastSchedulerExecution = time.Now()
						createReminder(0)
					}

					processRemindersDates(0)

					qry = `UPDATE scheduler SET end_exec = ? WHERE id = ?;`
					_, err = dbExec(qry, time.Now().Format(sqliteLayout), scheduler_id)
					if err != nil {
						log.Error().Stack().Err(err).Msg("")
					}
					log.Info().Msg("scheduler end")
				}
			}
		}
	}()
}
