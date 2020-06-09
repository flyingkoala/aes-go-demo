/*
@Time : 2020/6/9 16:52
@Author : wkang
@File : main
@Description:
*/
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
)

func main() {
	//设置key
	key := []byte("PsGC9MSijtRAMbYD4RmHbjR&eZXG#Z8_")
	realkey:= key[0:16]
	log.Println("key值：",hex.EncodeToString(realkey))
	//明文
	origData := []byte("需要加密的内容")
	//加密
	en := AESEncrypt(origData,realkey)
	log.Println("加密结果：",hex.EncodeToString(en))
	//解密
	de := AESDecrypt(en,realkey)
	log.Println("解密结果：",string(de))
}
//解密
func AESDecrypt(crypted,key []byte)[]byte{
	block,_ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block,key[:blockSize])
	origData := make([]byte,len(crypted))
	blockMode.CryptBlocks(origData,crypted)
	origData = PKCS7UnPadding(origData)
	return origData
}

//去补码
func PKCS7UnPadding(origData []byte)[]byte{
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}
//加密
func AESEncrypt(origData,key []byte)[]byte{
	//获取block块
	block,_ :=aes.NewCipher(key)
	//补码
	origData = PKCS7Padding(origData, block.BlockSize())
	//加密模式，
	blockMode := cipher.NewCBCEncrypter(block,key[:block.BlockSize()])
	//创建明文长度的数组
	crypted := make([]byte,len(origData))
	//加密明文
	blockMode.CryptBlocks(crypted,origData)
	return crypted
}
//补码
func PKCS7Padding(origData []byte,blockSize int)[]byte{
	//计算需要补几位数
	padding := blockSize-len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(origData,padtext...)
}