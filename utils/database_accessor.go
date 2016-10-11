package utils

import (

	// Postgre Driver
	"github.com/jinzhu/gorm"
	// GORM for PostgreSQL
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	//DBType  数据库类型
	DBType string = "postgres"
	// DBConn 连接对象
	// DBConn string = "postgres://andysilver:65570480@localhost/keerp?sslmode=disable"
	DBConn string = "host=localhost user=andysilver dbname=keerp sslmode=disable password=65570480"
)

// DataBaseAccessor 数据访问器
type DataBaseAccessor struct {
	// *sql.DB
	*gorm.DB
}

// NewDataBaseAccessor 创建数据库访问对象
func NewDataBaseAccessor() *DataBaseAccessor {
	//db, err := sql.Open(DB, DBConn)

	db, err := gorm.Open(DBType, DBConn)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	db.SingularTable(true)

	return &DataBaseAccessor{db}
}
