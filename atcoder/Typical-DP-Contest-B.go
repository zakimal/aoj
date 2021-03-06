package main

import (
	"fmt"
)

const SIZE = 1005

func main() {
	var A, B int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	a := make([]int, A+1)
	b := make([]int, B+1)
	for i := 1; i <= A; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 1; i <= B; i++ {
		fmt.Scanf("%d", &b[i])
	}
	reverse(&a, A)
	reverse(&b, B)
	// sente[i][j]: 先手のターンで左にi個右にj個残っている状態になってからゲーム終了までに先手が得られる価値の最大値
	sente := make([][]int, A+1)
	for i := range sente {
		sente[i] = make([]int, B+1)
	}
	// gote[i][j]: 後手のターンで左にi個右にj個残っている状態になってからゲーム終了までに先手が得られる価値の最大値
	gote := make([][]int, A+1)
	for i := range gote {
		gote[i] = make([]int, B+1)
	}
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			if i == 0 && j == 0 {
				// 先手のターンで左に0個右に0個残っている状態になってからゲーム終了までに先手が得られる価値の最大値は0
				sente[i][j] = 0

				// 後手のターンで左に0個右に0個残っている状態になってからゲーム終了までに先手が得られる価値の最大値は0
				gote[i][j] = 0
			} else if i == 0 {
				// 先手のターンで左に0個右にj個残っている状態になってから
				// ゲーム終了までに先手が得られる価値の最大値は
				// 「一つ前の（後手の）ターン（左に0個右にj-1個残っている状態）になってからゲーム終了までに先手が得られる価値の最大値」
				// +「そのターンでとる価値」
				sente[i][j] = b[j] + gote[0][j-1]

				// 後手のターンで左に0個右にj個残っている状態になってから
				// ゲーム終了までに先手が得られる価値の最大値は
				// 「一つ前の（先手の）ターン（左に0個右にj-1個残っている状態）になってからゲーム終了までに先手が得られる価値の最大値」
				// （後手のターンでは先手は得点を伸ばすことはできないから。）
				gote[i][j] = sente[0][j-1]
			} else if j == 0 {
				// 先手のターンで左にi個右に0個残っている状態になってから
				// ゲーム終了までに先手が得られる価値の最大値は
				// 「一つ前の（後手の）ターン（左にi-1個右に0個残っている状態）になってからゲーム終了までに先手が得られる価値の最大値」
				// +「そのターンでとる価値」
				sente[i][j] = a[i] + gote[i-1][0]

				// 後手のターンで左にi個右に0個残っている状態になってから
				// ゲーム終了までに先手が得られる価値の最大値は
				// 「一つ前の（先手の）ターン（左にi-1個右に0個残っている状態）になってからゲーム終了までに先手が得られる価値の最大値」
				// （後手のターンでは先手は得点を伸ばすことはできないから。）
				gote[i][j] = sente[i-1][0]
			} else {
				// 先手のターンで左にi個右にj個残っている状態になってから
				// ゲーム終了までに先手が得られる価値の最大値は
				// 「一つ前の（後手の）ターン（左にi-1個右にj個残っている状態 or 左にi個右にj-1個残っている状態）になってから
				// ゲーム終了までに先手が得られる価値の最大値」 +「そのターンでとる価値」の和のうち大きい方
				// （先手は最善の手を取るので大きい方）
				sente[i][j] = max(a[i]+gote[i-1][j], b[j]+gote[i][j-1])

				// 後手のターンで左にi個右にj個残っている状態になってから
				// ゲーム終了までに先手が得られる価値の最大値は
				// 「一つ前の（先手の）ターン（左にi-1個右にj個残っている状態 or 左にi個右にj-1個残っている状態）になってから
				// ゲーム終了までに先手が得られる価値の最大値」のうち小さい方
				// （後手は最善の手を取るので先手から見れば最悪の手を取ることになるので小さい方）
				gote[i][j] = min(sente[i-1][j], sente[i][j-1])
			}
		}
	}
	// 求めたいのは「先手のターンで左にi個右にj個残っている状態になってからゲーム終了までに先手が得られる価値の最大値」
	// つまりsente[A][B]
	fmt.Println(sente[A][B])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverse(a *[]int, l int) {
	for i, j := 1, l; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}
