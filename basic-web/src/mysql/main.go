
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main()  {
	db,err:=sql.Open("mysql","root:870439570@tcp(39.108.144.143:3306)/test2?charset=utf8")
	checkErr(err)
	
	//插入数据
	stmt, err := db.Prepare("INSERT INTO user SET name=?")
	checkErr(err)

	res, err := stmt.Exec("qxw")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id:",id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	//返回受影响数的行
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect:",affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()



}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
