package main

import "fmt"

type Magic struct {
	Target *Human
	Broken chan int
}

func NewMagic(h *Human) *Magic {
	return &Magic{
		Target: h,
		Broken: make(chan int, 1),
	}
}

func (m Magic) GenerateDress() Costume {
	fmt.Printf("%v は魔法でドレスを作ってもらった！\n", m.Target.Name)
	return NewDress(m.Target)
}
func (m Magic) GenerateGlassShoes() *Shoes {
	fmt.Printf("%v は魔法でガラスの靴を作ってもらった！\n", m.Target.Name)
	return NewShoes(m.Target)
}

func (m *Magic) Limit(limit chan int) {
	<-limit
	fmt.Println("0時が近づいている")
	fmt.Println("魔法が解けそう！！")
	m.Broken <- 1
}
