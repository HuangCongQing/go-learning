/*
 * @Description: 多态 https://www.yuque.com/huangzhongqing/lang/qso8oc#wrvna
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2023-08-20 19:25:23
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/10-OOP/test4_interface.go
 */

 package main

import "fmt"

//本质是一个指针<<<<<<所以下面调用的时候要写&xxx
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色  string代表一个返回值
	GetType() string  //获取动物的种类
}

//具体的类1
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类2   Dog
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// 传参选择是哪个具体的类
func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {

	var animal AnimalIF //接口的数据类型， 父类指针
	animal = &Cat{"Green"}

	animal.Sleep() //调用的就是Cat的Sleep()方法 , 多态的现象

	animal = &Dog{"Yellow"}

	animal.Sleep() // 调用Dog的Sleep方法，多态的现象

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)
}
