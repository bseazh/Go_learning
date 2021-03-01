package main

import (
	"fmt"
	"math"
	"strconv"
)

type IntQueue struct {
	num  []int
	rear int
}

//构造函数 : 构造出Queue
func New_IntQueue() *IntQueue {
	return &IntQueue{make([]int, 0, 100), 0}
}

//判断队伍是否为空(Is_Empty)
func (this *IntQueue) Is_Empty() bool {
	return this.rear == 0
}

//返回队首元素(Front)
func (this *IntQueue) Front() int {
	return this.num[0]
}

//返回队尾元素(Rear)
func (this *IntQueue) Rear() int {
	return this.num[this.rear-1]
}

//返回队伍元素个数(Size)
func (this *IntQueue) Size() int {
	return this.rear
}

//入队(EnQueue)
func (this *IntQueue) EnQueue(Elem ...int) {
	this.num = append(this.num, Elem...)
	this.rear += len(Elem)
}

//出队(DeQueue)
func (this *IntQueue) DeQueue() int {
	if this.Is_Empty() {
		fmt.Printf("Queue为空 \n")
		return math.MinInt32
	}
	this.num = append(this.num[:0], this.num[1:]...)
	this.rear--
	return this.Front()
}

//队列序列化打印(ToString)
//调试用的
func (this *IntQueue) ToString() string {
	res := ""
	for _, val := range this.num {
		res = res + " " + strconv.Itoa(val)
	}
	return res
}

//测试IntQueue
func Test_queue() {
	//Test : New_IntQueue
	Q := New_IntQueue()

	//Test : Is_Empty
	if Q.Is_Empty() {
		fmt.Println("队列Q", Q.num, "为空")
	}
	//Test : EnQueue
	Q.EnQueue(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(Q.ToString())

	//Test : Front
	fmt.Println("队首元素为: ", Q.Front())

	//Test : Rear
	fmt.Println("队尾元素为: ", Q.Rear())

	//Test : DeQueue
	for i := 0; i < 3; i++ {
		Q.DeQueue()
	}

	//Test : Size
	fmt.Println(Q.ToString())
	fmt.Println("队伍中的元素个数: ", Q.Size())
}
func main() {
	Test_queue()
}
