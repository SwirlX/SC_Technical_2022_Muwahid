package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {

	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}
	var response *folders.FetchFolderResponse
	response, _ = folders.GetAllFolders(req)
	f := response.Folders

	t.Run("test1", func(t *testing.T) {
		f0 := f[0]
		assert.Equal(t, uuid.FromStringOrNil("11130579-bb38-4317-a764-2efff98cbb04"), f0.Id)
		assert.Equal(t, "key-king-cobra", f0.Name)
		assert.Equal(t, uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), f0.OrgId)
		assert.Equal(t, false, f0.Deleted)
	})

	t.Run("test2", func(t *testing.T) {
		f1 := f[1]
		assert.Equal(t, uuid.FromStringOrNil("72dc6a27-14d5-4643-a882-087990903069"), f1.Id)
		assert.Equal(t, "balanced-bella", f1.Name)
		assert.Equal(t, uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), f1.OrgId)
		assert.Equal(t, true, f1.Deleted)
	})

	t.Run("test3", func(t *testing.T) {
		fn := f[len(f) - 1]
		assert.Equal(t, uuid.FromStringOrNil("96639516-0743-413c-8f53-2cc922e250ba"), fn.Id)
		assert.Equal(t, "correct-scalphunter", fn.Name)
		assert.Equal(t, uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), fn.OrgId)
		assert.Equal(t, true, fn.Deleted)
	})

	

}
