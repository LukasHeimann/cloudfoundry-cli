package running

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/net"

	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/securitygroups/defaults"
)

const urlPath = "/v2/config/running_security_groups"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SecurityGroupsRepo

type SecurityGroupsRepo interface {
	BindToRunningSet(string) error
	List() ([]models.SecurityGroupFields, error)
	UnbindFromRunningSet(string) error
}

type cloudControllerRunningSecurityGroupRepo struct {
	repoBase DefaultSecurityGroupsRepoBase
}

func NewSecurityGroupsRepo(configRepo coreconfig.Reader, gateway net.Gateway) SecurityGroupsRepo {
	return &cloudControllerRunningSecurityGroupRepo{
		repoBase: DefaultSecurityGroupsRepoBase{
			ConfigRepo: configRepo,
			Gateway:    gateway,
		},
	}
}

func (repo *cloudControllerRunningSecurityGroupRepo) BindToRunningSet(groupGUID string) error {
	return repo.repoBase.Bind(groupGUID, urlPath)
}

func (repo *cloudControllerRunningSecurityGroupRepo) List() ([]models.SecurityGroupFields, error) {
	return repo.repoBase.List(urlPath)
}

func (repo *cloudControllerRunningSecurityGroupRepo) UnbindFromRunningSet(groupGUID string) error {
	return repo.repoBase.Delete(groupGUID, urlPath)
}
