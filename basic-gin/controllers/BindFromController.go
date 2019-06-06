package controllers

import "github.com/gin-gonic/gin"

// 构造表单请求的参数
type StructA struct {
	FieldA string `form:"a"`  // form 后的参数对应表单中输入框name属性
}
type StructB struct {
	NestedStruct StructA
	FieldB string `form:"b"`
}
type StructC struct {
	NestedStructPointer *StructA
	FieldC string `form:"c"`
}
type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"x"`
	}
	FieldD string `form:"d"`
}

// 获取表单b的参数
func GetDataB(c *gin.Context)  {
	var b StructB
	c.Bind(&b)
	c.JSON(200,gin.H{"a":b.NestedStruct,"b":b.FieldB})
}

func GetDataC(this *gin.Context)  {
	var c StructC
	this.Bind(&c)
	this.JSON(200,gin.H{"a":c.NestedStructPointer,"b":c.FieldC})
}

func GetDataD(c *gin.Context) {
	var d StructD
	c.Bind(&d)
	c.JSON(200, gin.H{
		"x": d.NestedAnonyStruct,
		"d": d.FieldD,
	})
}



