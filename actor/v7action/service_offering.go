package v7action

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/actionerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/railway"
)

func (actor Actor) PurgeServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName string) (Warnings, error) {
	var serviceOffering resources.ServiceOffering

	warnings, err := railway.Sequentially(
		func() (warnings ccv3.Warnings, err error) {
			serviceOffering, warnings, err = actor.CloudControllerClient.GetServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName)
			err = actionerror.EnrichAPIErrors(err)
			return
		},
		func() (ccv3.Warnings, error) {
			return actor.CloudControllerClient.PurgeServiceOffering(serviceOffering.GUID)
		},
	)

	return Warnings(warnings), err
}
