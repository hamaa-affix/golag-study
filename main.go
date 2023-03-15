package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	// "io"
	// "log"
	// "net/http"
)

type DBConf struct {
	UserName   string
	PassWord   string
	DBName     string
	Port       int
	DriverName string
	Host       string
	Err        error
}

func NewDBConf(
	userName string,
	pass string,
	dbName string,
	port int,
	driverName string,
	host string,
) *DBConf {
	return &DBConf{
		UserName:   userName,
		PassWord:   pass,
		DBName:     dbName,
		Port:       port,
		DriverName: driverName,
		Host:       host,
	}
}

func (d *DBConf) CreateEndpoint() string {
	if d.Err != nil {
		log.Fatal("missing create db connection endpoint\nerror is %v", d.Err)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.UserName, d.PassWord, d.Host, d.Port, d.DBName)
}

func main() {
	fmt.Println(os.Getenv("DATABASE_CONECTION_URL"))
	conf := NewDBConf(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
		3306,
		"mysql",
		"mysql",
	)

	db, err := sql.Open("mysql", conf.CreateEndpoint())
	if err != nil {
		log.Fatal("mysql connection error\nerror message:\n %v", err)
	}
	result, err := db.Exec(`SHOW TABLES FROM hoge;`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// h1 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	// }
	// h2 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #2!\n")
	// }

	// http.HandleFunc("/", h1)
	// http.HandleFunc("/endpoint", h2)

	// log.Fatal(http.ListenAndServe(":8080", nil))
}
