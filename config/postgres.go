package config

import (
	"database/sql"
	"fmt"
	"os"
	
	_ "github.com/lib/pq"
)

type Postgres struct {
	Username string
	Password string
	Port string
	Address string
	Database string

	DB *sql.DB
}

type PsqlDb struct {
	*Postgres
}

var (
	PSQL *PsqlDb
)

func InitPostgres() error {
	PSQL = new(PsqlDb)

	PSQL.Postgres = &Postgres{
		Username: 	os.Getenv("POSTGRES_USER"),
		Password: 	os.Getenv("POSTGRES_PASSWORD"),
		Address: 	os.Getenv("POSTGRES_ADDRESS"),
		Port: 		os.Getenv("POSTGRES_PORT"),
		Database: 	os.Getenv("POSTGRES_DB"),
	}

	err := PSQL.Postgres.OpenConnection()
	if err != nil {
		return err
	}

	err = PSQL.DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) OpenConnection() error {

	query := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Address, p.Port, p.Username, p.Password, p.Database)

	dbConnection, err := sql.Open("postgres", query)
	if err != nil {
		return err
	}

	p.DB = dbConnection

	err = p.DB.Ping()
	if err != nil{
		return err
	}

	fmt.Println("Berhasil konek ke database")

	return nil

}