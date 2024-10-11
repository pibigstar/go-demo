package main

import (
	"fmt"

	"github.com/nsf/jsondiff"
)

/**
*  @Author: leikewei
*  @Date: 2024/9/9
*  @Desc:
 */

func main() {
	json1 := `{"a":1,"b":2,"c":3}`
	json2 := `{"dd":1,"b":2,"c":5}`
	opts := jsondiff.DefaultJSONOptions()
	p, str := jsondiff.Compare([]byte(json1), []byte(json2), &opts)
	fmt.Println(str)

	fmt.Println(p.String())
}
