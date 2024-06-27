package main

import "fmt"

func main() {
	/*
			fmt.Println(3)
			fmt.Println(4.117)
			fmt.Println(false)
			fmt.Println('g')
			fmt.Println("測試")

		var x int
		x = 4
		fmt.Println(x)
		x = 9
		fmt.Println(x)
		var g float64 = 4.7841
		fmt.Println(g)
		var n string = "你好"
		fmt.Println(n)
		var t bool = false
		fmt.Println(t)
		var c rune = 's'
		fmt.Println(c)
	*/
	fmt.Println(3, "lo", 878)
	var x, y, z, s int
	//var y int
	//var z int
	//var s int
	fmt.Print("輸入1個數字：")
	fmt.Scanln(&x, &y, &z)
	fmt.Print("輸入2個數字：")
	fmt.Scanln(&s)
	var result int = x*(y+z) - s
	fmt.Println(result)
}
