package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // we need to define the pgx driver to use to connect to database
	"github.com/joho/godotenv"
)

// to run server on nodemon we can use
// nodemon --exec go run main.go --signal SIGTERM
// const portNumber = "127.0.0.1:8080"

func main() {

	// loading .env file for database configuration
	// create .env file and it should have PostgresDBConfiguration = "host=localhost port=5432 dbname=animals user=_____ password=_____" // file details accordingly
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	DbConfiguration := os.Getenv("PostgresDBConfiguration")
	
	// connection to database using pgx package and driver
	// conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=animals user=_____ password=_____")
	// to use template given on github by jackc, do remeber to write EXPORT PGADMIN=postgres ie name of database server
	conn, err := sql.Open("pgx", DbConfiguration)
	
	if err != nil {
		log.Fatalf("Unable to connect %v\n",err)
	}

	defer conn.Close()

	log.Println("\nConnection established to Database")

	//test my connection by pinging it gives back error if not get a pong response
	err = conn.Ping()

	if err!= nil {
		log.Fatalf("Cannont ping database")
	}

	fmt.Println("Pinged database")

	// Get rows from database
	getAllRows(conn)
	
}

//fetch all rows from database
func getAllRows(conn *sql.DB) error {

	rows, err := conn.Query("SELECT id, animal_name FROM animals")

	if err != nil {
		log.Println(err)
		return err
	}
	// always defer-close the connection 
	defer rows.Close()

	var id int
	var animal_name string
	
	// we traverse the loop till rows are present in the database
	//Then we scan and save the columns in variables
	for rows.Next() {
		err := rows.Scan(&id, &animal_name)

		if err != nil {
			log.Println(err)
			return err
		}

		fmt.Println("Record is", id, animal_name)
	}

	// below is a good practice as i might get an error while trying to traversing above row
	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)	
	}

	fmt.Println("----------------------------------------------------")

	return nil
}