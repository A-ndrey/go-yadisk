//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/A-ndrey/go-yadisk"
)

func TestDiskMetaInfo(t *testing.T) {
	client := createClient(t)
	_, err := client.DiskMetaInfo(context.Background(), yadisk.DiskMetaInfoParams{})
	if err != nil {
		t.Fatal(err)
	}
}
