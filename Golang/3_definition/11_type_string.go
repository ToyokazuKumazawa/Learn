package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello World"[0])) //72 アスキー文字コードが表示される.byte配列.文字列で表示したいときはstring()でキャストする
	var s string = "Hello World"

	s = strings.Replace(s, "H", "X", 1) // HがあったらXに初めの1個置き換える
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println(`Test
                       Test
Test`)

	fmt.Println("\"") // "を表示したいときは\"を"で囲むか
	fmt.Println(`"`)  // `で"を囲む
}
