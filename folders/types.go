package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

// Added a PaginatedFolderResponse type hold paginated data
type FetchPaginatedFolderResponse struct {
	Folders map[string]folder_and_token_wrapper
}