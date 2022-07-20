package sword_to_offer

// f001 整数除法
func F001(x, y int) int {
	//处理符号
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

	//有效性校验
	if y == 0 {
		panic("y can't be 0")
	}
	//累积商
	result := 0
	for x >= y {
		result++
		//优化 TODO
		x = x - y
	}

	//处理符号
	if flag {
		return -result
	} else {
		return result
	}
}
