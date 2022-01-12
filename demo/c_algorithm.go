package demo

import (
	U128 "github.com/mengzhuo/uint128"
)

var (
	ZERO, _ = U128.NewFromString("0")
	ONE, _  = U128.NewFromString("1")
	Cnt     *U128.Uint128
)

func Solution(i, j int) int {
	Cnt, _ = U128.NewFromString("0")
	return plain(i, j)
}

func plain(i, j int) int {
	if i < j || i < 0 || j < 0 {
		return 0
	}
	if j == 0 {
		return 1
	}
	Cnt.Add(ONE)
	// fmt.Println(i,",",j)
	return plain(i-1, j) + plain(i-1, j-1)
}

// ========= 备忘录实现方式 ==============
var (
	m   [][]int          // 备忘录
	cnt [][]U128.Uint128 // 递归计数
)

func SolutionMemo(i, j int) int {
	Cnt, _ = U128.NewFromString("0")
	m = make([][]int, i+1)
	cnt = make([][]U128.Uint128, i+1)

	for i := range m {
		m[i] = make([]int, j+1)
		cnt[i] = make([]U128.Uint128, j+1)
	}
	res, _ := memo(i, j)

	// for idx := range cnt {
	// 	fmt.Println("\t",cnt[idx])
	// }

	Cnt = &(cnt[i][j])
	return res
}

func memo(i, j int) (int, *U128.Uint128) {
	if i < j || i < 0 || j < 0 {
		return 0, ZERO
	}
	if j == 0 {
		return 1, ZERO
	}

	if m[i][j] == 0 {
		r1, c1 := memo(i-1, j)
		r2, c2 := memo(i-1, j-1)

		m[i][j] = r1 + r2

		cnt[i][j].Add(c1)
		cnt[i][j].Add(c2)
		cnt[i][j].Add(ONE)
	}

	return m[i][j], &cnt[i][j]
}
