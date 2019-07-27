package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

var client *ossClient

type ossClient struct {
	*oss.Client
}

const (
	// oss bucket名
	ossBucket = "pibigstar"
	// oss endpoint
	ossEndpoint = "oss-cn-shanghai.aliyuncs.com"
	// oss访问key
	ossAccessKeyID = "LTAIKFU1CUmLErUw"
	// oss private key secret
	ossAccessKeySecret = "n0axekSPgKwCqIGyBa1oSZBQpOyzlp"
	// 默认失效时间，30天
	defaultBucketExpireTime = 30
)

// NewClient : 创建oss client对象
func NewClient() error {
	if client != nil {
		return nil
	}
	ossCli, err := oss.New(ossEndpoint, ossAccessKeyID, ossAccessKeySecret)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	client = &ossClient{Client: ossCli}
	return nil
}

// Bucket : 获取bucket存储空间
func (client *ossClient) GetBucket(bucketNames ...string) *oss.Bucket {
	if client != nil {
		bucketName := ossBucket
		if len(bucketNames) != 0 {
			bucketName = bucketNames[0]
		}
		bucket, err := client.Bucket(bucketName)
		if err != nil {
			return nil
		}
		return bucket
	}
	return nil
}

// init the oss client
func init() {
	err := NewClient()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// put the object to the oss
func (client *ossClient) Put(key string, reader io.Reader) error {
	return client.GetBucket().PutObject(key, reader)
}

// delete the object, actually use the nil replace the object
func (client *ossClient) DeleteObject(key string) {
	client.Put(key, nil)
}

// get the down url
func (client *ossClient) GetDownloadURL(objName string) string {
	signedURL, err := client.GetBucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedURL
}

// set lifecycle rule for the specified bucket
func (client *ossClient) BuildLifecycleRule(bucketName string) {
	// 表示前缀为test的对象(文件)距最后修改时间30天后过期。
	ruleTest1 := oss.BuildLifecycleRuleByDays("rule1", "test/", true, defaultBucketExpireTime)
	rules := []oss.LifecycleRule{ruleTest1}
	client.SetBucketLifecycle(bucketName, rules)
}
