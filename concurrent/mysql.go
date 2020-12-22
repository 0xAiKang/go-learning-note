package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Mysql 使用实例
func getConn() (*sql.DB) {
	// 建立链接
	db, err := sql.Open("mysql", "root:122410@tcp(127.0.0.1:3306)/award?charset=utf8")
	if err != nil {
		fmt.Println("connect fail")
		return nil
	}

	return db
}

func main()  {
	db := getConn()
	if db == nil {
		log.Println("conn is nil")
		return
	}

	// 插入数据（预处理）
	stmt, err := db.Prepare("insert into award_user_info(award_name, award_time, user_name) values (?, ?, ?);")
	if err != nil {
		log.Println("prepare insert sql error:", err)
	}
	awardName := "boo"
	awardTime := "2020-12-22 23:40:55"
	userName := "boo"
	log.Println("insert into award_user_info , award_name , award_time  , user_name ",awardName, awardTime, userName)

	_, err = stmt.Exec(awardName, awardTime, userName)
	if err != nil {
		log.Println("exec sql error:", err)
		return
	}

	// 修改数据（执行删除）
	result, err := db.Exec("update award_user_info set award_name = \"mac\" where award_name = \"boo\";")
	if err != nil {
		log.Println("exec sql error:", err)
		return
	}
	fmt.Println(result)

	// 删除数据（执行删除）
	_, err = db.Exec("delete from award_user_info where award_name = \"boo\";")

	// 查询数据
	rows, err := db.Query("select * from award_user_info")
	if err != nil {
		log.Println("exec select error", err)
		return
	}

	for rows.Next() {
		var id int
		var award_name, user_name, award_time string

		err = rows.Scan(&id, &award_name, &user_name, &award_time)
		log.Printf("id: %d, award_name: %s, user_name: %s, award_time: %s \n", id, award_name, user_name, award_time )
	}

	// 关闭数据库连接
	defer db.Close()
}