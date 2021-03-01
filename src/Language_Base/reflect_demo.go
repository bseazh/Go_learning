package main

import (
	"fmt"
	"reflect"
)

// Round 1.
// reflect.TypeOf()
//	接收 interface{} 空接口 , 任何类型都为空接口的子类
//	打印其 Name , Type
//	Name 是类的名称 , 种类是 系统内置 + struct

func reflect_Type(x interface{}) {
	obj := reflect.TypeOf(x)
	fmt.Println(obj, "\tName:", obj.Name(), "\tKind:", obj.Kind())
}

type Cat struct{}
type Dog struct{}

func Test_reflect_Type() {
	var a int32 = 10
	var b float32 = 1.23
	var c Cat
	var d Dog
	var e []int
	var f []string
	reflect_Type(e)
	reflect_Type(f)
	reflect_Type(a)
	reflect_Type(b)
	reflect_Type(c)
	reflect_Type(d)
}

// Round 2.
// reflect.ValueOf()
//	打印其 Type , 以及对应Value

func reflect_ValueOf(x interface{}) {
	v := reflect.ValueOf(x)
	//fmt.Printf("%v %T\n", v, v)
	k := v.Kind()
	//fmt.Println(k)

	switch k {
	case reflect.Int32:
		res := int32(v.Int())
		fmt.Printf("%v , %T\n", res, res)
	case reflect.Float32:
		res := float32(v.Float())
		fmt.Printf("%v , %T\n", res, res)
	}
}

func Test_reflect_ValueOf() {
	a := int32(10)
	b := float32(1.234)
	reflect_ValueOf(a)
	reflect_ValueOf(b)
}

//Round 3
// reflect.SetInt()
// 根据具体的类型进行设置
func reflect_SetValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()

	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)
	}
}

func Test_reflect_SetValue() {
	var a int32 = 10
	reflect_SetValue(&a)

	var b float32 = 1.23
	reflect_SetValue(&b)

	fmt.Println(a, b)
}

func main() {
	Test_reflect_Type()
	Test_reflect_ValueOf()
	Test_reflect_SetValue()
}
