package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

const key = "1111111111111111"

func AesEncryptByCBC(data string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	// 填充
	paddText := PKCS7Padding([]byte(data), block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, []byte(key)[:block.BlockSize()])

	// 加密
	result := make([]byte, len(paddText))
	blockMode.CryptBlocks(result, paddText)
	// 返回密文
	return base64.StdEncoding.EncodeToString(result)
}

func AesDecryptByCBC(data string) string {
	decodeString, err := base64.StdEncoding.DecodeString(data)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(key)[:block.BlockSize()])

	result := make([]byte, len(decodeString))

	blockMode.CryptBlocks(result, []byte(decodeString))
	// 去除填充
	result = UnPKCS7Padding(result)
	return string(result)
}

func UnPKCS7Padding(text []byte) []byte {
	// 取出填充的数据 以此来获得填充数据长度
	unPadding := int(text[len(text)-1])
	return text[:(len(text) - unPadding)]
}

func PKCS7Padding(text []byte, blockSize int) []byte {
	// 计算待填充的长度
	padding := blockSize - len(text)%blockSize
	var paddingText []byte
	if padding == 0 {
		// 已对齐，填充一整块数据，每个数据为 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		// 未对齐 填充 padding 个数据，每个数据为 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}
