//go:build integration

package integration

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/A-ndrey/go-yadisk"
)

func TestDirectory(t *testing.T) {
	ctx := context.Background()
	dirPath := testResourcePrefix + "dir"

	client := createClient(t)

	_, err := client.CreateDirectory(ctx, yadisk.CreateDirectoryParams{Path: dirPath})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.ResourceMetaInfo(ctx, yadisk.ResourceMetaInfoParams{Path: dirPath})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DeleteResource(ctx, yadisk.DeleteResourceParams{Path: dirPath, Permanently: true})
	if err != nil {
		t.Fatal(err)
	}
}

func TestFile(t *testing.T) {
	ctx := context.Background()
	filePath := testResourcePrefix + "file"

	client := createClient(t)

	uploadFileLink, err := client.UploadFileLink(ctx, yadisk.UploadFileLinkParams{Path: filePath, Overwrite: true})
	if err != nil {
		t.Fatal(err)
	}

	err = client.Upload(ctx, uploadFileLink.URL, strings.NewReader("test file content"))
	if err != nil {
		t.Fatal(err)
	}

	waitOperation(t, client, uploadFileLink.OperationID, 5*time.Second)

	_, err = client.ResourceMetaInfo(ctx, yadisk.ResourceMetaInfoParams{Path: filePath})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DeleteResource(ctx, yadisk.DeleteResourceParams{Path: filePath, Permanently: true})
	if err != nil {
		t.Fatal(err)
	}
}
