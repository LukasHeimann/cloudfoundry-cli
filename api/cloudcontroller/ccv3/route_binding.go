package ccv3

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/internal"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

func (client *Client) CreateRouteBinding(binding resources.RouteBinding) (JobURL, Warnings, error) {
	return client.MakeRequest(RequestParams{
		RequestName: internal.PostRouteBindingRequest,
		RequestBody: binding,
	})
}

func (client *Client) GetRouteBindings(query ...Query) ([]resources.RouteBinding, IncludedResources, Warnings, error) {
	var result []resources.RouteBinding

	included, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetRouteBindingsRequest,
		Query:        query,
		ResponseBody: resources.RouteBinding{},
		AppendToList: func(item interface{}) error {
			result = append(result, item.(resources.RouteBinding))
			return nil
		},
	})

	return result, included, warnings, err
}

func (client *Client) DeleteRouteBinding(guid string) (JobURL, Warnings, error) {
	return client.MakeRequest(RequestParams{
		RequestName: internal.DeleteRouteBindingRequest,
		URIParams:   internal.Params{"route_binding_guid": guid},
	})
}
