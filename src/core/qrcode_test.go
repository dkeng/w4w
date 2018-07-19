package core_test

import (
	"fmt"
	"testing"

	"github.com/dkeng/w4w/src/core"
)

func TestCreateQrcodeBase64(t *testing.T) {
	inputURL := "http://www.baidu.com"
	base64, err := core.CreateQrcodeBase64(inputURL)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(base64)
}
