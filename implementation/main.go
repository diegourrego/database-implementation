package implementation

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Init() {
	datasource := "root:Diego123@tcp(localhost:3306)/prueba_db"
	storageDB, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	if err := storageDB.Ping(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Println("database configuration")
}
