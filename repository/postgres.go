package repository

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"log"
)

type Repository interface {
	GetRootSpecializations() []string
	GetChildSpecializations(specialization string) []string
}

type PostgresService struct {
	Connection *pgx.Conn
}

func New() PostgresService {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "postgres",
	})

	if err != nil {
		log.Panic("Error while connecting to database: ", err)
	}

	return PostgresService{Connection: conn}
}

func (postgres PostgresService) GetRootSpecializations() []string {
	rows, err := postgres.Connection.Query("select getRootSpecs()")
	var specs []string

	if err != nil {
		log.Panic("Error while calling stored procedure getRootSpec(): ", err)
	}

	var row pgtype.Varchar
	for rows.Next() {
		err := rows.Scan(&row)
		if err != nil {
			log.Panic("Error while scanning root specializations: ", err)
		}
		specs = append(specs, row.String)
	}

	return specs
}

func (postgres PostgresService) GetChildSpecializations(specialization string) []string {
	rows, err := postgres.Connection.Query("select getChildByParent($1)", specialization)
	var specs []string

	if err != nil {
		log.Panic("Error while calling stored procedure getChildByParent("+specialization+"): ", err)
	}

	var row pgtype.Varchar
	for rows.Next() {
		err := rows.Scan(&row)
		if err != nil {
			log.Panic("Error while scanning child specializations: ", err)
		}
		specs = append(specs, row.String)
	}

	return specs
}
