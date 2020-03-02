package tools

import (
	"fmt"
	"testing"
)

func TestAppendToFile(t *testing.T) {
	fmt.Println(AppendToFile("/Users/zhangjie/xinbo/myself/go/goTools/test.txt", "\n4、第四行"))
}
