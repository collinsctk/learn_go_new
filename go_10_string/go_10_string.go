package main

import "fmt"

func main() {
	peace := "peace"
	var var_peace = "peace"
	var var_peace_string string = "peace"
	var blank string

	fmt.Println(
		peace,
		var_peace,
		blank,
		var_peace_string)

	// -------------------字符串字面值/原始字符串字面值-------------------
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence\nlike this`)

	// ---------------------Unicode---------------------
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	//---------------------%v表示值---------------------
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	//---------------------%c表示字符---------------------
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	// --------------------字符串是不可变的--------------------
	message := "qytang"
	c := message[5]
	fmt.Println(c)

	//message[5] = 'd' // 不能这样做
	//fmt.Println(message)

	// --------------------遍历字符串--------------------
	// --------------------方案1--------------------
	for i := 0; i < len(message); i++ {
		c := message[i]
		fmt.Println(c)
		//在Go语言中，当你使用字符串索引访问字符串中的字符时，得到的实际上是该字符的Unicode代码点（即rune值）。这就是为什么你看到的是数字，而不是字符。
		//在Go语言中，字符串是不可变的字节序列。访问字符串的索引会返回字节值，而不是字符。这些字节值是字符的UTF-8编码。因此，当你打印这些值时，它们显示为整数值。
		fmt.Printf("%c\n", c)
	}

	// --------------------方案2--------------------
	for i, c := range message {
		fmt.Printf("index: %d, character: %c\n", i, c)
	}
	// --------------------Go内置函数---------------------
	fmt.Println(len(message))

}
