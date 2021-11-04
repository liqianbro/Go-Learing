package utils_test

import (
	"Go-Learing/utils"
	"fmt"
	"testing"
)

func TestHideStar(t *testing.T) {
	idCard := "320482200007086917"
	fmt.Println(utils.HideStar(idCard))
	name := "龚智超"
	fmt.Println(utils.HideStar(name))
	em := "12121212121@qq.com"
	fmt.Println(utils.HideStar(em))
}
