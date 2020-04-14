package main

import (
	"log"
	"os"

	// 变量初始化 _ 编译器在导入未被引用的包时不报错
	_ "./matchers"
	"./search"
)

// init在main之前调用 // 程序里所有的 init 方法都会在 main 函数启动前被调用
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定的项做搜索
	search.Run("president")
}
