package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type OrderStatus int32

const (
	UnPaid OrderStatus = 1 //待付款
	Cancel OrderStatus = 2 //已取消
	Paid   OrderStatus = 3 //已付款
	Closed OrderStatus = 4 //已关闭
)

func main() {
	db, err := sql.Open("mysql", "root:sferwlwe@(host.docker.internal:3306)/user_db?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		panic(err)
	}
	sql := "select order_sn from tb_orders where order_status in(?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query("1,2,4") //查询待付款，已取消，已关闭
	defer rows.Close()
	for rows.Next() {
		var order_sn string
		if ers := rows.Scan(&order_sn); ers == nil {
			fmt.Println(order_sn)
		}
	}
}
