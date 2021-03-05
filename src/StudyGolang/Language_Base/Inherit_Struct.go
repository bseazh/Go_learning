package main

import "fmt"

//参考网站 : https://golangbot.com/learn-golang-series/
//结构体 如何实现结构体的方法继承

//举一个例子,父类是人
//子类进行实现继承 并拓展自身的属性
type Person struct {
	name string
}

func (this *Person) Self_introduction() {
	fmt.Println("My name is ", this.name)
}

//学生 在 "人"的基础上 增加 学号
//其中Person类作为父类,继承到Student类
//name属性(隐式)镶嵌到Student
type Student struct {
	Person
	ID_stu string
}

//重载函数
func (this *Student) Self_introduction() {
	fmt.Println("My name is ", this.name)
	fmt.Println("My ID_stu is ", this.ID_stu)
}

//老师 在 "人"的基础上 增加 工号 , 手机号
//其中Person类作为父类,继承到Teacher类
//name属性(隐式)镶嵌到Teacher
type Teacher struct {
	Person
	ID_Tea string
	Phone  string
}

//重载函数
func (this *Teacher) Self_introduction() {
	fmt.Println("My name is ", this.name)
	fmt.Println("My ID_stu is ", this.ID_Tea)
	fmt.Println("My Phone number is ", this.Phone)
}

func main() {

	Ben := Person{name: "罗翔"}
	Ben_stu := Student{Person: Ben, ID_stu: "S0001"}
	Ben_Tea := Teacher{Person: Ben, ID_Tea: "T0001", Phone: "13057295781"}

	Ben.Self_introduction()
	Ben_stu.Self_introduction()
	Ben_Tea.Self_introduction()

}
