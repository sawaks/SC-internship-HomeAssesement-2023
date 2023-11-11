package folders_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	orgID, err := uuid.FromString("3b9a868b-8cd9-4b6b-ba23-fd1e08f3e9fa")
	if err != nil {
		t.Fatalf("Failed to parse orgID: %v", err)
	}

	t.Run("Successful Folder Retrieval for Valid OrgID", func(t *testing.T) {

		req := &folders.FetchFolderRequest{OrgID: orgID}
		resp, err := folders.GetAllFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		respJSON, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			t.Fatalf("Failed to marshal response: %v", err)
		}
		fmt.Println("Response!!!!!:", string(respJSON))
	})

	t.Run("No Result for Non-Existing Organization ID", func(t *testing.T) {
		emptyOrgID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
		req := &folders.FetchFolderRequest{OrgID: emptyOrgID}
		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, resp.Folders)
	})

}
