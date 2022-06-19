package main

import "fmt"

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	fmt.Println(c1(3))

	c2 := circleArea(3)
	fmt.Println(c2(2))
	fmt.Println(c2(3))
}

// クロージャ：ネストされた関数定義において、内側の関数が外にある変数を使って処理をすることができるというもの
// 関数内で状態（値）を保持し続けて、処理をさせたい場合などにクロージャは有用

// 無名関数
// 一度しか使わない関数であれば名前を考える手間が省ける
// 他との名前の衝突を防ぐことができる
// 関数が使うプログラムの近くにあると処理の見通しが良くなる
