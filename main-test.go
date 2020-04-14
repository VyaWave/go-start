package main

import "fmt"

func init() {
	fmt.Println("Hello, World!  GO init....")
}

var g int = 99

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World! GO finish....")

	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)

	if ptr != nil {
		fmt.Printf("不是空指针")

	} /* ptr 不是空指针 */
	if ptr == nil {
		fmt.Printf("是空指针\n")

	} /* ptr 是空指针 */
}
