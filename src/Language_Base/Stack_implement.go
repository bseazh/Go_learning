package main

import (
	"fmt"
	"math"
)

type IntStack struct {
	num  []int
	p_St int
}

//构造函数:建立初始的栈
func New_IntStack() *IntStack {
	return &IntStack{num: make([]int, 0, 100), p_St: 0}
}

//栈 : 判空
//作用 : 判断栈是否为空
func (this *IntStack) IsEmpty() bool {
	return len(this.num) == 0
}

//栈 : Push
//作用 : 把元素按到栈顶
func (this *IntStack) Push(Elem ...int) {
	for _, val := range Elem {
		this.num = append(this.num, val)
		this.p_St++
	}
}

//栈 : Pop
//作用 : 弹出栈顶元素,并判断是否正确弹出栈顶
func (this *IntStack) Pop() bool {

	if this.IsEmpty() {
		return false
	} else {
		this.p_St--
		this.num = append(this.num[:this.p_St], this.num[this.p_St+1:]...)
		return true
	}
}

//栈 : Top
//作用 : 返回栈顶元素
func (this *IntStack) Top() int {
	if this.IsEmpty() {
		return math.MinInt32
	} else {
		return this.num[this.p_St-1]
	}
}

//栈 : Size
//作用 : 返回栈当前的容量
func (this *IntStack) Size() int {
	return this.p_St
}
func Test_Stack() {

	//Test : New_IntStack
	Stack := New_IntStack()

	//Test : IsEmpty
	if Stack.IsEmpty() {
		fmt.Println("Stack : ", Stack.num, " 为空")
	}

	//Test : Push
	Stack.Push(1, 2, 3, 4, 5)

	//Test : Top & Pop
	for i := 0; i < 3; i++ {
		fmt.Println("弹出元素: ", Stack.Top())
		Stack.Pop()
	}
	//Test : Size
	fmt.Println("栈里剩下的元素个数为: ", Stack.Size())
}
func main() {
	Test_Stack()
}
