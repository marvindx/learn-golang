package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "p"
	port     = 5432
	dbname   = "postgres"
)

// connStr host={host} user={user} password={password} port={port} dbname={dbname} [sslmode=disable]
var connStr = fmt.Sprintf(
	"host=%s user=%s password=%s port=%d dbname=%s sslmode=disable",
	host, user, password, port, dbname,
)

// connStr2 postgres://{user}:{password}@{host}:{port}/{dbname}[?sslmode=disable]
var connStr2 = fmt.Sprintf(
	"postgres://%s:%s@%s:%d/%s?sslmode=disable",
	user, password, host, port, dbname,
)

func Newdb() *sql.DB {
	db, err := sql.Open("postgres", connStr)
	CheckErr(err)
	return db
}

// queryPg 查询
func queryPg(db *sql.DB) {

	rows, _ := db.Query("select * from orders;")
	defer rows.Close()

	for rows.Next() {
		var id int
		//var desc string
		_ = rows.Scan(&id)
		fmt.Println(id)
	}

	//cols, _ := rows.Columns()
	//vals := make([]interface{}, len(cols))
	//for i, _ := range vals {
	//	vals[i] = new(sql.RawBytes)
	//}
	//for rows.Next() {
	//	_ = rows.Scan(vals...)
	//	for j, _ := range vals {
	//		fmt.Printf("%s\n", vals[j].(*sql.RawBytes))
	//	}
	//}
}

// insertPg 插入
func insertPg(db *sql.DB) {

	stmt, err := db.Prepare(`insert into orders("desc") values ($1);`)
	defer stmt.Close()
	CheckErr(err)
	_, err = stmt.Exec("test db")
	CheckErr(err)

	//var lastInsertId int
	//err = db.QueryRow("insert into orders values ($1, $2) returning id;", 31, "test db").Scan(&lastInsertId)
	//CheckErr(err)
	//fmt.Println(lastInsertId)
}

// updatePg 更新
func updatePg(db *sql.DB) {

	stmt, err := db.Prepare(`update orders set "desc"='test' where id = 4;`)
	CheckErr(err)
	defer stmt.Close()
	_, err = stmt.Exec()
	CheckErr(err)
}

// deletePg 删除
func deletePg(db *sql.DB) {
	stmt, err := db.Prepare("delete from orders where id = 5;")
	CheckErr(err)
	defer stmt.Close()
	_, err = stmt.Exec()
	CheckErr(err)
}

// txPg 事务
func txPg(db *sql.DB) {
	tx, err := db.Begin()
	CheckErr(err)
	stmt, err := tx.Prepare(`update orders set "desc"='test tx commit again again' where id = 4;`)
	CheckErr(err)
	defer stmt.Close()
	_, err = stmt.Exec()
	CheckErr(err)
	err = tx.Commit()
	CheckErr(err)
}

func dbDemo() {
	db := Newdb()
	defer db.Close()

	err := db.Ping()
	CheckErr(err)

	//insertPg(db)
	//updatePg(db)
	queryPg(db)
	//deletePg(db)
	//txPg(db)

}

func main() {
	//dbDemo()
}
