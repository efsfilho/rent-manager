package main

import (
	"database/sql"
	"strconv"
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

type Cue struct {
	Id     int64 `json:"id"`
	active bool
	Done   bool   `json:"done"`
	Date   string `json:"date"`
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

func initDataBase() error {
	qry := `
		CREATE TABLE IF NOT EXISTS cue (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			active BOOLEAN DEFAULT TRUE,
			done BOOLEAN DEFAULT FALSE,
			name TEXT NOT NULL,
			date TEXT	DEFAULT (DATE('0', 'unixepoch')) CHECK(length(date) == 10),
			created_at NUMERIC DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now'))
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
		old_name text not null,
		new_name text not null,
		old_date TEXT not null,
		new_date TEXT not null,
		change_type text not null,
		created_at NUMERIC not null
	);
	
`
	if _, err := db.Exec(qry); err != nil {
		return errors.Wrap(err, "log_cue not created")
	}
	qry = `
	CREATE TRIGGER log_cue_after_update 
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
		return errors.Wrap(err, "trigger not created")
	}

	stmt, err := db.Prepare("INSERT INTO cue (active, name) VALUES (?, ?)")
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer stmt.Close()
	if _, err = stmt.Exec(true, "a1a1a1"); err != nil {
		return errors.Wrap(err, "")
	}
	if _, err = stmt.Exec(true, "a2a22aa2"); err != nil {
		return errors.Wrap(err, "")
	}
	if _, err = stmt.Exec(true, "bb34b4b234bb234"); err != nil {
		return errors.Wrap(err, "")
	}
	if _, err = stmt.Exec(true, "sadjlhasklhasjkldh"); err != nil {
		return errors.Wrap(err, "")
	}
	if _, err = stmt.Exec(true, "ÇASDçasçdasçfkasjfas"); err != nil {
		return errors.Wrap(err, "")
	}

	// time.Local = time.UTC
	// time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	dt := time.Now()
	if _, err := db.Exec("insert into cue (name, date) values ('TESssssTESTE', ?)", dt.Format(time.DateOnly)); err != nil {
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
	rows, err := db.Query("SELECT id, active, done, name, date FROM cue WHERE active = TRUE;")
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("connections")

	for rows.Next() {
		cue := Cue{}
		err = rows.Scan(&(cue.Id), &(cue.active), &(cue.Done), &(cue.Name), &(cue.Date))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		cues = append(cues, cue)
	}

	return cues, nil
}

func createCue(cue Cue) error {
	qry := "INSERT INTO cue (active, name, date) VALUES (TRUE, ?, ?)"
	r, err := dbExec(qry, cue.Name, time.Now().Format(sqliteLayout))
	if err != nil {
		return errors.Wrap(err, "")
	}

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("connections")

	id, err := r.LastInsertId()
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

func updateCue(id string, newValues map[string]interface{}) error {
	var queryColumns []string
	var queryValues []any
	var allowedFields []string = []string{"done", "name", "date"}
	if _, err := strconv.Atoi(id); err != nil {
		log.Error().Msgf("id value: %v", id)
		return errors.Wrap(err, "id not valid")
	}

	for _, field := range allowedFields {
		if value, ok := newValues[field]; ok {
			queryColumns = append(queryColumns, field+" = ?")
			if field == "date" {
				dateValue, err := parseIsoDateTime(value)
				if err != nil {
					return errors.Wrap(err, "")
				}
				// t, _ := i.(time.Time)
				queryValues = append(queryValues, dateValue.Format(time.DateOnly))
			} else {
				queryValues = append(queryValues, value)
			}
		}
	}

	// last value used in where clause
	queryValues = append(queryValues, id)
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer tx.Commit()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("connections")

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
		log.Info().Msgf("Update not concluded. RowsAffected: %v", count)
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "")
		}
		log.Info().Msg("Rollback executed.")
		return errors.New("an error occur while updating")
	}

	return nil
}

func removeCue(id string) error {
	if _, err := strconv.Atoi(id); err != nil {
		log.Error().Msgf("id value: %v", id)
		return errors.Wrap(err, "id not valid")
	}
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer tx.Commit()

	log.Debug().
		Int("in_use", db.Stats().InUse).
		Int("open", db.Stats().OpenConnections).
		Msg("connections")

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
	log.Info().Msgf("RowsAffected: %v", count)
	if count != 1 {
		log.Info().Msgf("Update not concluded. RowsAffected: %v", count)
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "")
		}
		log.Info().Msg("Rollback executed.")
		return errors.New("an error occur while removing")
	}
	log.Info().Msgf("cue removed. id: %s", id)
	return nil
}
