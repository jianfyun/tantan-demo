package storage

import (
	"tantan-demo/config"
	"time"

	"github.com/go-pg/pg"
)

var (
	db *pg.DB
)

// Connect initializes database connection pool with configuration params.
func Connect() {
	db = pg.Connect(&pg.Options{
		Addr:         config.Config["postgres.addr"].Str,
		User:         config.Config["postgres.user"].Str,
		Password:     config.Config["postgres.password"].Str,
		Database:     config.Config["postgres.database"].Str,
		DialTimeout:  time.Duration(config.Config["postgres.timeout"].Int) * time.Second,
		ReadTimeout:  time.Duration(config.Config["postgres.timeout"].Int) * time.Second,
		WriteTimeout: time.Duration(config.Config["postgres.timeout"].Int) * time.Second,
	})
}

// Close closes the database connection pool.
func Close() {
	db.Close()
}
