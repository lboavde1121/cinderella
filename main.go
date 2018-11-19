package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Cinderella start ===")

	stepMother := NewHuman("StepMother", 52, Woman)
	sisterA := NewHuman("SisterA", 23, Woman)
	sisterB := NewHuman("SisterB", 20, Woman)
	cinderella := NewHuman("ella", 18, Woman)

	// シンデレラいじめられる
	stepMother.Say("今日もいじめてやるw")
	sisterA.Say("今日もいじめてやるw")
	sisterB.Say("今日もいじめてやるw")
	cinderella.Say("・・・")

	// 舞踏会が開催されます。
	// 19時開始
	ball := NewBall(19, 27)
	// 舞踏会用のドレスを用意します。
	dressRoom := NewDressRoom()
	dressRoom.Store(stepMother, sisterA, sisterB)

	// シンデレラのドレスは用意しません。
	// 継母のドレスはある
	dressRoom.GetDress(stepMother)
	// 舞踏会の開催
	ball.Entry(stepMother)
	// 姉Aのドレスもある
	dressRoom.GetDress(sisterA)
	ball.Entry(sisterA)
	// 姉Bのドレスもある
	dressRoom.GetDress(sisterB)
	ball.Entry(sisterB)

	// シンデレラだけドレスがない。。
	dressRoom.GetDress(cinderella)
	ball.Entry(cinderella)

	// 舞踏会に行きたがるシンデレラを、不可思議な力（魔法使い、仙女、ネズミ、母親の形見の木、白鳩など）が助け、準備を整える
	magic := NewMagic(cinderella)
	cinderella.SetCostume(magic.GenerateDress())
	cinderella.SetShoes(magic.GenerateGlassShoes())
	limit := make(chan int, 1)
	go magic.Limit(limit)
	// シンデレラ舞踏会へ参加する
	ball.Entry(cinderella)

	// 王子登場
	prince := NewHuman("王子", 18, Man)
	tailcoat := NewTailcoat(prince)
	prince.SetCostume(tailcoat)
	ball.Entry(prince)

	// 舞踏会開催
	ball.Start()
	finished := make(chan int, 1)
	go func() {
		for !ball.IsFinished() {
			<-time.After(1 * time.Second)
			ball.Dancing()

			if ball.Clock == 24 {
				limit <- 1
			}
		}
		ball.Finish()
		finished <- 1
	}()

	<-magic.Broken
	// シンデレラいそいで帰る
	ball.Exit(cinderella)
	// ガラスの靴を落としてしまう！！
	falledShoes := cinderella.Shoes
	cinderella.Shoes = nil

	// シンデレラ靴を落とす→王子靴を拾う
	foundShoes := falledShoes

	// 舞踏会終了
	<-finished

	fmt.Println("=== 後日 ===")
	fmt.Println("王子の部下たちはガラスの靴の持ち主を探している。")

	// 靴の持ち主を舞踏会の参加者の中から探す
	for _, h := range ball.Entries {
		if h.Gender == Woman {
			if foundShoes.Wear(h) {
				fmt.Println("見つけた！")
			} else {
				fmt.Printf("%v: %vさんの靴ではない\n", prince.Name, h.Name)
			}
		}
	}
	if foundShoes.Wear(cinderella) {
		fmt.Println("見つけた!")
	}
}
