package util

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/go-kratos/kratos/v2/log"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (
	pwHashBytes int = 32

	// 当盐值产生错误时，替代的默认值
	defaultSalt string = "1there2is3an4error5here6so7use8default9slat0now20220309084437895"
)

// 创建新的密码(值对象)
// 如果创建过程中出错，会返回nil
func NewPassword(rawPwd string) (string, string, error) {
	salt := MakeSalt()
	encryptedPwd, err := HashPassword(rawPwd, salt)
	if err != nil {
		log.Error("create pwd err: " + err.Error())
		return "", "", err
	}
	return encryptedPwd, salt, nil
}

// 验证密码
func CheckPassword(rawPwd, encryptedPwd, salt string) (bool, error) {
	if rawPwd == "" {
		return false, nil
	} else {
		inputEncryptedPwd, err := HashPassword(rawPwd, salt)
		if err != nil {
			log.Error("check pwd err: " + err.Error())
			return false, err
		}
		return inputEncryptedPwd == encryptedPwd, nil
	}
}

// 更新密码
func UpdatePassword(rawPwd string, salt string) (encryptedPwd string, err error) {
	encryptedPwd, err = HashPassword(rawPwd, salt)
	if err != nil {
		log.Error("update pwd err: " + err.Error())
		return "", err
	}
	return
}

// 密码加密：密码+Salt生成哈希序列
func HashPassword(rawPwd string, salt string) (encryptedPwd string, err error) {
	h, err := scrypt.Key([]byte(rawPwd), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h), nil
}

// 为密码生成Salt值
// 盐值允许异常情况下，使用默认值代替，为了保证密码安全，可以定期替换盐值
func MakeSalt() string {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		log.Error("make salt err: " + err.Error())
		return defaultSalt
	}
	return hex.EncodeToString(buf)
}
