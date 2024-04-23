package serialize

import (
	"context"
	"encoding/base64"
	reflect "reflect"

	"google.golang.org/protobuf/proto"
)

type ProtoSerializer[V proto.Message] struct{}

var (
	_ Serialize[proto.Message] = &ProtoSerializer[proto.Message]{}
)

func (*ProtoSerializer[V]) Encode(ctx context.Context, value V) (string, error) {
	v, err := proto.Marshal(value)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(v), nil
}

func (*ProtoSerializer[V]) Decode(ctx context.Context, value string) (V, error) {
	var vt V
	vType := reflect.TypeOf(vt).Elem()       // Vの型を取得
	vValue := reflect.New(vType).Interface() // 新しいV型のインスタンスを作成
	v := vValue.(V)
	buf, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return v, err
	}
	if err := proto.Unmarshal(buf, proto.Message(v)); err != nil {
		return v, err
	}
	return v, nil
}
