package folders

// This'github.com/gofrs/uuid' is a package imported.
import (
	"github.com/gofrs/uuid"
)

// My sugestions for improvement the code:
// fix issues that prevent the code from running - such as removing unused variables
// have better naming conventions as its quite confusing at the moment
// try to run the code after that
// add proper error handling
// create unit tests for the functions
// the first loop seems to be unnecessary as its just copying the data from a pointer to a value and then back to a pointer again - I believe we can remove this for improving performance
// the second loop also seems unnecessary and it seems to be wrong as it reuses the loop variable address which can leap to all elements in 'fp' pointing to the same instance

// This function retrieves all folders related to an organization ID returning in the response and error.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Fetch all folders by organization ID and handle any errors
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		// Return the error to the caller
		return nil, err
	}

	ffr := &FetchFolderResponse{Folders: folders}
	return ffr, nil
}

// This function retrieves all Folder instances that match an organization ID.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() // Calls function that return a sample set of folder data.

	resFolder := []*Folder{} // Initializes a slice to hold points to the Folder structs.

	// Iterates over folders.
	for _, folder := range folders {
		if folder.OrgId == orgID { // Checks if the folder's organization ID matches the provided 'orgID'.
			resFolder = append(resFolder, folder) // Appends the pointer to the matching Folder to 'resFolder'.
		}
	}
	//Returns the filtered slice and a nil error
	return resFolder, nil
}
