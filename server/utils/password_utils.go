package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)
func VerifyPassword(inputPwd, storedPwd string) bool {
	fmt.Println(inputPwd)
	// 1. 第一次加密：md5(pwd + "_" + pwd)
	data := []byte(inputPwd + "_" + inputPwd)
	first := md5.Sum(data)
	firstHex := hex.EncodeToString(first[:])

	// 2. 第二次加密：md5(firstHex + "_xf_2222_" + pwd)
	secondData := []byte(firstHex + "_xf_2222_" + inputPwd)
	second := md5.Sum(secondData)
	secondHex := hex.EncodeToString(second[:])

	// 3. 组合结果
	hashedInput := firstHex + "$$" + secondHex

	// 4. 比较结果
	fmt.Println(hashedInput, storedPwd)
	return hashedInput == storedPwd
}