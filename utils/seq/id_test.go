package seq

import (
	"context"
	"encoding/json"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc/metadata"
	"testing"
)

// 利用雪花算法生成不重复ID
func TestID(t *testing.T) {
	id := NextNumID()
	t.Log(id)
	body, err := json.Marshal(sonyflake.Decompose(id))
	if err != nil {
		t.Error(err)
	}
	t.Log(string(body))
}

//根据上下文生成带前缀的ID
func TestNextID(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(IdPrefixKey, "P66-"))
	id := NextID(ctx)
	t.Log(id)
}
