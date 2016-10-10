package utils

import (
	"database/sql"
	// Postgre Driver
	_ "github.com/lib/pq"
)

const (
	//DB 数据库类型
	DB string = "postgres"
	// DBConn 连接对象
	DBConn string = "postgres://andysilver:65570480@localhost/keerp?sslmode=disable"
)

// DataBaseAccessor 数据访问器
type DataBaseAccessor struct {
	*sql.DB
}

// NewDataBaseAccessor 创建数据库访问对象
func NewDataBaseAccessor() *DataBaseAccessor {
	db, err := sql.Open(DB, DBConn)

	if err != nil {
		panic(err)
	}

	return &DataBaseAccessor{db}
}
