package accept_test

import (
	"testing"

	"github.com/UiP9AV6Y/go-slog-adapter"

	"github.com/hashicorp/go-retryablehttp"
)

func TestRetryableHTTP(t *testing.T) {
	var _ retryablehttp.Logger = slogadapter.NewDefault()
	// retryablehttp.LeveledLogger is covered by the native slog.Logger
	// implementation aleady; we test its contract here for completeness sake.
	// not that we can change anything if the contract breaks, apart from
	// implementing our own bridge.
	var _ retryablehttp.LeveledLogger = slogadapter.NewDefault().Logger()
}
