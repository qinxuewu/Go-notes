package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

//  https://www.jb51.net/article/147071.htm
var db *sql.DB

func init() {
	log.Println("init............")
	db,_=sql.Open("mysql","root:870439570@tcp(39.108.144.143:3306)/test2?charset=utf8")
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)
	db.Ping()

}

func Db() sql.DB {
	db,_:=sql.Open("mysql","root:870439570@tcp(39.108.144.143:3306)/test2?charset=utf8")
	return  *db
}

func main() {
	startHttpServer()
}

func startHttpServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("【 http f服务启动成功 】")
}

func pool(w http.ResponseWriter, r *http.Request) {
	stmt,_:=db.Prepare("INSERT INTO user SET name=?")
	res, err := stmt.Exec("astaxie")
	//defer db.Close()

	checkErr2(err)
	fmt.Fprintln(w, "finish",res)


}

func checkErr2(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}