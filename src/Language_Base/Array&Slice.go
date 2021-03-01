package main

import (
	"fmt"
)

//Golang 中 数组和 slice 有什么区别?
//官方文档:https://blog.golang.org/slices-intro

//当 arr , slice 指向同一块连续开辟的区域时:
//1、当 slice := array[ idx1 : idx2 ]
//	实质上是指着同一块区域,修改其中都会影响另外一个.
//
//2、一旦利用切片获取array中部分区域时
//	切片默认设置其 len (切片的大小), cap [从切片结束 -- 数组剩余区域]
//
//3、array是固定开辟大小,从定义起就固定其长度.
//	 相比slice,则可以动态扩充其容量

//切片(slice) 扩充的功能(基本操作)
//	1.定义格式
//	2.添加(某一位置,追加到切片后面)
//	3.删除(某一位置(切片拼接) )
//	4.返回其地址

func main() {
	//相同点:
	//1、当 slice := array[ idx1 : idx2 ]
	//实质上是指着同一块区域,修改其中都会影响另外一个.

	var arr = [8]int{0, 1, 2, 3, 4, 5}
	slice := arr[2:4]
	fmt.Println(" arr: ", arr, " slice : ", slice)
	arr[2] = 6
	fmt.Println(" arr: ", arr, " slice : f", slice)

	fmt.Println("-----------------------------------------")
	//2、一旦利用切片获取其中数组,切片默认设置其 len , cap

	fmt.Printf(" Len : %d \n Cap : %d\n", len(slice), cap(slice))

	fmt.Println("-----------------------------------------")
	//3、slice支持追加功能 类似 realloc 的功能
	//容量方面：一旦超过其容量可以两倍扩充
	//如下例子,原本长度为2，增加7个元素后 , 9 > cap(6)
	//                                    cap(12) <- 12 ( 6 * 2 )
	//地址方面：管理区域
	//同时换了一个首地址,即拷贝原来的切片到另外一个分配区域中重新再添加这7个元素.
	//(&slice) = arr[2] :0xc0000b6010 -> (&slice) = 0xc00008c060

	fmt.Printf(" array : %p \n Slice : %p \n", &arr[2], slice)
	add_Slice := []int{1, 2, 3, 4, 5, 6, 7}
	slice = append(slice, add_Slice...)
	fmt.Printf(" array : %p \n Slice : %p \n Len : %d \n Cap : %d \n",
		&arr, slice, len(slice), cap(slice))

	fmt.Println("-----------------------------------------")

	//定义 利用make( type , len , cap )
	var s = make([]int, 3, 5)
	s = slice[:]
	fmt.Println(" s: ", s)

	fmt.Println("-----------------------------------------")

	//添加	(add)
	//1.(后面追加)
	//func append(s []T, x ...T) []T
	s = append(s, 1, 2, 3, 4)
	// 如果 s2 = [1,2,3,4]
	// 格式 s = append(s, s2... ) "..." 相当于拆包
	fmt.Println(s)

	fmt.Println("-----------------------------------------")

	//2.(中间插入)
	s = []int{1, 2, 4, 5}
	s = append(s[:2], append([]int{3}, s[2:]...)...)
	fmt.Println(s)
	fmt.Println("-----------------------------------------")

	//删除（实现切片拼接）
	s = []int{1, 2, 6, 3, 4}
	s = append(s[:2], s[3:]...)
	fmt.Println(s)

	fmt.Println("-----------------------------------------")

	//返回切片的地址
	fmt.Printf("%p\n", s)
	fmt.Println(&s)

}
