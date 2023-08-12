package repository

import (
	"database/sql"
	"log"
	"user-service/internal/core/port/repository"
	"user-service/internal/infra/configDB"
)

type database struct {
	db *sql.DB
}

func NewDB(conf configDB.ConfigDatabase) repository.Database {
	db, err := newDatabase(conf)

	if err != nil {
		log.Println("Create database failed: ", err.Error())
		return nil
	}

	return &database{
		db: db,
	}
}

func newDatabase(conf configDB.ConfigDatabase) (*sql.DB, error) {
	db, err := sql.Open(conf.Driver, conf.URL())

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *database) Close() error {
	return db.Close()
}

func (db *database) GetDB() *sql.DB {
	return db.db
}
