/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          utils.go
 * Description:   util
 */

package weigo

import (
	"fmt"
	"reflect"
)

func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func debugPrintln(message ...interface{}) {
	fmt.Println(message)
}

func debugTypeof(element interface{}) interface{} {
	return reflect.TypeOf(element)
}

func debugCheckError(err error) {
	if err != nil {
		debugPrintln(err)
	}
}
