package bit

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestBit(t *testing.T) {
	//& 与运算
	//清零
	assert.Equal(t, 0, And(3, 0))
	assert.Equal(t, 0, And(12, 0))

	//将某几位清零，高位或低位，上面的变种使用
	assert.Equal(t, 5, And(245, 15))

	//判断奇偶
	assert.Equal(t, isOdd(1), true)
	assert.Equal(t, isOdd(0), false)
	assert.Equal(t, isEven(1), false)
	assert.Equal(t, isEven(0), true)

	//| 或运算

	//^ 异或运算

	//~ 取反

}
