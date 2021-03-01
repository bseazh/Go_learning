package main

import "fmt"

//set对象的定义和初始化
//set中数据的插入
//从set中查找和读取元素
//从set中删除元素

//分析：set集合实则是堆的实现，查找、插入和删除是n(log)

//定义其Set类
type Set struct {
	num []int
	Map map[int]int
}

//实现底层的 swap
func swap(u *int, v *int) {
	temp := *u
	*u = *v
	*v = temp

}

//交换Set中 num切片中的内容，并同时把hash值一同交换
func Swap_Set(S1 *Set, u int, v int) {
	swap(&S1.num[u], &S1.num[v])
	tmp, _ := S1.Map[u]
	S1.Map[u] = S1.Map[v]
	S1.Map[v] = tmp
}

//上浮	--	堆排序的操作
func Up(S1 *Set, u int) {
	for {
		if u/2 > 0 && S1.num[u] < S1.num[u/2] {
			Swap_Set(S1, u, u/2)
			Down(S1, u)
			u /= 2
		} else {
			return
		}
	}
}

//下沉	--	堆排序的操作
func Down(S1 *Set, x int) {
	u := x
	Size := len(S1.num)
	if x*2 <= Size && S1.num[x*2] < S1.num[u] {
		u = x * 2
	}
	if x*2+1 <= Size && S1.num[x*2+1] < S1.num[u] {
		u = x*2 + 1
	}
	if u != x {
		Swap_Set(S1, u, x)
		Down(S1, u)
	}
}

//初始化Set	--	相当于构造函数
func initialization(S1 *Set) {
	S1.num = append(S1.num, -1)
	S1.Map = make(map[int]int)
	S1.Map[-1] = 0
}

//Set集合的基本操作之	插入
func Insert(s1 *Set, a ...int) {
	a_Len := len(a)
	Size := len(s1.num)
	for i := 0; i < a_Len; i++ {
		ok, _ := s1.Map[a[i]]
		if ok == 0 {
			s1.Map[a[i]] = Size
			s1.num = append(s1.num, a[i])
			Up(s1, Size)
			Size++
		} else {
			continue
		}
	}
}

//Set集合的基本操作之	查找是否存在
func Find(s1 *Set, x int) bool {
	if s1.Map[x] == 0 {
		return false
	} else {
		return true
	}
}

//Set集合的基本操作之	集合中的最值
func MinimumValue(s1 *Set) int {
	return s1.num[1]
}

//Set集合的基本操作之	删除
func Delete(s1 *Set, x int) bool {
	//删除元素时先判断是否存在

	if Find(s1, x) == false {
		return false
	} else {
		//取出堆中待删除元素的 索引
		idx := s1.Map[x]
		fmt.Println("idx**** : ", idx)
		Swap_Set(s1, len(s1.num)-1, idx)
		Swap_Set(s1, idx, 1)
		fmt.Println(s1.num[1:])
		fmt.Println(s1.num[1], "  ", s1.num[idx], " ", s1.num[len(s1.num)-1])
		s1.num = s1.num[:len(s1.num)-1]
		fmt.Println("Length**** : ", len(s1.num))

		Down(s1, 1)
		//Up(s1, idx)

		//取消其映射,使其 x 的索引为0  (合法的索引是[1,Size] )
		s1.Map[x] = 0
		return true
	}
}

func main() {
	var S1 Set
	initialization(&S1)

	Insert(&S1, 3, 15, 1, 10)
	fmt.Println("Start")
	fmt.Println(S1.num[1:])
	fmt.Println(S1.Map)
	fmt.Println("End")

	ok := Delete(&S1, 3)
	fmt.Printf("删除 %v , %t\n", 3, ok)

	fmt.Println(S1.Map)
	fmt.Println(S1.num[:])
}
