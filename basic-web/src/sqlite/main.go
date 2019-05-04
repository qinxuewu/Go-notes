
package main
import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"time"
)
func main()  {

	db,err:=sql.Open("sqlite3","./foo.db")
	checkErr(err)

	 //插入数据
	   stmt,err:=db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	   checkErr(err)

	   res,err:=stmt.Exec("秦学武","技术部","2012-12-09")

	   id,err:=res.LastInsertId();
	   checkErr(err)

	   fmt.Println("id:",id)

		// 更新数据
	   stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	   checkErr(err)

	   res, err = stmt.Exec("秦学武6666", id)
	   checkErr(err)

	   //返回受影响数的行
	   affect, err := res.RowsAffected()
	   checkErr(err)
	   fmt.Println(affect)

      //查询数据
      rows, err := db.Query("SELECT * FROM userinfo")
      checkErr(err)

	  for rows.Next() {
			var uid int
			var username string
			var department string
			var created time.Time
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

func checkErr(err error)  {
	if err !=nil{
		panic(err)
	}
}
