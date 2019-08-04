package internal

import (
	"database/sql"

	mapper "github.com/birkirb/loggers-mapper-logrus"
	"gopkg.in/birkirb/loggers.v1"

	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Resolver struct {
	config *Config
	logger loggers.Contextual
	db     *sql.DB
	router *mux.Router
}

func NewResolver(c *Config) *Resolver {
	return &Resolver{
		config: c,
	}
}

type LoggerApp interface {
	ResolveLogger() loggers.Contextual
}

func (r *Resolver) ResolveLogger() loggers.Contextual {
	if r.logger == nil {
		l := logrus.New()
		l.Out = os.Stdout
		l.Level = logrus.InfoLevel
		l.SetFormatter(&logrus.JSONFormatter{})
		r.logger = mapper.NewLogger(l)
	}
	return r.logger
}
func (r *Resolver) ResolveDB(dbName string) *sql.DB {
	if r.db == nil {
		db, err := sql.Open(dbName, r.config.DB.GetConnectionString())
		if err != nil {
			panic(err)
		}

		// find a place for these
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(5)

		return db
	}
	return r.db
}

func NewRouter() *mux.Router {
	return mux.NewRouter()
}
