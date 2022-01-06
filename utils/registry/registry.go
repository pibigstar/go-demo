// +build windows

package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

// 操作注册表
func main() {
	// 创建：指定路径的项
	// 路径：HKEY_CURRENT_USER\Software\Test
	key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\Test`, registry.ALL_ACCESS)
	defer func() {
		_ = key.Close()
	}()

	// 判断是否已经存在了
	if exists {
		fmt.Println(`键已存在`)
	} else {
		fmt.Println(`新建注册表键`)
	}

	// 写入：32位整形值
	_ = key.SetDWordValue(`32位整形值`, uint32(123456))
	// 写入：64位整形值
	_ = key.SetQWordValue(`64位整形值`, uint64(123456))
	// 写入：字符串
	_ = key.SetStringValue(`字符串`, `hello`)
	// 写入：字符串数组
	_ = key.SetStringsValue(`字符串数组`, []string{`hello`, `world`})
	// 写入：二进制
	_ = key.SetBinaryValue(`二进制`, []byte{0x11, 0x22})

	// 读取：字符串
	s, _, _ := key.GetStringValue(`字符串`)
	fmt.Println(s)

	// 读取：一个项下的所有子项
	keys, _ := key.ReadSubKeyNames(0)
	for _, key := range keys {
		// 输出所有子项的名字
		fmt.Println(key)
	}

	// 创建：子项
	subKey, _, _ := registry.CreateKey(key, `子项`, registry.ALL_ACCESS)
	defer func() {
		_ = subKey.Close()
	}()

	// 删除：子项
	// 该键有子项，所以会删除失败
	// 没有子项，删除成功
	err := registry.DeleteKey(key, `子项`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
