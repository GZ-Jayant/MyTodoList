package db

import (
	"database/sql"
	"encoding/json"
	"gin-webapi/config"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	TaskID        int    `json:"task_id"`
	TaskName      string `json:"task_name"`
	TaskDesc      string `json:"task_desc"`
	TaskStatus    string `json:"task_status"`
	TaskStartDate string `json:"task_start_date"`
	TaskEndDate   string `json:"task_end_date"`
}

var (
	dbEngine []string // database engine
	dbConn   *sql.DB  // database connection

	maxTaskID chan int // max task id
)

func init() {
	type dbConfig struct {
		DBType     string `json:"dbType"`     // database type
		DBHost     string `json:"dbHost"`     // database host
		DBPort     string `json:"dbPort"`     // database port
		DBUsername string `json:"dbUsername"` // database username
		DBPassword string `json:"dbPassword"` // database password
		DBDatabase string `json:"dbDatabase"` // database name
	}
	cfg := new(dbConfig) // database config

	bytes, err := config.LoadConfig("db_config.json")
	config.ErrorLog(err, true, "load db config error")

	err = json.Unmarshal(bytes, &cfg)
	config.ErrorLog(err, true, "unmarshal db config error")

	dbEngine = append(dbEngine, cfg.DBType, cfg.DBUsername+":"+cfg.DBPassword+"@tcp("+cfg.DBHost+":"+cfg.DBPort+")/"+cfg.DBDatabase)

}

// Connect connect to database;
// it's only called once when server start
func Connect() {
	var err error
	dbConn, err = sql.Open(dbEngine[0], dbEngine[1])
	config.ErrorLog(err, true, "open database error")

	createTask := `create table IF NOT EXISTS todo_task_info(
		task_id int primary key,
		task_name varchar(100),
		task_desc varchar(100),
		task_status varchar(100),
		task_start_date date,
		task_end_date date
	);`

	// check your table exist or not, if not, create it
	// todo_task
	_, err = dbConn.Exec(createTask)
	config.ErrorLog(err, true, "create table todo_task_info error")

	// get max task id
	maxTaskID = make(chan int, 1)
	go func() {
		var id sql.NullInt32
		err = dbConn.QueryRow("select max(task_id) from todo_task_info").Scan(&id)
		config.ErrorLog(err, true, "get max task id error")
		if id.Valid {
			maxTaskID <- int(id.Int32)
		} else {
			maxTaskID <- 0
		}
	}()

}

// Close close database connection;
// it's only called once when server stop
func Close() {
	dbConn.Close()
}
