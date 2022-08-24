package v7

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/sharedaction"
	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/v7action"
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3"
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/uaa"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/v7/shared"
	"code.cloudfoundry.org/clock"
)

type BaseCommand struct {
	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       Actor

	cloudControllerClient *ccv3.Client
	uaaClient             *uaa.Client
}

func (cmd *BaseCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	sharedActor := sharedaction.NewActor(config)
	cmd.SharedActor = sharedActor

	ccClient, uaaClient, routingClient, err := shared.GetNewClientsAndConnectToCF(config, ui, "")
	if err != nil {
		return err
	}
	cmd.cloudControllerClient = ccClient
	cmd.uaaClient = uaaClient

	cmd.Actor = v7action.NewActor(ccClient, config, sharedActor, uaaClient, routingClient, clock.NewClock())
	return nil
}

func (cmd *BaseCommand) GetClients() (*ccv3.Client, *uaa.Client) {
	return cmd.cloudControllerClient, cmd.uaaClient
}
