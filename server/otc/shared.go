package otc

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
)

const (
	genericOTCAPIError = "Error when calling the OTC API. Please create a ticket"
	wrongAPIUsageError = "Invalid API request: Argument doesn't match definition. Please create a ticket."
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/otc/ecs", listECSHandler)
	r.POST("/otc/stopecs", stopECSHandler)
	r.POST("/otc/startecs", startECSHandler)
	r.POST("/otc/rebootecs", rebootECSHandler)
	r.GET("/otc/flavors", listFlavorsHandler)
	r.GET("/otc/images", listImagesHandler)
	r.GET("/otc/availabilityzones", listAvailabilityZonesHandler)
	r.GET("/otc/volumetypes", listVolumeTypesHandler)
	r.GET("/otc/rds/versions", listRDSVersionsHandler)
	r.GET("/otc/rds/flavors", listRDSFlavorsHandler)
	r.GET("/otc/rds/instances", listRDSInstancesHandler)
	//	r.GET("/otc/rds/instances", listRDSTagsHandler)
}

func getProvider(to *token.TokenOptions) (*gophercloud.ProviderClient, error) {
	opts, err := TokenOptionsFromEnv(to)
	if err != nil {
		return nil, err
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func getComputeClient() (*gophercloud.ServiceClient, error) {
	provider, err := getProvider(nil)
	if err != nil {
		fmt.Println("Error while authenticating.", err.Error())
		return nil, errors.New(genericOTCAPIError)
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "eu-ch",
	})

	if err != nil {
		fmt.Println("Error getting client.", err.Error())
		return nil, errors.New(genericOTCAPIError)
	}

	return client, nil
}

func getRDSClient() (*gophercloud.ServiceClient, error) {
	to := token.TokenOptions{
		TenantName: "eu-ch_rds",
	}
	provider, err := getProvider(&to)
	if err != nil {
		fmt.Println("Error while authenticating.", err.Error())
		return nil, errors.New(genericOTCAPIError)
	}

	client, err := openstack.NewRDSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("Error getting client.", err.Error())
		return nil, errors.New(genericOTCAPIError)
	}

	return client, nil
}

func getImageClient() (*gophercloud.ServiceClient, error) {
	provider, err := getProvider(nil)
	if err != nil {
		fmt.Println("Error while authenticating.", err.Error())
		return nil, errors.New(genericOTCAPIError)
	}

	client, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{
		Region: "eu-ch",
	})

	if err != nil {
		fmt.Println("Error getting client.", err.Error())
		return nil, errors.New(genericOTCAPIError)
	}

	return client, nil
}

func getBlockStorageClient() (*gophercloud.ServiceClient, error) {
	provider, err := getProvider(nil)
	if err != nil {
		fmt.Println("Error while authenticating.", err.Error())
		return nil, errors.New(genericOTCAPIError)
	}

	client, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: "eu-ch",
	})

	if err != nil {
		fmt.Println("Error getting client.", err.Error())
		return nil, errors.New(genericOTCAPIError)
	}

	return client, nil
}
