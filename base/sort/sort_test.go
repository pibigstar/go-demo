package main

import (
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	ids := []int{1, 5, 6, 11, 19, 2, 7}
	//递增排序,方法一
	sort.Ints(ids)
	t.Log(ids)
	//递增排序，方法二
	sort.Sort(sort.IntSlice(ids))
	t.Log(ids)

	//递减排序
	reverse := sort.Reverse(sort.IntSlice(ids))
	sort.Sort(reverse)
	t.Log(ids)
}

func TestSortObj(t *testing.T) {
	persons := []Person{
		{18, "li", time.Now()},
		{11, "hua", time.Now()},
		{25, "tt", time.Now()},
	}
	// q 的年龄大，将其排到后面，所以返回true
	sort.Sort(PersonSwapper{persons, func(p, q *Person) bool {
		return p.Age < q.Age // 按年龄递增排序
	}})
	t.Log(persons)

	sort.Sort(PersonSwapper{persons, func(p, q *Person) bool {
		return p.Name < q.Name //按姓名递增
	}})
	t.Log(persons)

	// 先按年龄再按姓名
	sort.Sort(PersonSwapper{persons, func(p, q *Person) bool {
		if p.Age < q.Age {
			return true
		}
		if p.Age == q.Age {
			if p.Name < q.Name {
				return true
			}
		}
		return false
	}})
	t.Log(persons)
}

// 数组排序更简易的方式
func TestSortSlice(t *testing.T) {
	persons := []Person{
		{18, "li", time.Now()},
		{11, "hua", time.Now()},
		{25, "tt", time.Now()},
	}
	sort.Slice(persons, func(i, j int) bool {
		if persons[i].Age < persons[j].Age {
			return true
		}
		return false
	})
	t.Log(persons)
}

type Person struct {
	Age      int
	Name     string
	Birthday time.Time
}

type PersonSwapper struct {
	p  []Person
	by func(p, q *Person) bool
}

func (pw PersonSwapper) Swap(i, j int) {
	pw.p[i], pw.p[j] = pw.p[j], pw.p[i]
}

func (pw PersonSwapper) Len() int {
	return len(pw.p)
}

func (pw PersonSwapper) Less(i, j int) bool {
	return pw.by(&pw.p[i], &pw.p[j])
}
