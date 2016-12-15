package utils

import (

	// Postgre Driver

	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// GORM for PostgreSQL
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	//DBType  数据库类型
	DBType string = "postgres"
	// DBConnDev  连接对象(开发环境)
	DBConnDev string = "host=localhost user=andysilver dbname=keerp sslmode=disable password=65570480"
	// DBConnProc 连接对象(应用环境)
	DBConnProc string = "user=postgres host=/tmp dbname=keerp sslmode=disable password=postgres"
)

// DataBaseAccessor 数据访问器
type DataBaseAccessor struct {
	// *sql.DB
	*gorm.DB
}

// NewDataBaseAccessor 创建数据库访问对象
func NewDataBaseAccessor() *DataBaseAccessor {
	//db, err := sql.Open(DB, DBConn)

	//DBConn := DBConnProc

	if os.Getenv("GODEV") == "0" {
		//DBConn = DBConnDev
	}
	DBConn := DBConnDev

	fmt.Println(DBConn)

	db, err := gorm.Open(DBType, DBConn)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	db.SingularTable(true)

	return &DataBaseAccessor{db}
}
