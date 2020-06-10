package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
	logger "github.com/takashabe/bazel-exercise/lib/go"
)

type InnoDBStats struct {
	DatabaseName         string      `json:"database_name" db:"database_name"`
	TableName            string      `json:"table_name" db:"table_name"`
	LastUpdate           interface{} `json:"last_update" db:"last_update"`
	NRows                int         `json:"n_rows" db:"n_rows"`
	ClusteredIndexSize   int         `json:"clustered_index_size" db:"clustered_index_size"`
	SumOfOtherIndexSizes int         `json:"sum_of_other_index_sizes" db:"sum_of_other_index_sizes"`
}

func main() {
	if err := DBConn(); err != nil {
		panic(err)
	}
}

func DBConn() error {
	db, err := sqlx.Connect("mysql", "root@/mysql")
	if err != nil {
		return err
	}
	defer db.Close()

	stats := []*InnoDBStats{}
	if err := db.Select(&stats, "select * from innodb_table_stats"); err != nil {
		return err
	}
	pp.Println(stats)

	return nil
}

func A() string {
	return "A"
}

func B() {
	logger.Info("%s", "hoge")
}
