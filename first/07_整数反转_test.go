package first

import (
	"fmt"
	"math"
	"testing"
)

//int32 -2147483648 to 2147483647
//我们先假设 rev为正数,负数同理
//如果 temp=rev⋅10+pop 导致溢出，那么一定有 rev≥ INTMAX/10
//如果rev>INTMAX/10,那么temp=rev⋅10+pop一定会溢出
//如果rev==INTMAX/10,那么只要pop>7,temp=rev⋅10+pop就会溢出
func reverse(x int) int {
	var rev int
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
func TestReverse(t *testing.T) {

	out1 := reverse(123)
	out2 := reverse(-123)
	out3 := reverse(120)
	fmt.Println(out1, " ", out2, " ", out3)
}

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
// 示例 1:
// 	输入: 123
// 	输出: 321
// 示例 2:
// 	输入: -123
// 	输出: -321
// 示例 3:
// 	输入: 120
// 	输出: 21
// 注意:
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
