package core_test

import (
	"fmt"
	"testing"

	"github.com/dkeng/w4w/src/core"
)

func TestShort(t *testing.T) {
	resultURL, err := core.ShortURL("http://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resultURL)
}

func TestFormatURL(t *testing.T) {
	inputURL := "http://www.baidu.com"
	fmt.Println(core.FormatURL(inputURL))
}

func TestCheckURL(t *testing.T) {
	inputURL := "https://www.baidu.com"
	fmt.Println(core.CheckURL(inputURL))
}
