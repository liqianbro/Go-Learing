package crypto

import (
	"fmt"
	"testing"
)

func TestCrypto(t *testing.T) {
	text := "1111111111111111"
	result := AesEncryptByCBC(text)
	fmt.Println(result)
	fmt.Println(AesDecryptByCBC(result))
}
