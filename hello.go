package main

import "fmt"

//import "rsc.io/quote"
import "unsafe"

func main() {
	var x = 1
	x += x
	fmt.Println(x)
	fmt.Println("hello work")

	var b = "hello"
	fmt.Println(len(b))
	fmt.Println(unsafe.Sizeof(b))

	if Unknown == 0 {
		fmt.Println("unknow is zero")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//fmt.Println("i=" + string(i))
		str := "i="
		res := fmt.Sprintf("%s%d", str, i)
		fmt.Println(res)
	}

	fmt.Println(Female)

	fmt.Println(Hello("hellen"))

	a, b := swap("Google", "Runoob")

	fmt.Println(a, b)

	var c int = 10
	fmt.Println("变量的地址: %x\n", &c)

	//fmt.Println(quote.Go())

	var d int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &d /* 指针变量的存储地址 */

	fmt.Printf("d 变量的地址是: %x\n", &d)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// go空指针
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	fmt.Printf("ptr 的地址为 : %x\n", &ptr)

	//m := make(map[string]int)
	// 使用字面量创建 Map
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	// 获取键值对
	v1 := m["apple"]
	v2, ok := m["pear"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
	fmt.Println(v1)
	fmt.Println(v2, ok)
	// 遍历 Map
	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	fmt.Println(Factorial(4));
}

func swap(x, y string) (string, string) {
	return y, x
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)
