package monitoredsubscriptions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoredSubscriptionsClient struct {
	Client *resourcemanager.Client
}

func NewMonitoredSubscriptionsClientWithBaseURI(sdkApi sdkEnv.Api) (*MonitoredSubscriptionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "monitoredsubscriptions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonitoredSubscriptionsClient: %+v", err)
	}

	return &MonitoredSubscriptionsClient{
		Client: client,
	}, nil
}
