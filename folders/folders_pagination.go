package folders

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// Pagination parameters added to the request struct
type PaginatedFetchFolderRequest struct {
	OrgID uuid.UUID
	Token string // pagination token
	Limit int    // Limit is the number of items per page
}

// Pagination token and data added to the response struct
type PaginatedFetchFolderResponse struct {
	Folders []*Folder
	Token   string // token for the next page
}

// Pagination logic in GetAllFolders
func GetPaginatedAllFolders(req *PaginatedFetchFolderRequest) (*PaginatedFetchFolderResponse, error) {

	// Check for end of data
	if req.Token == "END_OF_DATA" {
		return generateEmptyResponse(), nil
	}

	// Use GetAllFolders to fetch all folders
	nonPaginatedResponse, err := GetAllFolders(&FetchFolderRequest{OrgID: req.OrgID})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all folders: %w", err)
	}
	allFolders := nonPaginatedResponse.Folders

	// Paginate folders
	foldersPage, nextToken, err := paginateFolders(allFolders, req.Token, req.Limit)
	if err != nil {
		return nil, err
	}

	// Return the paginated response
	return &PaginatedFetchFolderResponse{
		Folders: foldersPage,
		Token:   nextToken,
	}, nil
}

func generateEmptyResponse() *PaginatedFetchFolderResponse {
	return &PaginatedFetchFolderResponse{
		Folders: []*Folder{},
		Token:   "END_OF_DATA",
	}
}

func paginateFolders(allFolders []*Folder, token string, limit int) ([]*Folder, string, error) {
	startingAfter, err := ParsePaginationToken(token)
	if err != nil {
		return nil, "", err
	}

	startIndex, endIndex := calculatePaginationIndexes(allFolders, startingAfter, limit)

	foldersPage := allFolders[startIndex:endIndex]
	nextToken := generateNextToken(foldersPage, endIndex, len(allFolders))

	return foldersPage, nextToken, nil
}

func calculatePaginationIndexes(allFolders []*Folder, startingAfter *PaginationTokenStruct, limit int) (int, int) {
	startIndex := 0
	if startingAfter != nil {
		for i, folder := range allFolders {
			if folder.Id == startingAfter.LastID {
				startIndex = i + 1
				break
			}
		}
	}

	endIndex := startIndex + limit
	if endIndex > len(allFolders) {
		endIndex = len(allFolders)
	}

	return startIndex, endIndex
}

func generateNextToken(foldersPage []*Folder, endIndex, totalFolders int) string {
	if endIndex >= totalFolders {
		return "END_OF_DATA"
	}
	lastFolder := foldersPage[len(foldersPage)-1]
	return GeneratePaginationToken(lastFolder, false)
}
