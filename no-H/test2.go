package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 写入.tmp文件
//思路，for检测 /media 口 ，定时查看是否有U盘插入，  如果有的话进入

func Encrypt(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息

	if err != nil {
		if os.IsExist(err) {

			return true

		}
		return false
	}

	return true

}

//输入一串id，进行 加密

func main() {
	args := os.Args                   //获取用户输入的所有参数
	if args == nil || len(args) < 1 { //如果用户没有输入,或参数个数不够,则调用该函数提示用户
		fmt.Println("your str?")

		return
	}
	str := args[1] //获取输入的第一个参数
	key := []byte{0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
		0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
		0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
		0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
	}
	strEncrypted, err := Encrypt(str, key)
	if err != nil {
		log.Println("Encrypted err:", err)
	}
	fmt.Println("Encrypted:", strEncrypted)
	//拿到文件路径
	glob, _ := filepath.Glob(filepath.Join("/media", "*/*"))
	s := glob[0]
	fmt.Println(s)
	_, err = os.Stat(s + "/.tmp001") //os.Stat获取文件信息
	path := s + "/.tmp001"
	if Exists(path) {
		err := os.Remove(path)
		if err != nil {
			return
		}

	}
	create, err := os.Create(path)
	defer func(create *os.File) {
		err := create.Close()
		if err != nil {

		}
	}(create)
	_, err = create.WriteString(strEncrypted)
	if err != nil {
		log.Println(err)
		return
	}
}
