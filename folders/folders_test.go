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

func TestGetPaginatedAllFolders(t *testing.T) {
	orgID, err := uuid.FromString("3b9a868b-8cd9-4b6b-ba23-fd1e08f3e9fa")
	if err != nil {
		t.Fatalf("Failed to parse orgID: %v", err)
	}

	t.Run("Successful Pagination Retrieval", func(t *testing.T) {
		// Define a request with a specific limit
		req := &folders.PaginatedFetchFolderRequest{
			OrgID:  orgID,
			Limit:  1,  // Setting the limit to 1 for this test
			Cursor: "", // Start with an empty cursor to get the first page
		}

		// Call the paginated function
		resp, err := folders.GetPaginatedAllFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Folders, 1) // We expect only one folder due to the limit

		// Check the details of the first folder
		firstFolder := resp.Folders[0]
		assert.Equal(t, "71702b42-aee8-4c03-a05c-1a0cc5102a85", firstFolder.Id.String())
		assert.Equal(t, "sawa-test-1", firstFolder.Name)
		assert.Equal(t, orgID, firstFolder.OrgId)
		assert.False(t, firstFolder.Deleted)

		// Now, use the returned cursor to get the next page
		nextReq := &folders.PaginatedFetchFolderRequest{
			OrgID:  orgID,
			Limit:  1,
			Cursor: resp.NextCursor, // Use the cursor from the previous response
		}

		nextResp, err := folders.GetPaginatedAllFolders(nextReq)
		assert.NoError(t, err)
		assert.NotNil(t, nextResp)
		assert.Len(t, nextResp.Folders, 1) // We expect the second folder on the next page

		// Check the details of the second folder
		secondFolder := nextResp.Folders[0]
		assert.Equal(t, "71702b42-aee8-4c03-a05c-1a0cc5102a86", secondFolder.Id.String())
		assert.Equal(t, "sawa-test-2", secondFolder.Name)
		assert.Equal(t, orgID, secondFolder.OrgId)
		assert.True(t, secondFolder.Deleted)

		// Ensure there's no next cursor after the second folder, indicating the end of the data
		assert.Empty(t, nextResp.NextCursor)
	})

	//
	// 	t.Run("Paginate Beyond Data Set", func(t *testing.T) {
	// 		// Request the first page
	// 		req := &folders.PaginatedFetchFolderRequest{
	// 			OrgID:  orgID,
	// 			Limit:  1,
	// 			Cursor: "",
	// 		}
	// 		resp, _ := folders.GetPaginatedAllFolders(req)
	//
	// 		// Request a page beyond the available data
	// 		nextReq := &folders.PaginatedFetchFolderRequest{
	// 			OrgID:  orgID,
	// 			Limit:  1,
	// 			Cursor: resp.NextCursor,
	// 		}
	// 		nextResp, err := folders.GetPaginatedAllFolders(nextReq)
	// 		assert.NoError(t, err)
	// 		assert.NotNil(t, nextResp)
	//
	// 		// Since we only have 2 folders, and we're requesting the third page, it should be empty
	// 		assert.Empty(t, nextResp.Folders)
	// 		// The next cursor should also be empty since there is no more data
	// 		assert.Empty(t, nextResp.NextCursor)
	// 	})

	t.Run("Invalid Cursor Token", func(t *testing.T) {
		// Request with an invalid cursor token
		req := &folders.PaginatedFetchFolderRequest{
			OrgID:  orgID,
			Limit:  1,
			Cursor: "invalidCursor",
		}
		_, err := folders.GetPaginatedAllFolders(req)
		// We expect an error due to the invalid cursor
		assert.Error(t, err)
	})

	t.Run("Limit Larger Than Data Set", func(t *testing.T) {
		// Request with a limit larger than the data set
		req := &folders.PaginatedFetchFolderRequest{
			OrgID:  orgID,
			Limit:  10, // Large limit to fetch all data in one go
			Cursor: "",
		}
		resp, err := folders.GetPaginatedAllFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		// We expect to get all folders since the limit exceeds the number of available folders
		assert.Len(t, resp.Folders, 2)
		// Since all data is fetched, the next cursor should be empty
		assert.Empty(t, resp.NextCursor)
	})

}
