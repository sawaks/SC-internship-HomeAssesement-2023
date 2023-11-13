package folders_test

import (
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
		// your test/s here
		req := &folders.FetchFolderRequest{OrgID: orgID}
		resp, err := folders.GetAllFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		assert.Len(t, resp.Folders, 2)

		assert.Equal(t, "71702b42-aee8-4c03-a05c-1a0cc5102a85", resp.Folders[0].Id.String())
		assert.Equal(t, "sawa-test-1", resp.Folders[0].Name)
		assert.Equal(t, orgID, resp.Folders[0].OrgId)
		assert.False(t, resp.Folders[0].Deleted)
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
			OrgID: orgID,
			Limit: 1,  // Setting the limit to 1 for this test
			Token: "", // Start with an empty cursor to get the first page
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
			OrgID: orgID,
			Limit: 1,
			Token: resp.Token, // Use the cursor from the previous response
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
		assert.Equal(t, "END_OF_DATA", nextResp.Token)

	})

	t.Run("Paginate Beyond Data Set", func(t *testing.T) {
		// Request the first page
		firstReq := &folders.PaginatedFetchFolderRequest{
			OrgID: orgID,
			Limit: 1, // Assuming limit is set to the number of folders per page
			Token: "",
		}
		firstResp, _ := folders.GetPaginatedAllFolders(firstReq)

		// Request the second page using the cursor from the first response
		secondReq := &folders.PaginatedFetchFolderRequest{
			OrgID: orgID,
			Limit: 1,
			Token: firstResp.Token,
		}
		secondResp, _ := folders.GetPaginatedAllFolders(secondReq)

		// Now, try to get the third page, which should be empty
		thirdReq := &folders.PaginatedFetchFolderRequest{
			OrgID: orgID,
			Limit: 1,
			Token: secondResp.Token,
		}
		thirdResp, err := folders.GetPaginatedAllFolders(thirdReq)

		// Assertions to ensure the third response is empty
		assert.NoError(t, err)
		assert.Empty(t, thirdResp.Folders, "Expected no folders on the third page")
		assert.Equal(t, "END_OF_DATA", thirdResp.Token)

	})

	t.Run("Invalid Token", func(t *testing.T) {
		// Request with an invalid cursor token
		req := &folders.PaginatedFetchFolderRequest{
			OrgID: orgID,
			Limit: 1,
			Token: "invalidToken",
		}
		_, err := folders.GetPaginatedAllFolders(req)
		// We expect an error due to the invalid cursor
		assert.Error(t, err)
	})

	t.Run("Limit Larger Than Data Set", func(t *testing.T) {
		// Request with a limit larger than the data set
		req := &folders.PaginatedFetchFolderRequest{
			OrgID: orgID,
			Limit: 10, // Large limit to fetch all data in one go
			Token: "",
		}
		resp, err := folders.GetPaginatedAllFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		// We expect to get all folders since the limit exceeds the number of available folders
		assert.Len(t, resp.Folders, 2)
		// Since all data is fetched, the next cursor should be empty
		assert.Equal(t, "END_OF_DATA", resp.Token)
	})

}
