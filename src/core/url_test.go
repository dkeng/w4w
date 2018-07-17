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
