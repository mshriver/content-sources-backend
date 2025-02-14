package external_repos

import (
	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/dao"
	"github.com/content-services/yummy/pkg/yum"
	"github.com/stretchr/testify/mock"
)

// TestIntrospect

type MockRpmDao struct {
}

func (m MockRpmDao) List(orgID string, uuidRepo string, limit int, offset int, search string, sortBy string) (api.RepositoryRpmCollectionResponse, int64, error) {
	return api.RepositoryRpmCollectionResponse{}, 0, nil
}

func (m MockRpmDao) InsertForRepository(repoUuid string, pkgs []yum.Package) (int64, error) {
	return int64(len(pkgs)), nil
}

func (m MockRpmDao) Search(orgID string, request api.SearchRpmRequest) ([]api.SearchRpmResponse, error) {
	return []api.SearchRpmResponse{}, nil
}

type MockRepositoryDao struct {
	mock.Mock
}

func (m *MockRepositoryDao) List() ([]dao.Repository, error) {
	return []dao.Repository{}, nil
}

func (m *MockRepositoryDao) FetchForUrl(url string) (dao.Repository, error) {
	return dao.Repository{}, nil
}

func (m *MockRepositoryDao) Update(repo dao.RepositoryUpdate) error {
	args := m.Called(repo)
	return args.Error(0)
}

func (m MockRpmDao) OrphanCleanup() error {
	return nil
}
