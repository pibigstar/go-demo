/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          main.go
 * Description:   Reflect Toolkit
 */

package dig

import (
	"fmt"
	"github.com/going/toolkit/to"
	"reflect"
)

/*
	Returns a boolean starting from a Slice or Map.
*/
func Bool(src interface{}, route ...interface{}) bool {
	var b bool
	err := Get(src, &b, route...)
	if err != nil {
		return false
	}
	return b
}

/*
	Returns an uint64 starting from a Slice or Map.
*/
func Uint64(src interface{}, route ...interface{}) uint64 {
	var i uint64
	err := Get(src, &i, route...)
	if err != nil {
		return uint64(0)
	}
	return i
}

/*
	Returns an int64 starting from a Slice or Map.
*/
func Int64(src interface{}, route ...interface{}) int64 {
	var i int64
	err := Get(src, &i, route...)
	if err != nil {
		return int64(0)
	}
	return i
}

/*
	Returns a float32 starting from a Slice or Map.
*/
func Float32(src interface{}, route ...interface{}) float32 {
	var f float32
	err := Get(src, &f, route...)
	if err != nil {
		return float32(0)
	}
	return f
}

/*
	Returns a float64 starting from a Slice or Map.
*/
func Float64(src interface{}, route ...interface{}) float64 {
	var f float64
	err := Get(src, &f, route...)
	if err != nil {
		return float64(0)
	}
	return f
}

/*
	Returns an interface{} starting from a Slice or Map.
*/
func Interface(src interface{}, route ...interface{}) interface{} {
	var i interface{}
	err := Get(src, &i, route...)
	if err != nil {
		return nil
	}
	return i
}

/*
	Returns a string starting from a Slice or Map.
*/
func String(src interface{}, route ...interface{}) string {
	var s string
	err := Get(src, &s, route...)
	if err != nil {
		return ""
	}
	return s
}

/*
	Returns the element of the Slice or Map given by route.
*/
func pick(src interface{}, dig bool, route ...interface{}) (*reflect.Value, error) {
	var err error = nil

	v := reflect.ValueOf(src)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil, fmt.Errorf("Source is not a pointer.")
	}

	v = v.Elem()

	for _, key := range route {
		u := v
		switch v.Kind() {
		case reflect.Slice:
			switch i := key.(type) {
			case int:
				if i < v.Len() {
					v = v.Index(i)
				} else {
					return nil, fmt.Errorf("Undefined index: %d.", i)
				}
			}
		case reflect.Map:
			vkey := reflect.ValueOf(key)
			v = v.MapIndex(vkey)
			if dig == true && v.IsValid() == false {
				u.SetMapIndex(vkey, reflect.MakeMap(u.Type()))
				v = u.MapIndex(vkey)
			}
			if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
				v = v.Elem()
			}
		}
		if v.IsValid() == true {
			if v.CanInterface() == true {
				v = reflect.ValueOf(v.Interface())
			}
		}
	}

	return &v, err
}

/*
	Starts with src (pointer to Slice or Map) tries to follow the given route, if
	the route is found it then tries to set the node with the given value (val).
*/
func Set(src interface{}, val interface{}, route ...interface{}) error {
	l := len(route)

	if l < 1 {
		return fmt.Errorf("Missing route.")
	}

	parent := route[0 : l-1]
	last := route[l-1 : l]

	p, err := pick(src, false, parent...)

	if err != nil {
		return err
	}

	if p.IsValid() == false {
		return fmt.Errorf("Given route does not exists.")
	}

	p.SetMapIndex(reflect.ValueOf(last[0]), reflect.ValueOf(val))

	return nil
}

/*
	Starts with src (pointer to Slice or Map) tries to follow the given route,
	if the route is found it then tries to copy or convert the found node into
	the value pointed by dst.
*/
func Get(src interface{}, dst interface{}, route ...interface{}) error {

	if len(route) < 1 {
		return fmt.Errorf("Missing route.")
	}

	dv := reflect.ValueOf(dst)

	if dv.Kind() != reflect.Ptr || dv.IsNil() {
		return fmt.Errorf("Destination is not a pointer.")
	}

	sv := reflect.ValueOf(src)

	if sv.Kind() != reflect.Ptr || sv.IsNil() {
		return fmt.Errorf("Source is not a pointer.")
	}

	// Setting to zero before setting it again.
	dv.Elem().Set(reflect.Zero(dv.Elem().Type()))

	p, err := pick(src, false, route...)

	if err != nil {
		return err
	}

	if p.IsValid() == false {
		return fmt.Errorf("Could not find route.")
	}

	if dv.Elem().Type() != p.Type() {
		// Trying conversion
		if p.CanInterface() == true {
			var t interface{}
			t, err = to.Convert(p.Interface(), dv.Elem().Kind())
			if err == nil {
				tv := reflect.ValueOf(t)
				if dv.Elem().Type() == tv.Type() {
					p = &tv
				}
			}
		}
	}

	if dv.Elem().Type() == p.Type() || dv.Elem().Kind() == reflect.Interface {
		dv.Elem().Set(*p)
	} else {
		return fmt.Errorf("Could not assign %s to %s.", p.Type(), dv.Elem().Type())
	}

	return nil
}

/*
	Makes a path to the given route, if the route already exists it overwrites it
	with a zero value.
*/
func Dig(src interface{}, route ...interface{}) error {
	v, err := pick(src, true, route...)
	if v.IsValid() == false {
		return fmt.Errorf("Could not reach node.")
	}
	return err
}
