package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// AESEncrypt ase encrypt
func AESEncrypt(src []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if src == nil {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(src)
	content = pkcs5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted, nil
}

// AESDecrypt aes decrypt
func AESDecrypt(crypt []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return pkcs5Trimming(decrypted), nil
}

// Base64AndDecrypt as name
func Base64AndDecrypt(base64Crypt, key, iv []byte) ([]byte, error) {
	msgbuf := make([]byte, base64.StdEncoding.EncodedLen(len(base64Crypt)))
	n, _ := base64.StdEncoding.Decode(msgbuf, base64Crypt)
	decrypted, err := AESDecrypt(msgbuf[:n], key, iv)
	return decrypted, err
}

// AESCBCEncrypt AES CBC encrypt with pKCS7Padding
func AESCBCEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if origData == nil {
		fmt.Println("plain content empty")
	}
	blockSize := block.BlockSize()
	origData = pKCS7Padding32Blocksize(origData)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func pKCS7Padding32Blocksize(ciphertext []byte) []byte {
	padding := 32 - len(ciphertext)%32
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7UnPadding32Blocksize(origData []byte) []byte {
	length := len(origData)
	return origData[:(length - int(origData[length-1]))]
}
