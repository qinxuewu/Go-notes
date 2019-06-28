package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

// 配置文件结构体
type conf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}




func main()  {
	var c conf
	conf:=c.getConf()
	fmt.Println("解析yaml文件： ",conf.Host)

	testToml()
	tesJson()
	testXMl()
}


// 读取并解析配置文件
func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		fmt.Println("【读取配置文件失败：】",err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println("【解析配置文件：】",err.Error())
	}
	return c
}

type MySQlConfig struct {
	Host string
	User string
	Pwd string
	Dbname string
	Friend *Friends
}
type Friends struct {
	Age  int
	Name string
}

// 解析toml格式文件
func testToml() {
	var cp MySQlConfig
	var path string = "conf/conf.toml"
	if _, err := toml.DecodeFile(path, &cp); err != nil {
		log.Fatal(err)
	}

	fmt.Println("解析toml文件", cp.Host, cp.Friend.Name)
}

// 解析json格式文件
func tesJson()  {
	file,_:=os.Open("conf/conf.json")
	defer  file.Close()
	decoder:=json.NewDecoder(file)

	conf := MySQlConfig{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("解析json文件：",conf.Host,conf.User)
}

type XmlConf struct {
	Host string `xml:"host"`
	User string `xml:"user"`
	Pwd string `xml:"pwd"`
	Dbname string `xml:"dbname"`
}

// 解析xml格式文件
func testXMl()  {
	xmlFile,_:=os.Open("conf/conf.xml")
	defer xmlFile.Close()
	var conf XmlConf
	if err := xml.NewDecoder(xmlFile).Decode(&conf); err != nil {
		fmt.Println("Error Decode file:", err)
		return
	}
	fmt.Println("解析xml文件：",conf.Host,conf.User)

}
