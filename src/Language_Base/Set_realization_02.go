package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//定义类型
//参考: https://www.cnblogs.com/wanghui-garcia/p/10581388.html

//map[int]struct{} :
// 		其含义为只关心key,不需要非得需要映射的value值,强行变成set

type IntSet struct {
	num map[int]struct{}
}

//构造函数,New
//复杂度 : O( 1 )
func New_IntSet() *IntSet {
	return &IntSet{
		num: make(map[int]struct{}),
	}
}

//集合 中 当前元素数量( Size )
//实质 是 Map 存在key的个数
//复杂度 : O( map.Size ) =? O( N )
func (this *IntSet) setSize() int {
	return len(this.num)
}

//集合 中 添加元素(Element)
//实质 是 Map 添加Key值,对应的value值为空
//复杂度 : O( map->key.insert ) =? O( logN )
func (this *IntSet) Insert(Elem ...int) *IntSet {
	for _, elem := range Elem {
		this.num[elem] = struct{}{}
	}
	return this
}

//集合 中 判断元素(Element)是否存在
//实质 是 利用Map 映射返回结果来判断是否存在
//格式: value , (T|F) = map[key]
//复杂度 : O( map->key.search ) =? O( logN )
func (this *IntSet) Find(Elem int) bool {
	_, exist := this.num[Elem]
	return exist
}

//集合 中 删除元素(Element)
//实质 是 利用Map 删除key-value对的操作
//格式: delete( map , key )
//复杂度 : O( map.delete ) =? O( logN )
func (this *IntSet) Delete(Elem int) {
	delete(this.num, Elem)
}

//集合 中 清空整个集合
//实质 是 重新赋值一个空的map
//复杂度 : O( 1 )
func (this *IntSet) Clear() {
	this.num = make(map[int]struct{})
}

//集合序列化
//实质 是 把Map转化为切片
//复杂度 : O( S )
func (this *IntSet) ToSlice() []int {
	cnt_Set := len(this.num)
	if cnt_Set == 0 {
		return []int{}
	}
	res := make([]int, cnt_Set)
	i := 0
	// range map => 遍历其key值
	for Elem := range this.num {
		res[i] = Elem
		i++
	}
	return res
}

//集合 的 序列化后转化为字符串打印出来(方便调试)
//实质 是 利用strconv.Itoa()进行转化并拼接

func (this *IntSet) ToString() string {
	slice := this.ToSlice()
	res_String := ""
	for _, elem := range slice {
		res_String = res_String + " " + strconv.Itoa(elem)
	}
	return res_String
}

//集合 的 判断是否为子集 : s1 是否为 s2 的子集
//实质是遍历s1若发现其中有元素不是s2的元素则 不为s2的子集
//复杂度 : O( S )
func SubsetOf(s1, s2 IntSet) bool {
	if s1.setSize() > s2.setSize() {
		return false
	}
	for key, _ := range s1.num {
		if s2.Find(key) == false {
			return false
		}
	}
	return true
}

//集合 的 交集运算(Intersection)
//实质: 遍历集合s1每个元素,若发现其元素同样也在s2则保存下来
//复杂度: O( S )
func (s1 *IntSet) Intersection(s2 IntSet) *IntSet {
	res := New_IntSet()
	for key, _ := range s1.num {
		if s2.Find(key) == true {
			res.Insert(key)
		}
	}
	return res
}

//集合 的 并集运算(Union)
//实质: 遍历集合s1,s2的所有元素保存下来,注意去重
//复杂度: O( S )
func (s1 *IntSet) Union(s2 IntSet) *IntSet {
	res := New_IntSet()
	for key, _ := range s1.num {
		res.Insert(key)
	}
	for key, _ := range s2.num {
		res.Insert(key)
	}
	return res
}

//集合 的 差集运算()
//实质: 遍历集合s1每个元素,把不存在s2的保存下来
//复杂度: O( S )
func (s1 *IntSet) Difference(s2 IntSet) *IntSet {
	res := New_IntSet()
	for key, _ := range s1.num {
		if s2.Find(key) == false {
			res.Insert(key)
		}
	}
	return res
}

//测试函数
func Test_Set() {

	//Test : New_IntSet
	set := New_IntSet()

	//随机生成 6 个随机数 , Test : Insert
	var randNum int
	for i := 0; i < 6; i++ {
		randNum = rand.Int() % 100
		set.Insert(randNum)
		//fmt.Println(randNum)
	}

	fmt.Println("Set集合中的全部元素: ", set.num)

	//Test : Find & Delete
	if set.Find(randNum) == true {
		fmt.Println("被删除的最后一个元素是: ", randNum)
		set.Delete(randNum)
	}

	//Test : ToSlice
	var set_Slice []int = set.ToSlice()
	fmt.Println("序列化后: ", set_Slice)

	s1, s2, s3 := New_IntSet(), New_IntSet(), New_IntSet()
	s1.Insert(1, 2, 3)
	s2.Insert(1, 2, 3, 4, 5)
	s3.Insert(3, 4, 5, 6, 7)
	//Test: SubsetOf

	fmt.Println(" s1:", s1.ToString(), "\n",
		"s2:", s2.ToString(), "\n",
		"s3:", s3.ToString())

	if SubsetOf(*s1, *s2) {
		fmt.Println("集合s1: ", s1.ToString(), " 是 ", "集合s2: ", s2.ToString(), "的子集")
	} else {
		fmt.Println("集合s1: ", s1.ToString(), " 不是 ", "集合s2: ", s2.ToString(), "的子集")
	}

	//Test: Intersection
	fmt.Println("s1与s2的交集", (s1.Intersection(*s2)).ToString())
	fmt.Println("s2与s3的交集", (s2.Intersection(*s3)).ToString())

	//Test: Union
	fmt.Println("s1与s2的并集", (s1.Union(*s2)).ToString())
	fmt.Println("s2与s3的并集", (s2.Union(*s3)).ToString())

	//Test: Difference
	fmt.Println("s1与s2的差集", (s1.Difference(*s2)).ToString())
	fmt.Println("s2与s3的差集", (s2.Difference(*s3)).ToString())
}

//主函数直接调用 Test_Set()
func main() {
	Test_Set()
}
