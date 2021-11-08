package main

import "fmt"

func main() {
	score := 75

	if score >= 80 {
		fmt.Println("素晴らしい！")
		if score >= 60 {
			fmt.Println("まあまあ")
		}
	} else {
		fmt.Println("次回は頑張りましょうううううう。")
	}
}
