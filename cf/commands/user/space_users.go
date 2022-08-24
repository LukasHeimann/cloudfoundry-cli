package user

import (
	"fmt"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/actors/userprint"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/spaces"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/terminal"
	"github.com/LukasHeimann/cloudfoundrycli/v8/plugin/models"
)

type SpaceUsers struct {
	ui          terminal.UI
	config      coreconfig.Reader
	spaceRepo   spaces.SpaceRepository
	userRepo    api.UserRepository
	orgReq      requirements.OrganizationRequirement
	pluginModel *[]plugin_models.GetSpaceUsers_Model
	pluginCall  bool
}

func init() {
	commandregistry.Register(&SpaceUsers{})
}

func (cmd *SpaceUsers) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "space-users",
		Description: T("Show space users by role"),
		Usage: []string{
			T("CF_NAME space-users ORG SPACE"),
		},
	}
}

func (cmd *SpaceUsers) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 2 {
		cmd.ui.Failed(T("Incorrect Usage. Requires arguments\n\n") + commandregistry.Commands.CommandUsage("space-users"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 2)
	}

	cmd.orgReq = requirementsFactory.NewOrganizationRequirement(fc.Args()[0])

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
		cmd.orgReq,
	}

	return reqs, nil
}

func (cmd *SpaceUsers) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.userRepo = deps.RepoLocator.GetUserRepository()
	cmd.spaceRepo = deps.RepoLocator.GetSpaceRepository()
	cmd.pluginCall = pluginCall
	cmd.pluginModel = deps.PluginModels.SpaceUsers

	return cmd
}

func (cmd *SpaceUsers) Execute(c flags.FlagContext) error {
	spaceName := c.Args()[1]
	org := cmd.orgReq.GetOrganization()

	space, err := cmd.spaceRepo.FindByNameInOrg(spaceName, org.GUID)
	if err != nil {
		return err
	}

	printer := cmd.printer(org, space, cmd.config.Username())
	printer.PrintUsers(space.GUID, cmd.config.Username())
	return nil
}

func (cmd *SpaceUsers) printer(org models.Organization, space models.Space, username string) userprint.UserPrinter {
	var roles = []models.Role{models.RoleSpaceManager, models.RoleSpaceDeveloper, models.RoleSpaceAuditor}

	if cmd.pluginCall {
		return userprint.NewSpaceUsersPluginPrinter(
			cmd.pluginModel,
			cmd.userRepo.ListUsersInSpaceForRoleWithNoUAA,
			roles,
		)
	}

	cmd.ui.Say(T("Getting users in org {{.TargetOrg}} / space {{.TargetSpace}} as {{.CurrentUser}}",
		map[string]interface{}{
			"TargetOrg":   terminal.EntityNameColor(org.Name),
			"TargetSpace": terminal.EntityNameColor(space.Name),
			"CurrentUser": terminal.EntityNameColor(username),
		}))

	return &userprint.SpaceUsersUIPrinter{
		UI:         cmd.ui,
		UserLister: cmd.userRepo.ListUsersInSpaceForRoleWithNoUAA,
		Roles:      roles,
		RoleDisplayNames: map[models.Role]string{
			models.RoleSpaceManager:   T("SPACE MANAGER"),
			models.RoleSpaceDeveloper: T("SPACE DEVELOPER"),
			models.RoleSpaceAuditor:   T("SPACE AUDITOR"),
		},
	}
}
