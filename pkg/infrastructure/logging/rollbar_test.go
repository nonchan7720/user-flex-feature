package logging

import (
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	httpclient "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/httpClient"
	"github.com/nonchan7720/user-flex-feature/pkg/mock"
	"github.com/stretchr/testify/assert"
)

func testHTTPClient(transport mock.MockRoundTripper) *http.Client {
	client := httpclient.DefaultClient(transport)
	return client
}

func TestRollbar(t *testing.T) {
	done := make(chan struct{})
	called := false
	defer func() {
		<-done
		assert.True(t, called)
	}()
	transport := func(req *http.Request) (*http.Response, error) {
		called = true
		done <- struct{}{}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("")),
		}, nil
	}
	conf := RollbarConfig{
		Level:  "error",
		Token:  "DUMMY",
		Client: testHTTPClient(transport),
	}
	conf.Init("local", "v1", "test")
	h := NewRollbarHandler(&conf)
	defer h.Close()
	log := slog.New(h)
	slog.SetDefault(log)
	slog.Error("This is test")
}
