package folders

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
)

type PaginationTokenStruct struct {
	LastID uuid.UUID `json:"lastId"`
}

// Helper function to generate a pagination token from the last folder in a page
func GeneratePaginationToken(lastFolder *Folder, isLastPage bool) string {
	// If lastFolder is nil, return an empty string (no more data to paginate)
	if lastFolder == nil || isLastPage {
		return "END_OF_DATA"
	}

	// Generate token based on the last folder's ID
	tokenStruct := PaginationTokenStruct{
		LastID: lastFolder.Id,
	}
	tokenBytes, err := json.Marshal(tokenStruct)
	if err != nil {
		fmt.Printf("Error generating pagination token: %v\n", err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(tokenBytes)
}

// Helper function to parse a pagination token back into the last ID
func ParsePaginationToken(token string) (*PaginationTokenStruct, error) {
	if token == "" {
		return nil, nil // No cursor provided, start from the beginning
	}

	tokenBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	var tokenStruct PaginationTokenStruct
	err = json.Unmarshal(tokenBytes, &tokenStruct)
	if err != nil {
		return nil, err
	}

	return &tokenStruct, nil
}
