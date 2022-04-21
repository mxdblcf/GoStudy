package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// 进行读取文件
//思路，for检测 /media 口 ，定时查看是否有U盘插入，  如果有的话进入

func Decrypt(encrypted string, key []byte) (string, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}

//读取tmp文件,并解码
func myreadfile(s string) string {
	//拿到一个路径数组
	//fmt.Println(s)
	userFile := s + "/.tmp001"    //文件路径
	fin, err := os.Open(userFile) //打开文件,返回File的内存地址
	defer fin.Close()             //延迟关闭资源
	if err != nil {
		log.Println("i/o错误")
		fmt.Println(userFile, err)
		return "err"
	}
	var str string
	buf := make([]byte, 1024) //创建一个初始容量为1024的slice,作为缓冲容器
	for {
		//循环读取文件数据到缓冲容器中,返回读取到的个数
		n, _ := fin.Read(buf)
		if 0 == n {
			break //如果读到个数为0,则读取完毕,跳出循环
		}
		//从缓冲slice中写出数据,从slice下标0到n,
		bytes := buf[:n]
		str += string(bytes)
	}

	decodeString, err := hex.DecodeString(str)
	encodeToString := hex.EncodeToString(decodeString)
	key := []byte{0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
		0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
		0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
		0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
	}
	//转码
	strDecrypted, err := Decrypt(encodeToString, key)
	if err != nil {
		log.Println("Decrypted err:", err)
	}
	fmt.Println("Decrypted:", strDecrypted)
	return strDecrypted
}
func main() {

	//strDecrypted,err := Decrypt(strEncrypted, key)

	//	fmt.Println("Decrypted:",strDecrypted)

	//定时拿文件解码,能拿到就结束工作，拿不到一直循环
	for {
		time.Sleep(1 * time.Second)
		//判断目录下有没有文件

		glob, _ := filepath.Glob(filepath.Join("/media", "*/*"))

		if glob != nil {
			s := glob[0]
			s2 := myreadfile(s)
			if s2 != "" {
				break
			}
		}
	}

}
