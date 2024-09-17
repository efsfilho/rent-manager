package main

import (
	"database/sql"
	"flag"

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

	// Data base init
	dbFile := "./rentmem.db"
	os.Remove(dbFile)
	db, err = sql.Open("sqlite3", "file:./rentmem.db?cache=shared")
	if err != nil {
		log.Fatal().Err(err).Msg("Can't open sqlite file")
	}
	db.SetMaxOpenConns(2)
	defer db.Close()

	err = initDataBase()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("initDataBase error")
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
	e.GET("/cue", getCue)
	e.POST("/cue", postCue)
	e.PUT("/cue/:id", putCue)
	e.DELETE("/cue/:id", delCue)
	// e.POST("/tenants", postTenant)
	// e.GET("/tenants", getTenant)
	// e.PUT("/tenants/:id", putTenant)
	// e.DELETE("/tenants/:id", deleteTenant)

	// e.POST("/properties", postProperty)
	// e.GET("/properties", getProperty)
	// e.PUT("/properties/:id", putProperty)
	// e.DELETE("/properties/:id", deleteProperty)

	// e.POST("/rents", postRent)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
