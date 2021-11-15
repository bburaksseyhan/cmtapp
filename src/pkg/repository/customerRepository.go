package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bburaksseyhan/ctmapp/src/cmd/utils"
	"github.com/bburaksseyhan/ctmapp/src/pkg/entities"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// CustomerRepository related with context and response model
type CustomerRepository interface {
	List(cntxt context.Context, timeout int) ([]entities.CustomerEntity, error)
	Add(customer entities.CustomerEntity, cntxt context.Context, timeout int) (entities.CustomerEntity, error)
	Delete(id int, cntxt context.Context, timeout int) (bool, error)
	Get(id int, cntxt context.Context, timeout int) (entities.CustomerEntity, error)
}

type postgresCustomerRepository struct {
	dbSetting *utils.DbSettings
	conn      string
}

// NewPostgresCustomerRepository crerate new postgresCustomerRepository
func NewPostgresCustomerRepository(dbSettings *utils.DbSettings) CustomerRepository {
	//initial log formatter
	log.SetFormatter(&log.JSONFormatter{})

	repo := &postgresCustomerRepository{
		dbSetting: dbSettings,
		conn: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbSettings.Host, dbSettings.Port, dbSettings.User, dbSettings.Password, dbSettings.DbName),
	}

	return repo
}

func (r *postgresCustomerRepository) List(cntxt context.Context, timeout int) ([]entities.CustomerEntity, error) {

	//context
	ctx, cancel := context.WithTimeout(cntxt, time.Duration(timeout)*time.Second)
	defer cancel()

	//connect database
	db := openDatabaseConn(r)
	defer db.Close()

	//read data from server
	rows, _ := db.QueryContext(ctx, "Select id,firstname,lastname from Users")
	defer rows.Close()

	//define slice for store customer information
	var customerEntity []entities.CustomerEntity

	//read data row by row
	for rows.Next() {
		var userId int
		var firstName string
		var lastName string

		_ = rows.Scan(&userId, &firstName, &lastName)

		customerEntity = append(customerEntity, entities.CustomerEntity{Id: userId, FirstName: firstName, LastName: lastName})
	}

	return customerEntity, nil
}

func (r *postgresCustomerRepository) Add(customer entities.CustomerEntity, cntxt context.Context, timeout int) (entities.CustomerEntity, error) {

	//context
	ctx, cancel := context.WithTimeout(cntxt, time.Duration(timeout)*time.Second)
	defer cancel()

	//connect database
	db := openDatabaseConn(r)
	defer db.Close()

	//add data
	query := "Insert into Users(FirstName, LastName) values($1,$2)"
	if _, err := db.ExecContext(ctx, query, customer.FirstName, customer.LastName); err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *postgresCustomerRepository) Delete(id int, cntxt context.Context, timeout int) (bool, error) {

	//context
	ctx, cancel := context.WithTimeout(cntxt, time.Duration(timeout)*time.Second)
	defer cancel()

	//connect database
	db := openDatabaseConn(r)
	defer db.Close()

	query := "Delete From Users Where Id=$1"
	affectedRow, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return false, nil
	}

	fmt.Println(affectedRow.LastInsertId())
	fmt.Println(affectedRow.RowsAffected())

	return true, nil
}

func (r *postgresCustomerRepository) Get(id int, cntxt context.Context, timeout int) (entities.CustomerEntity, error) {

	//context
	ctx, cancel := context.WithTimeout(cntxt, time.Duration(timeout)*time.Second)
	defer cancel()

	//connect database
	db := openDatabaseConn(r)
	defer db.Close()

	data := db.QueryRowContext(ctx, "Select id,firstname,lastname from Users Where Id=$1", id)

	var userId int
	var firstName string
	var lastName string

	_ = data.Scan(&userId, &firstName, &lastName)

	return entities.CustomerEntity{Id: userId, FirstName: firstName, LastName: lastName}, nil
}

func openDatabaseConn(r *postgresCustomerRepository) *sql.DB {
	db, err := sql.Open("postgres", r.conn)
	if err != nil {
		log.Error("Connection failed")
	}

	pingError := db.Ping()
	if pingError != nil {
		log.Error("Ping != pong")
	}

	log.Info("Postgres connection success!!!")

	return db
}
