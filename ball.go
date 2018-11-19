package main

import (
	"fmt"
)

// 舞踏会
type Ball struct {
	Entries    []*Human
	Clock      int // 時刻
	FinishedAt int
}

func NewBall(startedAt, finishedAt int) *Ball {
	return &Ball{
		Clock:      startedAt,  // 開始時刻
		FinishedAt: finishedAt, // 終了時刻
	}
}

func (b Ball) Start() {
	fmt.Println("舞踏会開始")
}

func (b *Ball) Dancing() {
	fmt.Printf("現在 %d時\n", b.Clock)
	for _, h := range b.Entries {
		fmt.Printf("%v は踊っている\n", h.Name)
	}
	b.Clock++
}

func (b Ball) Finish() {
	fmt.Println("舞踏会は終了")
}

func (b Ball) IsFinished() bool {
	return b.Clock >= b.FinishedAt
}

func (b *Ball) Entry(h *Human) {
	if h.Cos != nil {
		b.Entries = append(b.Entries, h)
		fmt.Printf("%v は舞踏会に参加します。\n", h.Name)
	} else {
		fmt.Println("衣装を持っていないと参加できません")
		fmt.Printf("%v は舞踏会に参加できない。\n", h.Name)
	}
}

func (b *Ball) Exit(h *Human) {
	var entries []*Human
	for _, e := range b.Entries {
		if e != h {
			entries = append(entries, e)
		}
	}
	b.Entries = entries
	fmt.Printf("%v は舞踏会から抜け出し、帰宅した。\n", h.Name)
}
