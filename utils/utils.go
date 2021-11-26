package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Password(plainPwd string) string {
	//谷歌的加密包
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encodePWD
}

// CheckPassword 密码校验
func CheckPassword(plain, cropped string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cropped), []byte(plain)) //验证（对比）
	return err == nil
}
