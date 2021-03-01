package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
type teacher struct {
	Name  string `json:"Miss"`
	Phone string `json:"Phone"`
}

//反射,两种方式获取字段
//1.根据索引依次遍历其字段信息
//2.根据字段名称来搜索相应的字段的信息
//

func Reflect_Struct_Field(x interface{}) {

	fmt.Println("--------------------")
	v := reflect.TypeOf(x)
	fmt.Println(v.Name(), "\t", v.Kind())

	//利用循环把结构体中的每一个字段进行遍历
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Println("Name:", field.Name, "\tType:", field.Type, "\tTag:", field.Tag)
		fmt.Println(field.Tag.Get("json"))
	}

	//根据字段名进行搜索
	if nameField, ok := v.FieldByName("Name"); ok {
		fmt.Printf("name:%s\tindex:%d\ttype:%v\tjson tag:%v\n", nameField.Name, nameField.Index, nameField.Type, nameField.Tag.Get("json"))
	}
}

//测试reflect 结构体获取其字段信息
func Test_Reflect_Struct_Field() {
	stu1 := student{Name: "Lesile", Score: 88}
	teacher1 := teacher{Name: "Miss Alam", Phone: "13458293814"}
	Reflect_Struct_Field(stu1)
	Reflect_Struct_Field(teacher1)
}

// 用于测试的两个函数

func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

// 根据结构体中的内容 , 依次获取 TypeOf , ValueOf , Kind
//	1.	方法的个数需要 从 TypeOf 下的 NumMethod() 方法获取
//  2.	方法的名称需要 从 TypeOf 下的 Method(i).Name 获取
//  3.	方法的返回类型需要 从 ValueOf 下的 Method(i).Type() 获取
//  4.  调用方法需要 在 ValueOf 下的 Call( args )
//      其中 args 为参数 需要提前定义 reflect.Value{}类型

func Reflect_Struct_Method(x interface{}) {

	fmt.Println("--------------------")
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	k := v.Kind()
	fmt.Println("Type : ", t, "\tKind : ", k, "\tvalue : ", v)

	//方法的个数需要 从 TypeOf 下的 NumMethod() 方法获取
	NumMethod := t.NumMethod()
	fmt.Println(NumMethod)

	for i := 0; i < NumMethod; i++ {
		//方法的名称需要 从 TypeOf 下的 Method(i).Name 获取
		fmt.Println("Name: ", t.Method(i).Name)

		//方法的返回类型需要 从 ValueOf 下的 Method(i).Type() 获取
		fmt.Printf("Type: %s\n", v.Method(i).Type())

		//args 为参数 需要提前定义 reflect.Value{}类型
		var args = []reflect.Value{}

		//调用方法需要 在 ValueOf 下的 Call( args )
		v.Method(i).Call(args)
	}
}

func Test_Reflect_Struct_Method() {
	stu1 := student{Name: "Lesile", Score: 88}
	Reflect_Struct_Method(stu1)

}
func main() {
	Test_Reflect_Struct_Field()
	Test_Reflect_Struct_Method()
}
