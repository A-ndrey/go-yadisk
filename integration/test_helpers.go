//go:build integration

package integration

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/A-ndrey/go-yadisk"
)

const (
	testResourcePrefix = "integration-test-"
)

func createClient(t *testing.T) *yadisk.Client {
	token := os.Getenv("YA_DISK_TOKEN")
	if token == "" {
		t.Fatal(errors.New("YA_DISK_TOKEN is not set"))
	}

	return yadisk.NewClient(token, yadisk.DefaultHost, yadisk.WithHttpClient(&http.Client{Timeout: 10 * time.Second}))
}

func waitOperation(t *testing.T, client *yadisk.Client, operationID string, retryInterval time.Duration) {
	for {
		op, err := client.OperationStatus(context.Background(), operationID)
		if err != nil {
			t.Fatal(err)
		}

		switch op.Status {
		case yadisk.OpStatusInProgress:
			time.Sleep(retryInterval)
		case yadisk.OpStatusSuccess:
			return
		case yadisk.OpStatusFailed:
			t.Fatal("operation failed")
		}
	}
}
