package main

import (
	"database/sql"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation int8

const (
	OperationAdd    Operation = 0
	OperationUpdate Operation = 1
	OperationDelete Operation = 2
)

type Log struct {
	Operation Operation
	Log       string
	NewValue  string
	OldValue  string
	date      time.Time
	// User            int32
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
	case overdue:
		return "overdue"
	case paid:
		return "paid"
	}
	return "unknown"
}

type Cue struct {
	Id     int64 `json:"id"`
	active bool
	Done   bool   `json:"done"`
	Status status `json:"status"`
	Date   string `json:"date"`
	Dt     string `json:"dt"`
	Name   string `json:"name"`
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

type Rent struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Active     bool               `json:"active"`
	TenantId   primitive.ObjectID `json:"tenant_id"`
	PropertyId primitive.ObjectID `json:"property_id"`
}

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
	case Rent, *Rent, *[]Rent:
		config, err = getTypeConfig("Rent")
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

func initDB() error {
	qry := `
		CREATE TABLE IF NOT EXISTS cue (
			id 		INTEGER
					NOT NULL
					PRIMARY KEY AUTOINCREMENT,

			active 	BOOLEAN
					DEFAULT TRUE,

			done 	BOOLEAN 
					DEFAULT FALSE,

			status 	INTEGER
					NOT NULL
					DEFAULT 0
					CHECK(status >= 0 AND status <= 3), -- 0=pending, 1=due, 2=overdue, 3=paid

			name 	TEXT
					NOT NULL,

			date 	TEXT
					DEFAULT (DATE('0', 'unixepoch'))
					CHECK(length(date) == 10),

			created_at NUMERIC
					DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now'))
		)`
	// date NUMERIC DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime'))
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "cue table not created")
	}

	qry = `
		CREATE TABLE IF NOT EXISTS log (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			date NUMERIC,
			log TEXT
		)`
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "log not created")
	}

	qry = `
	CREATE TABLE IF NOT EXISTS log_cue (
		id integer not null primary key AUTOINCREMENT,
		row_id integer not null,
		old_active BOOLEAN,
		new_active BOOLEAN,
		old_status INTEGER,
		new_status INTEGER,
		old_name TEXT not null,
		new_name TEXT not null,
		old_date TEXT not null,
		new_date TEXT not null,
		change_type TEXT not null,
		created_at NUMERIC not null
	);
	`
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "log_cue not created")
	}
	qry = `
	CREATE TRIGGER IF NOT EXISTS log_cue_after_update 
	AFTER UPDATE ON cue
		WHEN old.name <> new.name
			OR old.active <> new.active
			OR old.date <> new.date
	BEGIN
		insert into log_cue (
			row_id,
			old_active,
			new_active,
			old_name,
			new_name,
			old_date,
			new_date,
			change_type,
			created_at
		) 
		values (
			old.id,
			old.active,
			new.active,
			old.name,
			new.name,
			old.date,
			new.date,
			'UPDATE',
			strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
		);
	END;
	`
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "tr4igger not created")
	}

	// stmt, err := db.Prepare("INSERT INTO cue (active, name, dt) VALUES (?, ?, ?)")
	stmt, err := db.Prepare("INSERT INTO cue (active, name, date) VALUES (?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer stmt.Close()
	dt := time.Now()
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
	dt = time.Date(dt.Year(), dt.Month(), 10, 0, 0, 0, 0, time.UTC)
	// if _, err = stmt.Exec(true, "bb34b4b234bb234", time.Now().Add(time.Duration(rand.IntN(10))*time.Minute).Format(sqliteLayout)); err != nil {
	if _, err = stmt.Exec(true, "bb34b4b234bb234", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "")
	}
	dt = time.Date(dt.Year(), dt.Month(), 2, 0, 0, 0, 0, time.UTC)
	if _, err = stmt.Exec(true, "sadjlhasklhasjkldh", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "")
	}
	dt = time.Date(dt.Year(), dt.Month(), rand.IntN(20), 0, 0, 0, 0, time.UTC)
	if _, err = stmt.Exec(true, "ÇASDçasçdasçfkasjfas", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "")
	}

	dt = dt.Add(time.Duration(rand.IntN(10)) * time.Minute)
	// time.Now().Add(time.Duration(rand.IntN(10)) * time.Minute)
	if _, err := db.Exec("insert into cue (name, date) values ('TESssssTESTE', ?)", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "trigger not created")
	}
	// dt = dt.AddDate(0, 0, 1)
	dt = time.Date(dt.Year(), dt.Month(), rand.IntN(20), 0, 0, 0, 0, time.UTC)
	if _, err := db.Exec("insert into cue (name, date) values ('TESTESTE', ?)", dt.Format(time.DateOnly)); err != nil {
		return errors.Wrap(err, "trigger not created")
	}

	qry = `
	CREATE TABLE IF NOT EXISTS scheduler (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		start_exec DATETIME,
		end_exec DATETIME,
		next_exec DATETIME
	)
	`
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "trigger not created")
	}

	return nil
}

func saveLog() {
	// CREATE TABLE IF NOT EXISTS log (
	// 	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	//  table TEXT,
	//  rowId INTEGER
	// 	date NUMERIC,
	// 	log TEXT
	// )`
	qry := "INSERT INTO log () VALUES ();"
	if _, err := db.Exec(qry); err != nil {
		log.Error().Stack().Err(err).Msg("log error")
		// return err
	}
}

func listCue() ([]Cue, error) {
	var cues []Cue = []Cue{}

	rows, err := db.Query("SELECT id, active, done, status, name, date FROM cue WHERE active = TRUE;")
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("listCue() / connections")
	for rows.Next() {
		cue := Cue{}
		err = rows.Scan(&(cue.Id), &(cue.active), &(cue.Done), &(cue.Status), &(cue.Name), &(cue.Date))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		cues = append(cues, cue)
	}

	return cues, nil
}

func createCue(cue Cue) error {
	log.Debug().Interface("cue", cue).Msg("createCue()")
	parsedDate, err := parseIsoDateTime(cue.Date)
	if err != nil {
		return errors.Wrap(err, "")
	}
	qry := "INSERT INTO cue (active, name, date) VALUES (TRUE, ?, ?)"
	result, err := dbExec(qry, cue.Name, parsedDate.Format(time.DateOnly))
	if err != nil {
		return errors.Wrap(err, "")
	}

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("createCue() / connections")

	id, err := result.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.Info().Msgf("new cue. id: %d", id)

	return nil
}

func parseIsoDateTime(dt interface{}) (time.Time, error) {
	v, ok := dt.(string)
	if !ok {
		return time.Time{}, errors.New("datetime is not a string")
	}
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString
	// expected format: YYYY-MM-DDTHH:mm:ss.sssZ from javascript new Date().toISOString()
	dtParsed, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return time.Time{}, err
	}
	return dtParsed, nil
}

func updateCue(id int64, newValues map[string]interface{}) error {
	var queryColumns []string
	var queryValues []any
	var allowedFields []string = []string{"status", "name", "date", "dt"}
	// detect fields to be used in where clause
	for _, field := range allowedFields {
		if value, ok := newValues[field]; ok {
			queryColumns = append(queryColumns, field+" = ?")
			if field == "date" {
				parsedDate, err := parseIsoDateTime(value)
				if err != nil {
					return errors.Wrap(err, "parseIsoDateTime()")
				}
				queryValues = append(queryValues, parsedDate.Format(time.DateOnly))
			} else if field == "dt" {
				// TODO remove field dt
				log.Info().Str("dt 1", value.(string)).Msg("")
				dtParsed, _ := parseIsoDateTime(value)
				// dtParsed, _ := time.Parse(sqliteLayout, value.(string))
				log.Info().Str("dt 2", dtParsed.Format(sqliteLayout)).Msg("")
				queryValues = append(queryValues, dtParsed.Format(sqliteLayout))
			} else {
				queryValues = append(queryValues, value)
			}
		}
	}

	queryValues = append(queryValues, id)
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "db.Begin()")
	}
	defer tx.Commit()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("updateCue() / connections")

	qry := "UPDATE cue SET " + strings.Join(queryColumns, ", ") + " WHERE id = ? AND active = TRUE;"

	log.Debug().Stack().Str("query_prepared", qry).Msg("")
	log.Debug().Interface("query_values", queryValues).Msg("")

	stmt, err := tx.Prepare(qry)
	if err != nil {
		return errors.Wrap(err, "Prepare")
	}
	defer stmt.Close()
	res, err := stmt.Exec(queryValues...)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}
	count, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "")
	}
	if count != 1 {
		log.Info().Msgf("rowsAffected: %v", count)
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "")
		}
		log.Info().Msg("update not concluded.pdate not concluded. rollback executed.")
		return errors.New("an error occur while updating")
	}
	return nil
}

func removeCue(id int64) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer tx.Commit()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("removeCue() / connections")

	stmt, err := tx.Prepare("UPDATE cue SET active = FALSE WHERE id = ?;")
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "")
	}
	count, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "")
	}
	if count != 1 {
		log.Info().Msgf("rowsAffected: %v", count)
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "")
		}
		log.Info().Msg("update not concluded. rollback executed.")
		return errors.New("an error occur while removing")
	}
	log.Info().Msgf("cue removed. id: %d", id)
	return nil
}

func changeCueStatus(id int64, s status) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer tx.Commit()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("changeCueStatus() / connections")

	stmt, err := tx.Prepare("UPDATE cue SET status = ? WHERE id = ?;")
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(s, id)
	if err != nil {
		return errors.Wrap(err, "")
	}
	count, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "")
	}
	if count != 1 {
		log.Info().Msgf("rowsAffected: %v", count)
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "")
		}
		log.Info().Msg("update not concluded. rollback executed.")
		return errors.New("an error occur while removing")
	}
	log.Info().
		Int64("id", id).
		Str("new_status", s.String()).
		Msg("cue status changed.")
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
		Int("open", db.Stats().OpenConnections).
		Msg("listSchedulerHistory() / connections")
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
	cues, err := listCue()
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
		if cue.Done {
			newStat := make(map[string]interface{})
			newStat["status"] = 1
			err := updateCue(cue.Id, newStat)
			if err != nil {
				log.Error().Stack().Err(err).Msg("")
			} else {
				log.Debug().Int("id", int(cue.Id)).Msg("uptated")
			}
		}
	}
	// }
	// log.Info().Interface("cues", cuesStats).Msg("")
}

func checkDueDates() error {
	cues, err := listCue()
	if err != nil {
		return errors.Wrap(err, "")
	}
	for _, cue := range cues {
		cueDate, err := time.Parse(time.DateOnly, cue.Date)
		if err != nil {
			return errors.Wrap(err, "")
		}
		now := time.Now()
		now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

		if cueDate.Equal(now) && cue.Status == pending && cue.Status != due {
			if err = changeCueStatus(cue.Id, due); err != nil {
				return errors.Wrap(err, "")
			}
		} else if cueDate.Before(now) && cue.Status != paid && cue.Status != overdue {
			if err = changeCueStatus(cue.Id, overdue); err != nil {
				return errors.Wrap(err, "")
			}
		} else if cueDate.After(now) && cue.Status != paid && cue.Status != pending {
			if err = changeCueStatus(cue.Id, pending); err != nil {
				return errors.Wrap(err, "")
			}
		}
	}
	return nil
}

func executeScheduler(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				{
					log.Info().Msg("scheduler start")
					result, err := dbExec("INSERT INTO scheduler (start_exec) VALUES (?);", time.Now().Format(sqliteLayout))
					if err != nil {
						log.Error().Stack().Err(err).Msg("")
					}
					scheduler_id, err := result.LastInsertId()
					if err != nil {
						log.Error().Stack().Err(err).Msg("")
					}

					checkDueDates()

					_, err = dbExec("UPDATE scheduler SET end_exec = ? WHERE id = ?;", time.Now().Format(sqliteLayout), scheduler_id)
					if err != nil {
						log.Error().Stack().Err(err).Msg("")
					}
					log.Info().Msg("scheduler end")
				}
			}
		}
	}()
}
