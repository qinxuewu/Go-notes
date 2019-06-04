package main

type StructA struct {
	FieldA string `form:"field_a"`
}
type StructB struct {
	NestedStruct StructA
	FieldB string `form:"field_b"`
}
type StructC struct {
	NestedStructPointer *StructA
	FieldC string `form:"field_c"`
}
type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}



// 使用自定义结构绑定表单数据请求
func main()  {
	
}
