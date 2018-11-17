package main

// 靴
type Shoes struct {
	Owner *Human
}

func NewShoes(h *Human) *Shoes {
	return &Shoes{
		Owner: h,
	}
}

// 所有者のみ履くことができる
func (s *Shoes) Wear(h *Human) bool {
	return h == s.Owner
}
