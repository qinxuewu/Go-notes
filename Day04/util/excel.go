package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"strings"
)

// 导出csv
func main()  {
	//f, err := os.Create("/Users/qinxuewu/go/src/we-blog/util/test.csv")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//f.WriteString("\xEF\xBB\xBF")
	//
	//w := csv.NewWriter(f)
	//data := [][]string{
	//	{"1", "test1", "test1-1"},
	//	{"2", "test2", "test2-1"},
	//	{"3", "test3", "test3-1"},
	//}
	//
	//w.WriteAll(data)

	ReadCsv()

}

// 读取csv
func  ReadCsv()  {
	f, err := ioutil.ReadFile("/Users/qinxuewu/go/src/we-blog/util/test.csv")
	if err != nil {
		panic(err)
	}


	r:=csv.NewReader(strings.NewReader(string(f)))
	var records [][]string
  	records,_=r.ReadAll()
	size := len(records)
	for i:=0; i < size; i++{
		log.Print(records[i][0]+","+records[i][1]+","+records[i][2])
	}
}
