package commands

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/authentication"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/terminal"
	"github.com/LukasHeimann/cloudfoundrycli/v8/plugin/models"
)

type OAuthToken struct {
	ui          terminal.UI
	config      coreconfig.ReadWriter
	authRepo    authentication.Repository
	pluginModel *plugin_models.GetOauthToken_Model
	pluginCall  bool
}

func init() {
	commandregistry.Register(&OAuthToken{})
}

func (cmd *OAuthToken) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "oauth-token",
		Description: T("Retrieve and display the OAuth token for the current session"),
		Usage: []string{
			T("CF_NAME oauth-token"),
		},
	}
}

func (cmd *OAuthToken) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
	}

	return reqs, nil
}

func (cmd *OAuthToken) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.authRepo = deps.RepoLocator.GetAuthenticationRepository()
	cmd.pluginCall = pluginCall
	cmd.pluginModel = deps.PluginModels.OauthToken
	return cmd
}

func (cmd *OAuthToken) Execute(c flags.FlagContext) error {
	token, err := cmd.authRepo.RefreshAuthToken()
	if err != nil {
		return err
	}

	if cmd.pluginCall {
		cmd.pluginModel.Token = token
	} else {
		cmd.ui.Say(token)
	}
	return nil
}
