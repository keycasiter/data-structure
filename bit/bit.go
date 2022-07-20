package bit

import "fmt"

// & 与运算 （AND）
// | 或运算 （OR）
// ^ 异或运算 （XOR）
// ~ 取反运算 （NOT）
// >> 右移运算 （RIGHT SHIFT）
// << 左移运算 （LEFT SHIFT）

func And(x, y int) int {
	fmt.Printf("%d & %d = %d\n", x, y, x&y)
	fmt.Printf("%.8b\n%.8b\n%.8b\n", x, y, x&y)
	return x & y
}

func Or(x, y int) int {
	return x | y
}

func Xor(x, y int) int {
	return x ^ y
}

func LeftShift(x, idx int) int {
	fmt.Printf("%d << %d = %d\n", x, idx, x<<idx)
	fmt.Printf("%.8b\n%d\n%.8b\n", x, idx, x<<idx)
	return x << idx
}

func RightShift(x, idx int) int {
	fmt.Printf("%d >> %d = %d\n", x, idx, x>>idx)
	fmt.Printf("%.8b\n%d\n%.8b\n", x, idx, x>>idx)
	return x >> idx
}

//判断奇数
func isOdd(x int) bool {
	return And(x, 1) == 1
}

//判断偶数
func isEven(x int) bool {
	return And(x, 1) != 1
}
