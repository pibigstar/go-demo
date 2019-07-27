//+build !test

package elastic

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

var client *esClient

type esClient struct {
	*elastic.Client
	ctx context.Context
}

const host = "http://localhost:9200"

// build the elasticSearch the client
func NewElasticSearchClient() error {
	errLog := log.New(os.Stdout, "Elastic", log.LstdFlags)

	esCli, err := elastic.NewClient(elastic.SetErrorLog(errLog), elastic.SetURL(host))
	if err != nil {
		return err
	}
	client = &esClient{Client: esCli, ctx: context.Background()}
	result, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		return err
	}
	log.Printf("Elasticsearch returned with code: %d and version: %s", code, result.Version.Number)

	version, err := client.ElasticsearchVersion(host)
	if err != nil {
		return err
	}
	log.Printf("Elasticsearch version :%s", version)

	return nil
}

// init the elastic search
func init() {
	err := NewElasticSearchClient()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// insert a document to the index
func (client *esClient) Insert(index, typeName string, value interface{}) (*elastic.IndexResponse, error) {
	// access by the http://localhost:9700/pibigstar/employee/id
	response, err := client.Index().
		Index(index).
		Type(typeName).
		BodyJson(value).
		Do(client.ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// get the document by id
func (client *esClient) GetById(index, typeName, id string) ([]byte, error) {
	// id 必须存在，不然会报错，如果想查找请用search
	result, err := client.Get().Index(index).Type(typeName).Id(id).Do(client.ctx)
	if err != nil && !result.Found {
		return nil, err
	}
	if result.Found {
		bytes, _ := result.Source.MarshalJSON()
		return bytes, nil
	}
	return nil, nil
}

// search the result by query strings
func (client *esClient) Query(index, typeName string, queryStrings ...string) (*elastic.SearchResult, error) {
	var queryString string
	if len(queryStrings) > 0 {
		queryString = queryStrings[0]
	}
	// 根据名字查询
	query := elastic.NewQueryStringQuery(queryString)
	result, err := client.Search().Index(index).Type(typeName).Query(query).Do(client.ctx)
	if err != nil {
		return nil, err
	}
	if result.Hits.TotalHits > 0 {
		return result, nil
	}
	return nil, errors.New("query the result is null")
}

// delete the document by id
func (client *esClient) DeleteById(index, typeName, id string) (*elastic.DeleteResponse, error) {
	response, err := client.Delete().Index(index).Type(typeName).Id(id).Do(client.ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// update the document by id
// values : map[string]interface{}{"age": 12}
func (client *esClient) UpdateById(index, typeName, id string, values map[string]interface{}) (*elastic.UpdateResponse, error) {
	response, err := client.Update().
		Index(index).
		Type(typeName).
		Id(id).
		Doc(values).Do(client.ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}
