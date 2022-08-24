package commandsloader

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/application"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/buildpack"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/domain"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/environmentvariablegroup"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/featureflag"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/organization"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/plugin"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/pluginrepo"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/quota"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/route"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/routergroups"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/securitygroup"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/service"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/serviceaccess"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/serviceauthtoken"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/servicebroker"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/servicekey"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/space"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/spacequota"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/user"
)

/*******************
This package make a reference to all the command packages
in cf/commands/..., so all init() in the directories will
get initialized

* Any new command packages must be included here for init() to get called
********************/

func Load() {
	_ = commands.API{}
	_ = application.ListApps{}
	_ = buildpack.ListBuildpacks{}
	_ = domain.CreateDomain{}
	_ = environmentvariablegroup.RunningEnvironmentVariableGroup{}
	_ = featureflag.ShowFeatureFlag{}
	_ = organization.ListOrgs{}
	_ = plugin.Plugins{}
	_ = pluginrepo.RepoPlugins{}
	_ = quota.CreateQuota{}
	_ = route.CreateRoute{}
	_ = routergroups.RouterGroups{}
	_ = securitygroup.ShowSecurityGroup{}
	_ = service.ShowService{}
	_ = serviceauthtoken.ListServiceAuthTokens{}
	_ = serviceaccess.ServiceAccess{}
	_ = servicebroker.ListServiceBrokers{}
	_ = servicekey.ServiceKey{}
	_ = space.CreateSpace{}
	_ = spacequota.SpaceQuota{}
	_ = user.CreateUser{}
}
