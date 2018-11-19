package main

import "fmt"

type Gender int

const (
	_ Gender = iota
	Woman
	Man
)

// 投稿人物用
type Human struct {
	Name   string
	Age    int
	Gender Gender
	Cos    Costume
	Shoes  *Shoes
}

func NewHuman(n string, age int, g Gender) *Human {
	return &Human{
		Name:   n,
		Age:    age,
		Gender: g,
	}
}

func (h Human) Say(s string) {
	fmt.Printf("%v: %v\n", h.Name, s)
}

func (h *Human) SetCostume(c Costume) {
	h.Cos = c
}
func (h *Human) SetShoes(s *Shoes) {
	h.Shoes = s
}
