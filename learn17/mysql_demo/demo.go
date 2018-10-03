package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type User struct {
	Id   int            `db:"id"`
	Name sql.NullString `db:"name"`
}

func queryRowTest() {
	sql := "select id, name from test_data where id=?"
	row := Db.QueryRow(sql, 1)
	var user User
	err := row.Scan(&user.Id, &user.Name) // row.Scan() 后，才会释放连接。所以必须Scan！
	if err != nil {
		fmt.Println("scan row failed, err:", err)
		return
	}
	fmt.Printf("user: %v \n", user)
}

func queryAllTest() {
	sql := "select id, name from test_data"
	rows, err := Db.Query(sql)
	if err != nil {
		fmt.Println("query all failed, err:", err)
		return
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name) // row.Scan() 后，才会释放连接。所以必须Scan！
		if err != nil {
			fmt.Println("scan row failed, err:", err)
			return
		}
		fmt.Printf("user: %v \n", user)
	}
}

func insertTest() {
	sql := "insert test_data(id, name) values(?, ?)"
	res, err := Db.Exec(sql, 2, "lisi")
	if err != nil {
		fmt.Println("insert data failed, err:", err)
		return
	}
	id, err := res.LastInsertId()
	count, err := res.RowsAffected()
	fmt.Printf("插入的行数: %d，插入id：%d \n", count, id)
}

func updateTest() {
	sql := "update test_data set name=? where id=?"
	res, err := Db.Exec(sql, "lisi02", 2)
	if err != nil {
		fmt.Println("insert data failed, err:", err)
		return
	}
	id, err := res.LastInsertId()
	count, err := res.RowsAffected()
	fmt.Printf("插入的行数: %d，插入id：%d \n", count, id)
}

func deleteTest() {
	sql := "delete from test_data where id=?"
	res, err := Db.Exec(sql, 2)
	if err != nil {
		fmt.Println("insert data failed, err:", err)
		return
	}
	id, err := res.LastInsertId()
	count, err := res.RowsAffected()
	fmt.Printf("插入的行数: %d，插入id：%d \n", count, id)
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Println("init db failed, err:", err)
		return
	}
	// queryRowTest()
	// insertTest()
	// updateTest()
	deleteTest()
	queryAllTest()
}

func initDb() error {
	var err error
	dsn := "root:123@tcp(localhost:3306)/caipiao_admin"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open mysql failed, err:", err)
		return err
	}
	Db.SetMaxOpenConns(100) // 打开最大连接数
	Db.SetMaxIdleConns(16)  // 空闲连接数

	return nil
}
