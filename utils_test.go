package oanda

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func Test_request(t *testing.T) {
	// 到達不能なエンドポイントを指定
	t.Run("ConnectionRefused", func(t *testing.T) {
		connection := newConnection(t, oandaDummy)
		connection.Timeout = time.Nanosecond // 即タイムアウトさせるため最小の待ち時間にする
		_, err := connection.Accounts().Get(context.Background())

		if _, ok := errors.Unwrap(err).(*url.Error); !ok {
			t.Fatalf("Connection was not refused.\n%+v", err)
		}
	})

	// 認証エラー
	t.Run("Unauthorized", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		connection.Token = "hogehoge" // 不正なトークンに書き換え
		_, err := connection.Accounts().Get(context.Background())

		if _, ok := errors.Unwrap(err).(*UnauthorizedError); !ok {
			t.Fatalf("Request was authorized.\n%+v", err)
		}
	})
}
