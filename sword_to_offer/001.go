package sword_to_offer

// f001 整数除法
func F001(x, y int) int {
	flag := 1
	if x < 0 {
		x = -x
		flag ^= -1
	}
	if y < 0 {
		y = -y
		flag ^= -1
	}

	if y == 0 {
		panic("y can't be 0")
	}
	result := 0
	for x >= y {
		result++
		x = x - y
	}

	if flag == 1 {
		return -result
	} else {
		return result
	}
}
