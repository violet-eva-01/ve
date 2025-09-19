// Package utils @author: Violet-Eva @date  : 2025/9/19 @notes :
package utils

func ListSplit[T any](input []T, length int) [][]T {

	times := len(input) / length    // 10001 / 2001 = 4
	residual := len(input) % length // 10001 % 2001 = 1997
	if residual > 0 {
		times += 1
	}
	output := make([][]T, times)

	if times <= 1 {
		output[0] = input
	} else {
		starLen := 0
		endLen := length
		for index := 0; index < times; index++ {
			output[index] = input[starLen:endLen]
			starLen += length
			if residual != 0 && index == times-1 {
				endLen += residual
			} else {
				endLen += length
			}
		}
	}
	return output
}
