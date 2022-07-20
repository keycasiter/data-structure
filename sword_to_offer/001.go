package sword_to_offer

// f001 整数除法
func F001(x, y int) int {
	flag := false
	if (x < 0 || y < 0) && !(x < 0 && y < 0) {
		flag = true
	}

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	if y == 0 {
		panic("y can't be 0")
	}
	result := 0
	for x >= y {
		result++
		x = x - y
	}

	if flag {
		return -result
	} else {
		return result
	}
}
