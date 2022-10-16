package folders

import (
	"github.com/gofrs/uuid"
	"github.com/thanhpk/randstr"
)

type folder_and_token_wrapper struct {
	Folder_list []*Folder
	Token string
}

func GetAllFoldersPaginated(req *FetchFolderRequest) (*FetchPaginatedFolderResponse, error) {
	r, _ := FetchAllFoldersByOrgIDPaginated(req.OrgID)
	paginated_data := PaginateData(r)
	ffr := &FetchPaginatedFolderResponse{Folders: paginated_data}
	return ffr, nil
}

func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}


/*
SOLUTION EXPLANATION:
My solution for paginating the data was to use 
a map data structure to split the data chunks where
tokens are used as keys, and the value is a struct
containing a small slice of data on the folders and a
token indicating the key for the next chunk of data.
The starting token is the string "start", and the 
final token is the string "null".

SOLUTION JUSTIFICATION:
I have gone with this approach as a map provides 
a simple data structure for splitting the folder data
into smaller chunks, which can all be accessed by tokens. 
The tokens for the next data chunk can also be
easily accessed by the token element of the "folder_and_token_wrapper"
struct stored that is stored as values in the map 
*/
func PaginateData(folders []*Folder) map[string]folder_and_token_wrapper {

	// Put the paginated data into a map, using tokens as keys
	paginated_data := make(map[string]folder_and_token_wrapper)

	// Place initial 10 folder elements into a page
	temp_folder_list := []*Folder{}
	for i := 0; i < 10; i++ {
		temp_folder_list = append(temp_folder_list, folders[i])
	}
	
	prev_token := "start"
	for i := 10; i < len(folders); i++ {
		temp_folder_list = append(temp_folder_list, folders[i])

		if (i == len(folders) - 1) { // For final folder element
			sep_data := folder_and_token_wrapper{
				Folder_list: temp_folder_list,
				Token: "null",
			}
			paginated_data[prev_token] = sep_data
			temp_folder_list = nil
			break
		}

		if (i % 10 == 0) { // Split into pages of 10
			sep_data := folder_and_token_wrapper{
				Folder_list: temp_folder_list,
				Token: randstr.String(5),
			}
			paginated_data[prev_token] = sep_data
			prev_token = sep_data.Token
			temp_folder_list = nil
		}

	}
	return paginated_data
}

/*
[
	Folder type struct 0 -> {
		id: "11130579-bb38-4317-a764-2efff98cbb04"
		name: "key-kling-conbra"
		org_id: "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
		deleted: true
	},
	{

	}
]
*/

/*
EXAMPLE OUTPUT OF PAGINATED DATA
{
	"start": {
		data: [
			Folder type struct 0 -> {
				id: "11130579-bb38-4317-a764-2efff98cbb04"
				name: "key-kling-conbra"
				org_id: "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
				deleted: true
			},
			Folder type struct 1 -> {
				...
			}
		],
		token: "Mw=="
	},
	"Mw==": {
		data: [
			Folder type struct 0 -> {
				id: "11130579-bb38-4317-a764-2efff98cbb04"
				name: "key-kling-conbra"
				org_id: "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
				deleted: true
			},
			Folder type struct 1 -> {
				...
			}
		],
		token: "ASDFG"
	}

}
*/