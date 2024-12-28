package accept_test

import (
	"testing"

	"github.com/UiP9AV6Y/go-slog-adapter"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func TestPromHTTP(t *testing.T) {
	var _ promhttp.Logger = slogadapter.NewDefault()
}
