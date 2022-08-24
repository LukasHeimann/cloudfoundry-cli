package apifakes

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/errors"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
)

type OldFakeBuildpackBitsRepository struct {
	UploadBuildpackErr         bool
	UploadBuildpackAPIResponse error
	UploadBuildpackPath        string
}

func (repo *OldFakeBuildpackBitsRepository) UploadBuildpack(buildpack models.Buildpack, dir string) error {
	if repo.UploadBuildpackErr {
		return errors.New("Invalid buildpack")
	}

	repo.UploadBuildpackPath = dir
	return repo.UploadBuildpackAPIResponse
}
