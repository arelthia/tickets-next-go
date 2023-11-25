package data

import (
	"fmt"
	"os"
	"tickets-next-go/api/model"

	"github.com/golobby/dotenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DataSource string `env:"DATA_SOURCE"`
}

type DbRunner struct {
	db *sqlx.DB
}

var database *DbRunner

func GetConnection() {
	driverName := "postgres"
	var envvars Config
	file, err := os.Open("assets/.env")
	if err != nil {
		fmt.Printf("error opening env file: %v\n", err)
		return
	}
	err = dotenv.NewDecoder(file).Decode(&envvars)
	if err != nil {
		fmt.Printf("error decoding env file: %v\n", err)
		return
	}
	dataSource := envvars.DataSource
	db, err := sqlx.Connect(driverName, dataSource)
	if err != nil {
		fmt.Printf("error connecting to database: %v\n", err)
		return
	}
	database = &DbRunner{db: db}
}

func GetDb() *DbRunner {
	if database != nil {
		return database
	}
	GetConnection()
	fmt.Println("connected to database")
	return database
}

// func (d *DbRunner) Close() {
// 	d.db.Close()
// }

func (dataSource *DbRunner) GetTickets() ([]model.Ticket, error) {
	tickets := []model.Ticket{}
	err := dataSource.db.Select(&tickets, "SELECT * FROM tickets")
	return tickets, err
}

func (dataSource *DbRunner) GetTicket(id int) (model.Ticket, error) {
	ticket := model.Ticket{}
	err := dataSource.db.Get(&ticket, "SELECT * FROM tickets WHERE id = $1", id)
	return ticket, err
}

func (dataSource *DbRunner) CreateTicket(t model.Ticket) (int64, error) {
	result, err := dataSource.db.Exec("INSERT INTO tickets (first_name, last_name, email, issue, priority) VALUES ($1, $2, $3, $4, $5)",
		t.FirstName, t.LastName, t.Email, t.Issue, t.Priority)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (dataSource *DbRunner) UpdateTicket(id int, t model.Ticket) (int64, error) {
	result, err := dataSource.db.Exec("UPDATE tickets SET first_name = $1, last_name = $2, email = $3, issue = $4, priority = $5 WHERE id = $6",
		t.FirstName, t.LastName, t.Email, t.Issue, t.Priority, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (dataSource *DbRunner) DeleteTicket(id int) (int64, error) {
	result, err := dataSource.db.Exec("DELETE FROM tickets WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
