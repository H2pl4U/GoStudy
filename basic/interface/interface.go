package main

import "fmt"

//会叫的动物
type sayer interface {
	say()
}

//会动的动物
type mover interface {
	move()
}

type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (d dog) move() {
	fmt.Println("狗会动")
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	var x sayer
	a := dog{}
	b := cat{}
	x = a
	x.say()
	x = b
	x.say()

	var y mover
	y = a
	c := &dog{}
	y = c
	y.move()
}
