package main

import (
	"log"
	"sort"
	"time"
)

type Person struct {
	Age      int
	Name     string
	Birthday time.Time
}

type PersonSwapper struct {
	p  []Person
	by func(p, q *Person) bool
}

func main() {
	persons := []Person{{18, "li", time.Now()}, {11, "hua", time.Now()}, {25, "tt", time.Now()}}

	// q 的年龄大，将其排到后面，所以返回true
	sort.Sort(PersonSwapper{persons, func(p, q *Person) bool {
		return p.Age < q.Age // 按年龄递增排序
	}})

	for _, person := range persons {
		log.Printf("%+v \n", person)
	}

	sort.Sort(PersonSwapper{persons, func(p, q *Person) bool {
		return p.Name < q.Name //按姓名递增
	}})

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
