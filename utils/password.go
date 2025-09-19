// Package utils @author: Violet-Eva @date  : 2025/9/19 @notes :
package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// NewPassword
// @Description:
// @param limits 0: digits, 1: lowers, 2:uppers, 3:chars
// @return str
// @return err
func NewPassword(limits ...[4]int) (str string, err error) {

	var (
		limit [4]int
		sb    strings.Builder
	)
	if len(limits) <= 0 {
		limit = [4]int{1, 1, 1, 1}
	} else {
		limit = limits[0]
	}
	rand.NewSource(time.Now().UnixNano())
	digits := []byte("0123456789")
	lowers := []byte("abcdefghijklmnopqrstuvwxyz")
	uppers := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	chars := []byte(",.<>!@#$%^&*()_=-[]{}|;:/?")
	byteS := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	passwordLength := 18 + rand.Intn(6)
	sb.WriteString(fmt.Sprintf("密码长度为: %d", passwordLength))
	leftPasswordLength := passwordLength - limit[0] - limit[1] - limit[2] - limit[3]
	if leftPasswordLength < 0 {
		err = fmt.Errorf("密码限制为:[%d]位数字,[%d]位小写字母,[%d]位大写字母.已超过密码的长度[%d],请重新指定密码限制", limit[0], limit[1], limit[2], passwordLength)
		return
	}
	var result []byte
	sb.WriteString(fmt.Sprintf("至少取[%d]位数字", limit[0]))
	for i := 0; i < limit[0]; i++ {
		result = append(result, byteS[rand.Intn(len(digits))])
	}
	sb.WriteString(fmt.Sprintf("至少取[%d]位小写字母", limit[1]))
	for i := 0; i < limit[1]; i++ {
		result = append(result, byteS[rand.Intn(len(lowers))])
	}
	sb.WriteString(fmt.Sprintf("至少取[%d]位大写字母", limit[2]))
	for i := 0; i < limit[2]; i++ {
		result = append(result, byteS[rand.Intn(len(uppers))])
	}
	sb.WriteString(fmt.Sprintf("至少取[%d]位特殊字符", limit[3]))
	for i := 0; i < limit[2]; i++ {
		result = append(result, byteS[rand.Intn(len(chars))])
	}
	rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < leftPasswordLength; i++ {
		result = append(result, byteS[rand.Intn(len(byteS))])
	}
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	str = string(result)
	fmt.Println(sb.String())
	return
}
