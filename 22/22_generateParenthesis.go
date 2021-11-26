package main

import "fmt"

// 求n个括号的所有组合, 如2时: ()()  (())
/*
	思路: 动态规划, n个括号时可能的情况为
		在已经有一个括号的情况下, 去增加 n-1 一个括号, 即是n个括号了
		新加n-1个括号如下情况
			括号内部加i个 右边加j个 满足 i+j=n-1
			为什么没有左边加k个? 这与右边加j个重复
		0的时候为 空
		1的时候为 ()
		这样我们就可以从2个括号情况开始递推到n
		左右两边i/j个的可能情况会出现在之前枚举过的情况
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	var resI [][]string                 // 枚举 0-n个括号 的所有结果
	resI = append(resI, []string{""})   // i为0时表示空
	resI = append(resI, []string{"()"}) // i为1时只有一种可能
	for i := 2; i <= n; i++ {           // 递推第i个情况
		var tempRes []string
		for j := 0; j < i; j++ { // j表示括号内部个数 那么右边括号个数肯定是 i-1-j
			leftRes := resI[j]
			rightRes := resI[i-1-j]
			for _, l := range leftRes { // 遍历左右情况进行组合
				for _, r := range rightRes {
					tempRes = append(tempRes, "("+l+")"+r)
				}
			}
		}
		resI = append(resI, tempRes)
	}
	return resI[n]
}

func main() {
	fmt.Println(generateParenthesis(5))
}
