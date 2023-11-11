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

// The function GetAllFolders takes a FetchFolderRequest as a parameter and returns a FetchFolderResponse and an error.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	// It set the variable for err (is error), f1(is Folder ) and fs (is a slice of Folder pointer).
	var (
		err error
		f1  Folder
		fs  []*Folder
	)

	// It set the variable named f is an empty slice.
	f := []Folder{}

	// It call the function FetchAllFoldersByOrgID with OrgId provided in req parameter.
	r, _ := FetchAllFoldersByOrgID(req.OrgID)

	// With a for loop, r(is the result of FetchAllFoldersByOrgID) is iterated and appended each folder to the f slice.
	for k, v := range r {
		f = append(f, *v)
	}

	// It set the variable named fp that is an empty slice of Folder pointers.
	var fp []*Folder

	// It appends pointers to the folders from f to fp.
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	// It set the variable named ffr is the FetchFolderResponse pointers and assigns it a value by creating a FetchFolderResponse struct with the fp slice.
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	// It returns ffr and a nil error, indicating a successful execution of the function.
	return ffr, nil
}

// The function FetchAllFoldersByOrgID takes an orgID of type UUID as a parameter and returns a slice of Folder pointers and an error.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {

	// It set the variable named folders contains the function GetSampleData.
	folders := GetSampleData()

	// It set the variable named resFolder is the slice of Folder pointers.
	resFolder := []*Folder{}

	// While folders iterated, if orgId belong to folder equale to orgID, it appends folder to resFolder.
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	// It returns resFolder and nil.
	return resFolder, nil
}
