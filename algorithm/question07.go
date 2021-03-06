package main

import "fmt"

/**
7. 整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	a := 123
	b := 321
	c := 120

	reta := reverse(a)
	retb := reverse(b)
	retc := reverse(c)

	fmt.Printf("the number %d reverse is %d \n", a, reta)
	fmt.Printf("the number %d reverse is %d \n", b, retb)
	fmt.Printf("the number %d reverse is %d \n", c, retc)
}

func reverse(x int) int {
	var ret int = 0

	if x >= 0 {
		for x > 0 {
			ret = ret*10 + x%10
			x = x / 10
		}

		if ret < 0 || ret > 0x7FFFFFFF {
			ret = 0
		}
	} else {
		for x < 0 {
			ret = ret*10 + x%10
			x = x / 10
		}

		if ret > 0 || ret < -0x7FFFFFFF {
			ret = 0
		}
	}

	return ret
}
