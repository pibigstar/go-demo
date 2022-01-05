// +build windows

package disk

import (
	"fmt"
	"strconv"
	"syscall"

	"golang.org/x/sys/windows/registry"
)

// 获取系统中所有盘符
func GetSystemDisks() []string {
	// 获取系统dll
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	// 获取dll中函数
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	// 调用dll中函数
	n, _, _ := GetLogicalDrives.Call()
	s := strconv.FormatInt(int64(n), 2)
	var allDrives = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:",
		"I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：",
		"U：", "V：", "W：", "X：", "Y：", "Z："}
	temp := allDrives[0:len(s)]
	var d []string
	for i, v := range s {
		if v == 49 {
			l := len(s) - i - 1
			d = append(d, temp[l])
		}
	}
	var drives []string
	for i, v := range d {
		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	}
	return drives
}

// 获取插入的U盘盘符
func GetUDisk() []string {
	//查询注册表，判断是否插入U盘
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\USBSTOR\Enum`, registry.QUERY_VALUE)
	if err != nil {
		fmt.Println("Not have U-Disk")
		return nil
	}
	defer k.Close()
	// 获取注册表中值，得到插入了几个U盘
	count, _, err := k.GetIntegerValue("Count")
	// 获取全部盘符
	disks := GetSystemDisks()

	return disks[len(disks)-int(count):]
}
