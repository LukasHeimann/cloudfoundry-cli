package v7action

import "github.com/LukasHeimann/cloudfoundrycli/v8/actor/sharedaction"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SSHActor

type SSHActor interface {
	ExecuteSecureShell(sshOptions sharedaction.SSHOptions) error
}
