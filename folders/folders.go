package folders

import (
	"github.com/gofrs/uuid"
)

/*
CODE:
This function retrieves all the folders that satisfies the 'req' FetchFolderRequest
and returns the required Folders in a slice within a FetchFolderResponse type

The FetchFolderRequest type holds an OrgId, so the FetchAllFoldersByOrgId function is called
which fetches all the folders that match the given OrgId from req into a slice of Folder pointers

IMPROVEMENT SUGGESTIONS:

  - Could remove the starting variable initialisation block since it looks like they are unused

  - Could use variable + assignment operator (:=) instead of declaring "var ffr *FetchFolderResponse" in one line and then
    assigning in another line e.g. ffr := &FetchFolderResponse{Folders: fp}

  - The second for loop appears to just revert what the first for loop had done
    as it converts the dereferenced pointer, back into a pointer again.
    Instead we could just remove the loops and at the end have
    ffr := &FetchFolderResponse{Folders: r}
    and return ffr
*/
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	for k, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

/*
Finds all the folders that match the given orgID, and returns a slice of pointers
to those folders
*/
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
