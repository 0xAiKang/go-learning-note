package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/**
	获取Mysql 连接
 */
func getMysqlConn() (*sql.DB, error)  {
	db, err := sql.Open("mysql", Conf.Mysql.Dsn)

	if err != nil {
		fmt.Println("get mysql conn error: ", err)
		return nil, err
	}

	return db, nil
}

// 保存中奖记录
func SaveRecords(awardName string, awardTime string, userName string) error {
	db, err := getMysqlConn()
	if err != nil {
		fmt.Println("mysql conn error: ", err)
		return err
	}
	defer db.Close()

	// 预处理
	stmt, err := db.Prepare("insert into award.award_user_info (award_name, award_time, user_name) values (?, ?, ?);")
	if err != nil {
		fmt.Println("prepare insert into sql error: ", err)
		return errors.New("prepare insert sql error")
	}

	fmt.Println("insert into award.award_user_info, award_name, award_time, user_name", awardName, awardTime, userName)

	_, err = stmt.Exec(awardName, awardTime, userName)
	if err != nil {
		fmt.Println("exec sql error: ", err)
		return errors.New("exec sql error")
	}
	return nil
}

/**
	查询中间记录
 */
func QueryRecords() error {
	db, err := getMysqlConn()
	if err != nil {
		fmt.Println("mysql conn error: ", err)
		return err
	}
	defer db.Close()

	rows, err := db.Query("select * from award.award_user_info;")
	if err != nil {
		fmt.Println("exec select error: ", err)
		return err
	}
	// 遍历数据
	for rows.Next() {
		var id int64
		var awardName, awardTime, userName string

		err = rows.Scan(&id, &awardName, awardTime, userName)
		fmt.Println("id : %d, awardName : %s, userName : %s, awardTime : %s \n", id, awardName, userName, awardTime)
	}

	return nil
}