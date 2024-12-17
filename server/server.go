package main

import (
	"database/sql"
	"flag"
	"time"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var db *sql.DB

// Same layout as sqlite strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
var sqliteLayout string = "2006-01-02 15:04:05.000"

func main() {
	// time.Local = time.UTC
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: sqliteLayout})
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Load env variables from file
	err := godotenv.Load()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Error while reading .env file")
	}

	renew := false
	// renew = true
	// Data base init
	dbFile := "rentmem.db"
	if renew {
		os.Remove(dbFile)
	}
	db, err = sql.Open("sqlite3", dbFile+"?cache=shared")
	if err != nil {
		log.Fatal().Err(err).Msg("Can't open sqlite file")
	}
	db.SetMaxOpenConns(2)
	defer db.Close()
	if renew {
		err = initDB()
		if err != nil {
			log.Fatal().Stack().Err(err).Msg("initDataBase error")
		}
	}
	// Start server
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:       true,
		LogStatus:       true,
		LogURI:          true,
		LogURIPath:      true,
		LogProtocol:     true,
		LogResponseSize: true,
		LogRemoteIP:     true,
		LogUserAgent:    true,
		LogFormValues:   []string{},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Status >= 400 {
				log.Error().
					Int("status", v.Status).
					Str("method", v.Method).
					Str("uri", v.URI).
					Int("resp_size", int(v.ResponseSize)).
					Msg("")
			} else {
				log.Info().
					Int("status", v.Status).
					Str("method", v.Method).
					Str("uri", v.URI).
					Int("resp_size", int(v.ResponseSize)).
					Msg("")
			}

			return nil
		},
		// LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		// 	// msg := fmt.Sprintf(" \"%-6s %v %v\" %v %v %v ",
		// 	msg := fmt.Sprintf(" \"%v %v %v\" %v %v %v ",
		// 		v.Method,
		// 		v.URI,
		// 		v.Protocol,
		// 		v.Status,
		// 		v.ResponseSize,
		// 		v.UserAgent)
		// 	log.Info().Msg(msg)
		// 	return nil
		// },
	}))

	// Routes
	e.Static("/", "web")
	e.GET("/rent", getRent)
	e.POST("/rent", postRent)
	e.PUT("/rent/:id", putRent)
	e.DELETE("/rent/:id", delRent)

	e.POST("/process-rent/:id", processRent)
	e.POST("/process-rent", processRent)
	e.GET("/reminder-detail/:id", getReminderDetail)
	e.GET("/reminders", getReminders)
	e.POST("/pay/rent/:id", payRent)
	e.POST("/process-reminders", processReminders)

	e.GET("/scheduler/history", getHistory)
	// e.GET("/stats", getHistory)

	e.GET("/initdb/:clear", initdb)
	e.GET("/initdb", initdb)

	// e.POST("/start", sstart)
	// e.POST("/stop", sstop)
	// e.POST("/test", test)
	// e.POST("/tenants", postTenant)
	// e.GET("/tenants", getTenant)
	// e.PUT("/tenants/:id", putTenant)
	// e.DELETE("/tenants/:id", deleteTenant)

	// e.POST("/properties", postProperty)
	// e.GET("/properties", getProperty)
	// e.PUT("/properties/:id", putProperty)
	// e.DELETE("/properties/:id", deleteProperty)

	// e.POST("/rents", postRent)
	executeScheduler(1 * time.Hour)

	// fmt.Print(time.Local)
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
