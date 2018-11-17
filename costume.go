package main

// 衣装インタフェース
type Costume interface {
	Wear(h *Human) bool
}

// ドレス
type Dress struct {
	Owner *Human
}

func NewDress(h *Human) Costume {
	return &Dress{
		Owner: h,
	}
}

// 所有者のみ着ることができる
func (d *Dress) Wear(h *Human) bool {
	return h == d.Owner && h.Gender == Woman
}

// 燕尾服
type Tailcoat struct {
	Owner *Human
}

func NewTailcoat(h *Human) Costume {
	return &Tailcoat{
		Owner: h,
	}
}

func (t *Tailcoat) Wear(h *Human) bool {
	return h == t.Owner && h.Gender == Man
}

// ドレスルーム
type DressRoom struct {
	Dresses []*Dress
}

func NewDressRoom() *DressRoom {
	return &DressRoom{}
}

// ドレスを収納する
func (d *DressRoom) Store(humans ...*Human) {
	for _, h := range humans {
		cos := NewDress(h)
		if dress, ok := cos.(*Dress); ok {
			d.Dresses = append(d.Dresses, dress)
		}
	}
}

func (dr *DressRoom) GetDress(h *Human) {
	for _, dress := range dr.Dresses {
		if dress.Wear(h) {
			h.SetCostume(dress)
		}
	}
}
