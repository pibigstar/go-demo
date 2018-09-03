/*
  Copyright (c) 2012-2013 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
	A helper package for converting between datatypes.

	If a certain value can not be directly converted to another, the zero value
	of the destination type is returned instead.
*/
package to

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

var (
	durationType = reflect.TypeOf(time.Duration(0))
	timeType     = reflect.TypeOf(time.Time{})
)

const (
	digits     = "0123456789"
	uintbuflen = 20
)

const (
	KindTime reflect.Kind = iota + 1000000000
	KindDuration
)

var strToTimeFormats = []string{
	"2006-01-02 15:04:05 Z0700 MST",
	"2006-01-02 15:04:05 Z07:00 MST",
	"2006-01-02 15:04:05 Z0700 -0700",
	"Mon Jan _2 15:04:05 -0700 MST 2006",
	time.RFC822Z, // "02 Jan 06 15:04 -0700"
	time.RFC3339, // "2006-01-02T15:04:05Z07:00", RFC3339Nano
	"2006-01-02 15:04:05 -0700",
	"2006-01-02 15:04:05 Z07:00",
	time.RubyDate, // "Mon Jan 02 15:04:05 -0700 2006"
	time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700"
	time.RFC822,   // "02 Jan 06 15:04 MST",
	"2006-01-02 15:04:05 MST",
	time.UnixDate, // "Mon Jan _2 15:04:05 MST 2006",
	time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST",
	time.RFC850,   // "Monday, 02-Jan-06 15:04:05 MST",
	time.Kitchen,  // "3:04PM"
	"01/02/06",
	"2006-01-02",
	"2006/01/02",
	"01/02/2006",
	"Jan _2, 2006",
	"01/02/06 15:04",
	time.Stamp, // "Jan _2 15:04:05", time.StampMilli, time.StampMicro, time.StampNano,
	time.ANSIC, // "Mon Jan _2 15:04:05 2006"
	"2006-01-02 15:04",
	"2006-01-02T15:04",
	"01/02/2006 15:04",
	"01/02/06 15:04:05",
	"01/02/2006 15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05",
	"_2/Jan/2006 15:04:05",
}

var strToDurationMatches = map[*regexp.Regexp]func([][][]byte) (time.Duration, error){
	regexp.MustCompile(`^(\-?\d+):(\d+)$`): func(m [][][]byte) (time.Duration, error) {
		sign := 1

		hrs := time.Hour * time.Duration(Int64(m[0][1]))

		if hrs < 0 {
			hrs = -1 * hrs
			sign = -1
		}

		min := time.Minute * time.Duration(Int64(m[0][2]))

		return time.Duration(sign) * (hrs + min), nil
	},
	regexp.MustCompile(`^(\-?\d+):(\d+):(\d+)$`): func(m [][][]byte) (time.Duration, error) {
		sign := 1

		hrs := time.Hour * time.Duration(Int64(m[0][1]))

		if hrs < 0 {
			hrs = -1 * hrs
			sign = -1
		}

		min := time.Minute * time.Duration(Int64(m[0][2]))
		sec := time.Second * time.Duration(Int64(m[0][3]))

		return time.Duration(sign) * (hrs + min + sec), nil
	},
	regexp.MustCompile(`^(\-?\d+):(\d+):(\d+).(\d+)$`): func(m [][][]byte) (time.Duration, error) {
		sign := 1

		hrs := time.Hour * time.Duration(Int64(m[0][1]))

		if hrs < 0 {
			hrs = -1 * hrs
			sign = -1
		}

		min := time.Minute * time.Duration(Int64(m[0][2]))
		sec := time.Second * time.Duration(Int64(m[0][3]))
		lst := m[0][4]

		for len(lst) < 9 {
			lst = append(lst, '0')
		}
		lst = lst[0:9]

		return time.Duration(sign) * (hrs + min + sec + time.Duration(Int64(lst))), nil
	},
}

func strToDuration(v string) time.Duration {

	var err error
	var d time.Duration

	d, err = time.ParseDuration(v)

	if err == nil {
		return d
	}

	b := []byte(v)

	for re, fn := range strToDurationMatches {
		m := re.FindAllSubmatch(b, -1)
		if m != nil {
			r, err := fn(m)
			if err == nil {
				return r
			}
		}
	}

	return time.Duration(0)
}

func uint64ToBytes(v uint64) []byte {
	buf := make([]byte, uintbuflen)

	i := len(buf)

	for v >= 10 {
		i--
		buf[i] = digits[v%10]
		v = v / 10
	}

	i--
	buf[i] = digits[v%10]

	return buf[i:]
}

func int64ToBytes(v int64) []byte {
	negative := false

	if v < 0 {
		negative = true
		v = -v
	}

	uv := uint64(v)

	buf := uint64ToBytes(uv)

	if negative {
		buf2 := []byte{'-'}
		buf2 = append(buf2, buf...)
		return buf2
	}

	return buf
}

func float32ToBytes(v float32) []byte {
	slice := strconv.AppendFloat(nil, float64(v), 'g', -1, 32)
	return slice
}

func float64ToBytes(v float64) []byte {
	slice := strconv.AppendFloat(nil, v, 'g', -1, 64)
	return slice
}

func complex128ToBytes(v complex128) []byte {
	buf := []byte{'('}

	r := strconv.AppendFloat(buf, real(v), 'g', -1, 64)

	im := imag(v)
	if im >= 0 {
		buf = append(r, '+')
	} else {
		buf = r
	}

	i := strconv.AppendFloat(buf, im, 'g', -1, 64)

	buf = append(i, []byte{'i', ')'}...)

	return buf
}

/*
	Converts a date string into a time.Time value, several date formats are tried.
*/
func Time(val interface{}) time.Time {
	switch t := val.(type) {
	// We could use this later.
	default:
		s := String(t)
		for _, format := range strToTimeFormats {
			r, err := time.ParseInLocation(format, s, time.Local)
			if err == nil {
				return r
			}
		}
	}
	return time.Time{}
}

/*
	Tries to convert the argument into a time.Duration value. Returns
	time.Duration(0) if any error occurs.
*/
func Duration(val interface{}) time.Duration {
	switch t := val.(type) {
	case int:
		return time.Duration(int64(t))
	case int8:
		return time.Duration(int64(t))
	case int16:
		return time.Duration(int64(t))
	case int32:
		return time.Duration(int64(t))
	case int64:
		return time.Duration(t)
	case uint:
		return time.Duration(int64(t))
	case uint8:
		return time.Duration(int64(t))
	case uint16:
		return time.Duration(int64(t))
	case uint32:
		return time.Duration(int64(t))
	case uint64:
		return time.Duration(int64(t))
	default:
		return strToDuration(String(val))
	}
	panic("Reached")
}

/*
	Tries to convert the argument into a []byte array. Returns []byte{} if any
	error occurs.
*/
func Bytes(val interface{}) []byte {

	if val == nil {
		return []byte{}
	}

	switch t := val.(type) {

	case int:
		return int64ToBytes(int64(t))

	case int8:
		return int64ToBytes(int64(t))
	case int16:
		return int64ToBytes(int64(t))
	case int32:
		return int64ToBytes(int64(t))
	case int64:
		return int64ToBytes(int64(t))

	case uint:
		return uint64ToBytes(uint64(t))
	case uint8:
		return uint64ToBytes(uint64(t))
	case uint16:
		return uint64ToBytes(uint64(t))
	case uint32:
		return uint64ToBytes(uint64(t))
	case uint64:
		return uint64ToBytes(uint64(t))

	case float32:
		return float32ToBytes(t)
	case float64:
		return float64ToBytes(t)

	case complex128:
		return complex128ToBytes(t)
	case complex64:
		return complex128ToBytes(complex128(t))

	case bool:
		if t == true {
			return []byte("true")
		}
		return []byte("false")

	case string:
		return []byte(t)

	case []byte:
		return t

	default:
		return []byte(fmt.Sprintf("%v", val))
	}

	panic("Reached.")
}

/*
	Tries to convert the argument into a string. Returns "" if any error occurs.
*/
func String(val interface{}) string {
	var buf []byte

	if val == nil {
		return ""
	}

	switch t := val.(type) {

	case int:
		buf = int64ToBytes(int64(t))
	case int8:
		buf = int64ToBytes(int64(t))
	case int16:
		buf = int64ToBytes(int64(t))
	case int32:
		buf = int64ToBytes(int64(t))
	case int64:
		buf = int64ToBytes(int64(t))

	case uint:
		buf = uint64ToBytes(uint64(t))
	case uint8:
		buf = uint64ToBytes(uint64(t))
	case uint16:
		buf = uint64ToBytes(uint64(t))
	case uint32:
		buf = uint64ToBytes(uint64(t))
	case uint64:
		buf = uint64ToBytes(uint64(t))

	case float32:
		buf = float32ToBytes(t)
	case float64:
		buf = float64ToBytes(t)

	case complex128:
		buf = complex128ToBytes(t)
	case complex64:
		buf = complex128ToBytes(complex128(t))

	case bool:
		if val.(bool) == true {
			return "true"
		}

		return "false"

	case string:
		return t

	case []byte:
		return string(t)

	default:
		return fmt.Sprintf("%v", val)
	}

	return string(buf)
}

/*
// Lets wait until Go 1.1
func List(val interface{}) []interface{} {
	list := []interface{}{}

	if val == nil {
		return list
	}

	switch reflect.TypeOf(val).Kind() {
	case reflect.Slice:
		vval := reflect.ValueOf(val)

		size := vval.Len()
		list := make([]interface{}, size)
		vlist := reflect.ValueOf(list)

		for i := 0; i < size; i++ {
			vlist.Index(i).Set(vval.Index(i))
		}

		return list
	}

	return list
}

// Lets wait until Go 1.1
func Map(val interface{}) map[string]interface{} {

	list := map[string]interface{}{}

	if val == nil {
		return list
	}

	switch reflect.TypeOf(val).Kind() {
	case reflect.Map:
		vval := reflect.ValueOf(val)
		vlist := reflect.ValueOf(list)

		for _, vkey := range vval.MapKeys() {
			key := String(vkey.Interface())
			vlist.SetMapIndex(reflect.ValueOf(key), vval.MapIndex(vkey))
		}

		return list
	}

	return list
}
*/

/*
	Tries to convert the argument into an int64. Returns int64(0) if any error
	occurs.
*/
func Int64(val interface{}) int64 {

	switch t := val.(type) {
	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return int64(t)
	case uint:
		return int64(t)
	case uint8:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case bool:
		if t == true {
			return int64(1)
		}
		return int64(0)
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	default:
		i, _ := strconv.ParseInt(String(val), 10, 64)
		return i
	}

	panic("Reached")

}

/*
	Tries to convert the argument into an uint64. Returns uint64(0) if any error
	occurs.
*/
func Uint64(val interface{}) uint64 {

	switch t := val.(type) {
	case int:
		return uint64(t)
	case int8:
		return uint64(t)
	case int16:
		return uint64(t)
	case int32:
		return uint64(t)
	case int64:
		return uint64(t)
	case uint:
		return uint64(t)
	case uint8:
		return uint64(t)
	case uint16:
		return uint64(t)
	case uint32:
		return uint64(t)
	case uint64:
		return uint64(t)
	case float32:
		return uint64(t)
	case float64:
		return uint64(t)
	case bool:
		if t == true {
			return uint64(1)
		}
		return uint64(0)
	default:
		i, _ := strconv.ParseUint(String(val), 10, 64)
		return i
	}

	panic("Reached")

}

/*
	Tries to convert the argument into a float64. Returns float64(0.0) if any
	error occurs.
*/
func Float64(val interface{}) float64 {

	switch t := val.(type) {
	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case uint:
		return float64(t)
	case uint8:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case float32:
		return float64(t)
	case float64:
		return float64(t)
	case bool:
		if t == true {
			return float64(1)
		}
		return float64(0)
	case string:
		f, _ := strconv.ParseFloat(val.(string), 64)
		return f
	}

	panic("Reached")
}

/*
	Tries to convert the argument into a bool. Returns false if any error occurs.
*/
func Bool(value interface{}) bool {
	b, _ := strconv.ParseBool(String(value))
	return b
}

/*
	Tries to convert the argument into a reflect.Kind element.
*/
func Convert(value interface{}, t reflect.Kind) (interface{}, error) {

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		switch t {
		case reflect.String:
			if reflect.TypeOf(value).Elem().Kind() == reflect.Uint8 {
				return string(value.([]byte)), nil
			} else {
				return String(value), nil
			}
		case reflect.Slice:
		default:
			return nil, fmt.Errorf("Could not convert slice into non-slice.")
		}
	case reflect.String:
		switch t {
		case reflect.Slice:
			return Bytes(value), nil
		}
	}

	switch t {

	case reflect.String:
		return String(value), nil

	case reflect.Uint64:
		return Uint64(value), nil

	case reflect.Uint32:
		return uint32(Uint64(value)), nil

	case reflect.Uint16:
		return uint16(Uint64(value)), nil

	case reflect.Uint8:
		return uint8(Uint64(value)), nil

	case reflect.Uint:
		return uint(Uint64(value)), nil

	case reflect.Int64:
		return int64(Int64(value)), nil

	case reflect.Int32:
		return int32(Int64(value)), nil

	case reflect.Int16:
		return int16(Int64(value)), nil

	case reflect.Int8:
		return int8(Int64(value)), nil

	case reflect.Int:
		return int(Int64(value)), nil

	case reflect.Float64:
		return Float64(value), nil

	case reflect.Float32:
		return float32(Float64(value)), nil

	case reflect.Bool:
		return Bool(value), nil

	case reflect.Interface:
		return value, nil

	case KindTime:
		return Time(value), nil

	case KindDuration:
		return Duration(value), nil

	}

	return nil, fmt.Errorf("Could not convert %s into %s.", reflect.TypeOf(value).Kind(), t)
}
