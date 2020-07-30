package main

import "fmt"

func main() {
	// 請問以下印出兩次 total 分別為多少?
	// 提示 := 會重新宣告變數

	var total int

	if true {
		total := 10
		fmt.Println("if 內的 total:", total)
	}

	fmt.Println("main 內的 total:", total)
}
