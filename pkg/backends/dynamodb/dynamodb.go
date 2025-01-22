package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/k3s-io/kine/pkg/server"
	"log"
)

// Options 包含 DynamoDB 的配置信息
type Options struct {
	Ip      string
	Port    int
	Table   string
	Timeout int
}

/*

aws dynamodb create-table --table-name balancer-routes \
--attribute-definitions AttributeName=pk,AttributeType=S AttributeName=sk,AttributeType=S \
--key-schema AttributeName=pk,KeyType=HASH AttributeName=sk,KeyType=RANGE \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --endpoint-url http://127.0.0.1:8080

aws dynamodb delete-table --table-name balancer-routes --endpoint-url http://127.0.0.1:8080
*/

// dynamodbCache 是使用 DynamoDB 实现的后端缓存
type dynamodbCache struct {
	options *Options
	client  *dynamodb.Client
}

// NewDynamoDBCache 返回一个实现了 server.Backend 接口的 DynamoDB 后端实例
func NewDynamoDBCache(ctx context.Context, options *Options) (server.Backend, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "test2", SecretAccessKey: "test2",
			},
		}),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg, func(awsOptions *dynamodb.Options) {
		awsOptions.BaseEndpoint = aws.String(fmt.Sprintf("http://%s:%d", options.Ip, options.Port))
		awsOptions.RetryMaxAttempts = 2
	})
	backend := &dynamodbCache{options: options, client: client}
	return backend, nil
}

func (dy *dynamodbCache) Start(ctx context.Context) error {
	fmt.Println("DynamoDB backend started")
	return nil
}

func (dy *dynamodbCache) Count(ctx context.Context, prefix string) (int64, int64, error) {
	log.Println("func (dy *dynamodbCache) Count(ctx context.Context, prefix string) (int64, int64, error) {")
	//count, err := f.db.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
	//	rangeResult := tr.GetRange(f.dir.Sub(prefix), fdb.RangeOptions{}).GetSliceOrPanic()
	//	return int64(len(rangeResult)), nil
	//})
	//
	//if err != nil {
	//	return 0, 0, err
	//}

	revision := int64(1)
	return 1, revision, nil
}

func (dy *dynamodbCache) Get(ctx context.Context, key string, rangeEnd string, limit, revision int64) (int64, *server.KeyValue, error) {
	fmt.Println("func (dy *dynamodbCache) Get(ctx context.Context, key string, rangeEnd string, limit, revision int64) (int64, *server.KeyValue, error) {")
	//var kv *server.KeyValue
	//var newRevision int64
	//
	//_, err := f.db.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
	//	result := tr.Get(fdb.Key(key)).MustGet()
	//	log.Printf("Get: key=%s, result=%s", key, string(result))
	//
	//	if len(result) != 0 {
	//		kv = &server.KeyValue{
	//			Key:   key,
	//			Value: result,
	//		}
	//		newRevision = revision
	//	}
	//
	//	return nil, nil
	//})
	//
	//if err != nil {
	//	log.Printf("Get: error=%v", err)
	//	return 0, nil, err
	//}
	//
	//if kv == nil {
	//	log.Printf("Get: kv is nil for key=%s", key)
	//	return 0, nil, nil
	//}
	//
	//log.Printf("Get: returning kv=%v, newRevision=%d", kv, newRevision)
	return 1, nil, nil
}

func (dy *dynamodbCache) Put(ctx context.Context, key string, value []byte) error {
	fmt.Println("func (dy *dynamodbCache) Put(ctx context.Context, key string, value []byte) error {")
	//_, err := f.db.Transact(func(tr fdb.Transaction) (interface{}, error) {
	//	tr.Set(fdb.Key(key), value)
	//	return nil, nil
	//})
	return nil
}

func (dy *dynamodbCache) Delete(ctx context.Context, key string, revision int64) (int64, *server.KeyValue, bool, error) {
	fmt.Println("func (dy *dynamodbCache) Delete(ctx context.Context, key string, revision int64) (int64, *server.KeyValue, bool, error) {")
	//var deletedRev int64
	//var deletedKV *server.KeyValue
	//var deleted bool
	//
	//_, err := f.db.Transact(func(tr fdb.Transaction) (interface{}, error) {
	//	val := tr.Get(fdb.Key(key)).MustGet()
	//	if len(val) == 0 {
	//		return nil, nil
	//	}
	//
	//	tr.Clear(fdb.Key(key))
	//
	//	deletedRev = revision
	//	deletedKV = &server.KeyValue{
	//		Key:   key,
	//		Value: val,
	//	}
	//	deleted = true
	//
	//	return nil, nil
	//})
	//
	//if err != nil {
	//	return 0, nil, false, err
	//}

	return 1, nil, false, nil
}

func (dy *dynamodbCache) Create(ctx context.Context, key string, value []byte, revision int64) (int64, error) {
	fmt.Println("func (dy *dynamodbCache) Create(ctx context.Context, key string, value []byte, revision int64) (int64, error) {")
	//rev, err := f.db.Transact(func(tr fdb.Transaction) (interface{}, error) {
	//	existing := tr.Get(fdb.Key(key)).MustGet()
	//	if len(existing) != 0 {
	//		return nil, errors.New("key already exists")
	//	}
	//	tr.Set(fdb.Key(key), value)
	//	return revision, nil
	//})
	//
	//if err != nil {
	//	return 0, err
	//}
	return 1, nil
}

func (dy *dynamodbCache) DbSize(ctx context.Context) (int64, error) {
	fmt.Println("func (dy *dynamodbCache) DbSize(ctx context.Context) (int64, error) {")
	//size, err := f.db.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
	//	rangeResult := tr.GetRange(f.dir, fdb.RangeOptions{}).GetSliceOrPanic()
	//	var totalSize int64
	//	for _, kv := range rangeResult {
	//		totalSize += int64(len(kv.Key) + len(kv.Value))
	//	}
	//	return totalSize, nil
	//})
	//
	//if err != nil {
	//	return 0, err
	//}
	return 1, nil
}

func (dy *dynamodbCache) List(ctx context.Context, prefix string, rangeEnd string, limit, revision int64) (int64, []*server.KeyValue, error) {
	fmt.Println("func (dy *dynamodbCache) List(ctx context.Context, prefix string, rangeEnd string, limit, revision int64) (int64, []*server.KeyValue, error) {")
	//var kvs []*server.KeyValue
	//var newRevision int64
	//
	//_, err := f.db.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
	//	rangeResult := tr.GetRange(f.dir.Sub(prefix), fdb.RangeOptions{Limit: int(limit)}).GetSliceOrPanic()
	//
	//	for _, item := range rangeResult {
	//		if limit > 0 && int64(len(kvs)) >= limit {
	//			break
	//		}
	//		kv := &server.KeyValue{
	//			Key:   string(item.Key),
	//			Value: item.Value,
	//		}
	//		kvs = append(kvs, kv)
	//	}
	//
	//	newRevision = revision
	//	return nil, nil
	//})
	//
	//if err != nil {
	//	return 0, nil, err
	//}

	return 1, nil, nil
}

func (dy *dynamodbCache) Update(ctx context.Context, key string, value []byte, oldRevision, newRevision int64) (int64, *server.KeyValue, bool, error) {
	fmt.Println("func (dy *dynamodbCache) Update(ctx context.Context, key string, value []byte, oldRevision, newRevision int64) (int64, *server.KeyValue, bool, error) {")
	var updatedKV *server.KeyValue
	//var updated bool
	//
	//_, err := f.db.Transact(func(tr fdb.Transaction) (interface{}, error) {
	//	existing := tr.Get(fdb.Key(key)).MustGet()
	//
	//	if len(existing) == 0 {
	//		return nil, errors.New("key does not exist")
	//	}
	//
	//	if oldRevision != newRevision {
	//		return nil, errors.New("revision mismatch")
	//	}
	//
	//	tr.Set(fdb.Key(key), value)
	//	updatedKV = &server.KeyValue{
	//		Key:   key,
	//		Value: value,
	//	}
	//	updated = true
	//
	//	return nil, nil
	//})
	//
	//if err != nil {
	//	return 0, nil, false, err
	//}
	return newRevision, updatedKV, false, nil
}

func (dy *dynamodbCache) Watch(ctx context.Context, key string, revision int64) <-chan []*server.Event {

	events := make(chan []*server.Event)

	go func() {
		defer close(events)
	}()

	return events
}
