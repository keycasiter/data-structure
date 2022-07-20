package sword_to_offer

import (
	"fmt"
	"testing"
)

// 问题1：不能使用乘法、除法、求余
//       通过循环减法累加来求除法的商
// 问题2：除数与被除数相差巨大，循环减法效率低下
//       每次减去除数的n次方
// 问题3：商的正负数判断
//       正正=正 负负=正 正负=负
// 问题4：整数范围越界

func Test001(t *testing.T) {
	fmt.Println(F001(-11, -2))
}
