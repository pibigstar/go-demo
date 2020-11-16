package copy

import (
	"bytes"
	"encoding/gob"
	"github.com/jinzhu/copier"
)

// 基于序列化和反序列化的深度拷贝
func DeepCopy(src, dst interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// 基于反射的深度拷贝
func Copy(src, dst interface{}) error {
	return copier.Copy(dst, src)
}
