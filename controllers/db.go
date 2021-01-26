package controllers

import (
	"database/sql"
	"fmt"
	"goapi/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	host     string
	user     string
	password string
	port     string
	database string
	table    string
}

func InitDB(user string, password string, host string, port string, database string, table string) *Database {
	return &Database{user: user, password: password, host: host, port: port, database: database, table: table}
}

func (db Database) connString() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%s", db.user, db.password, db.host, db.port, db.database)
}

func (db Database) CheckDB() (err error) {
	conStr := db.connString()
	conn, err := sql.Open("mysql", conStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	return conn.Ping()
}

func (db Database) CheckTable(dbName string) (err error) {
	conStr := db.connString()
	conn, err := sql.Open("mysql", conStr)
	if err != nil {
		return err
	}
	defer conn.Close()
	testQ := fmt.Sprintf(`SELECT TABLE_NAME 
			FROM information_schema.tables
			WHERE table_schema = "%s" 
			AND table_name = "%s" LIMIT 1;`, dbName, db.table)

	row := conn.QueryRow(testQ)
	returnRow := ""
	err = row.Scan(&returnRow)
	if err != nil {
		st := fmt.Sprintf("CREATE Table %s (id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, Username VARCHAR(30) NOT NULL, Password VARCHAR(30) NOT NULL, Firstname VARCHAR(30) NOT NULL, Lastname VARCHAR(30) NOT NULL);", db.table)
		stmt, err := conn.Prepare(st)
		if err != nil {
			return err
		}
		_, err = stmt.Exec()
		fmt.Println("creation error: ", err)

		if err != nil {
			return err
		}
	}
	return nil
}

func (db Database) SelectAll(targetUser *[]models.User) error {
	conStr := db.connString()
	conn, err := sql.Open("mysql", conStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	q := fmt.Sprintf("SELECT * FROM %s", db.table)
	rows, err := conn.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.User
		err = rows.Scan(&t.Id, &t.Username, &t.Password, &t.Firstname, &t.Lastname)
		if err != nil {
			return err
		}
		*targetUser = append(*targetUser, t)
	}
	return nil
}

func (db Database) InsertUSER(user models.User) error {
	conStr := db.connString()
	conn, err := sql.Open("mysql", conStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	q := fmt.Sprintf(`INSERT INTO %s (Username, Password, Firstname, Lastname) VALUES ("%s", "%s", "%s", "%s")`, db.table, user.Username, user.Password, user.Firstname, user.Lastname)
	insert, err := conn.Query(q)
	if err != nil {
		return err
	}
	defer insert.Close()
	return err
}

func (db Database) GetUser(user *models.User, id string) error {
	conStr := db.connString()
	conn, err := sql.Open("mysql", conStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	q := fmt.Sprintf("SELECT * FROM user WHERE id=%v LIMIT 1", id)
	row := conn.QueryRow(q)
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Firstname, &user.Lastname)
	if err != nil {
		return err
	}
	return nil
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
