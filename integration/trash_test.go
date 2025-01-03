package integration

import (
	"context"
	"testing"

	"github.com/A-ndrey/go-yadisk"
)

func TestTrash(t *testing.T) {
	ctx := context.Background()
	client := createClient(t)

	dirPath := testResourcePrefix + "dir-for-trash"

	findTrashPath := func(tr *yadisk.TrashResource) string {
		for _, item := range tr.Embedded.Items {
			if item.Name == dirPath {
				return item.Path
			}
		}

		t.Fatal("resource not found in trash")

		return ""
	}

	// create directory that will be deleted (moved to trash)
	_, err := client.CreateDirectory(ctx, yadisk.CreateDirectoryParams{Path: dirPath})
	if err != nil {
		t.Fatal(err)
	}

	// move directory to trash
	_, err = client.DeleteResource(ctx, yadisk.DeleteResourceParams{Path: dirPath})
	if err != nil {
		t.Fatal(err)
	}

	// find directory in trash
	tcontent, err := client.TrashContent(ctx, yadisk.TrashContentParams{Path: "/"})
	if err != nil {
		t.Fatal(err)
	}

	// restore directory
	_, err = client.RestoreResource(ctx, yadisk.RestoreResourceParams{Path: findTrashPath(tcontent)})
	if err != nil {
		t.Fatal(err)
	}

	// delete restored directory again
	_, err = client.DeleteResource(ctx, yadisk.DeleteResourceParams{Path: dirPath})
	if err != nil {
		t.Fatal(err)
	}

	// find directory in trash
	tcontent, err = client.TrashContent(ctx, yadisk.TrashContentParams{Path: "/"})
	if err != nil {
		t.Fatal(err)
	}

	// delete directory from trash permanently
	_, err = client.CleanTrash(ctx, yadisk.CleanTrashParams{Path: findTrashPath(tcontent)})
	if err != nil {
		return
	}
}
