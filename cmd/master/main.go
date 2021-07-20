package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

type User struct {
	Id            int64
	Email         string
	EncryptedPass string
}

// ErrorState Function checking error state and shutting down program if state is err
func ErrorState(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CheckDataBaseConn DataBase conn use lazy creating, cause this
//function ping db and return error if status unavailable
func CheckDataBaseConn(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	dbAddr := "0.0.0.0:5432"
	conn, err := ConnDB("passwd", dbAddr, "master")
	if err != nil {
		ErrorState(err)
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			ErrorState(err)
		}
	}(conn)

	err = CheckDataBaseConn(conn)
	if err != nil {
		ErrorState(err)
	}

	query := `CREATE TABLE IF NOT EXISTS users (
		id bigserial not null primary key,
		email varchar not null unique,
		encrypted_password varchar not null
	);`

	_, err = conn.Exec(query)
	if err != nil {
		log.Warning(err)
		time.Sleep(10 * time.Second)

		_, err = conn.Exec(query)
		if err != nil {
			log.Warning(err)
		}
	}

}
