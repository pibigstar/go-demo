package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"log"
	"os"
	"reflect"
)

var client *elastic.Client
var host = "http://localhost:9700"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int32    `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

// 初始化
func init() {
	errLog := log.New(os.Stdout, "App", log.LstdFlags)

	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errLog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}

	result, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	log.Printf("Elasticsearch returned with code: %d and version: %s", code, result.Version.Number)

	version, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	log.Printf("Elasticsearch version :%s", version)
}

// 创建
func create() {
	employee := &Employee{
		FirstName: "Pi",
		LastName:  "bigstar",
		Age:       18,
		About:     "test",
		Interests: []string{"henan", "zhengzhou"},
	}
	// http://localhost:9700/pibigstar/employee/1
	response, err := client.Index().
		Index("pibigstar").
		Type("employee").
		Id("1").
		BodyJson(employee).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Printf("Index: %s, Id: %s, Type: %s \n", response.Index, response.Id, response.Type)
}

// 查询
func get(id string) {
	var employee Employee
	// id 必须存在，不然会报错，如果想查找请用search
	result, err := client.Get().Index("pibigstar").Type("employee").Id(id).Do(context.Background())
	if err != nil && result.Found {
		//panic(err)
	}
	if result.Found {
		bytes, _ := result.Source.MarshalJSON()
		json.Unmarshal(bytes, &employee)
		fmt.Printf("%+v \n", employee)
	}
}

// 搜索
func search() {
	// 查询全部
	result, err := client.Search().Index("pibigstar").Type("employee").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if result.Hits.TotalHits > 0 {
		printResult(result)
	}
	// 根据名字查询
	query := elastic.NewQueryStringQuery("last_name:bigstar")
	searchResult, err := client.Search().Index("pibigstar").Type("employee").Query(query).Do(context.Background())
	if err != nil {
		panic(err)
	}
	printResult(searchResult)

}

// 删除
func delete(id string) {
	response, err := client.Delete().Index("pibigstar").Type("employee").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete the result is: %s \n", response.Result)
}

// 更新
func update(id string) {
	response, err := client.Update().
		Index("pibigstar").
		Type("employee").
		Id(id).
		Doc(map[string]interface{}{"age": 22}).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("update the result is : %s \n", response.Result)
}

// 打印结果
func printResult(result *elastic.SearchResult) {
	var typ Employee
	for _, item := range result.Each(reflect.TypeOf(typ)) {
		employee := item.(Employee)
		fmt.Printf("%+v\n", employee)
	}
}

func main() {
	get("1")
}
