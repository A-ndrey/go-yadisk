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

	// create directory
	_, err := client.CreateDirectory(ctx, yadisk.CreateDirectoryParams{Path: dirPath})
	if err != nil {
		t.Fatal(err)
	}

	// get directory meta info
	_, err = client.ResourceMetaInfo(ctx, yadisk.ResourceMetaInfoParams{Path: dirPath})
	if err != nil {
		t.Fatal(err)
	}

	// delete directory permanently
	_, err = client.DeleteResource(ctx, yadisk.DeleteResourceParams{Path: dirPath, Permanently: true})
	if err != nil {
		t.Fatal(err)
	}
}

func TestFile(t *testing.T) {
	ctx := context.Background()
	filePath := testResourcePrefix + "file"
	fileContent := "test file content"

	client := createClient(t)

	hasPublicFile := func(pubRes *yadisk.PublicResourceList) bool {
		for _, item := range pubRes.Items {
			if item.Name == filePath {
				return true
			}
		}

		return false
	}

	// get upload link
	uploadFileLink, err := client.UploadFileLink(ctx, yadisk.UploadFileLinkParams{Path: filePath, Overwrite: true})
	if err != nil {
		t.Fatal(err)
	}

	// upload file
	err = client.Upload(ctx, uploadFileLink.URL, strings.NewReader(fileContent))
	if err != nil {
		t.Fatal(err)
	}

	// wait for operation
	waitOperation(t, client, uploadFileLink.OperationID, 5*time.Second)

	// get file meta info
	_, err = client.ResourceMetaInfo(ctx, yadisk.ResourceMetaInfoParams{Path: filePath})
	if err != nil {
		t.Fatal(err)
	}

	// copy file
	copyFilePath := filePath + "-copy"
	_, err = client.CopyResource(ctx, yadisk.CopyResourceParams{FromPath: filePath, ToPath: copyFilePath})
	if err != nil {
		t.Fatal(err)
	}

	// get files copy meta info
	_, err = client.ResourceMetaInfo(ctx, yadisk.ResourceMetaInfoParams{Path: copyFilePath})
	if err != nil {
		t.Fatal(err)
	}

	// get link for downloading files copy
	downloadLink, err := client.DownloadFileLink(ctx, yadisk.DownloadFileLinkParams{Path: copyFilePath})
	if err != nil {
		t.Fatal(err)
	}

	// download files copy
	sb := strings.Builder{}
	err = client.Download(ctx, downloadLink.URL, &sb)
	if err != nil {
		t.Fatal(err)
	}

	if sb.String() != fileContent {
		t.Fatal("file content mismatch")
	}

	// rename files copy
	renamedFilePath := filePath + "-renamed"
	_, err = client.MoveResource(ctx, yadisk.MoveResourceParams{FromPath: copyFilePath, ToPath: renamedFilePath})
	if err != nil {
		t.Fatal(err)
	}

	// delete renamed file permanently
	_, err = client.DeleteResource(ctx, yadisk.DeleteResourceParams{Path: renamedFilePath, Permanently: true})
	if err != nil {
		t.Fatal(err)
	}

	// publish file
	_, err = client.PublishResource(ctx, yadisk.PublishResourceParams{Path: filePath})
	if err != nil {
		t.Fatal(err)
	}

	// check if file is public
	listPublic, err := client.ListPublicResources(ctx, yadisk.ListPublicResourcesParams{})
	if err != nil {
		t.Fatal(err)
	}

	if !hasPublicFile(listPublic) {
		t.Fatal("public file not found")
	}

	// unpublish file
	_, err = client.UnpublishResource(ctx, yadisk.UnpublishResourceParams{Path: filePath})
	if err != nil {
		t.Fatal(err)
	}

	// check if file is public
	listPublic, err = client.ListPublicResources(ctx, yadisk.ListPublicResourcesParams{})
	if err != nil {
		t.Fatal(err)
	}

	if hasPublicFile(listPublic) {
		t.Fatal("file still public")
	}

	// delete file permanently
	_, err = client.DeleteResource(ctx, yadisk.DeleteResourceParams{Path: filePath, Permanently: true})
	if err != nil {
		t.Fatal(err)
	}
}
