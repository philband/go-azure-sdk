package dscpconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscpConfigurationClient struct {
	Client *resourcemanager.Client
}

func NewDscpConfigurationClientWithBaseURI(api environments.Api) (*DscpConfigurationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "dscpconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DscpConfigurationClient: %+v", err)
	}

	return &DscpConfigurationClient{
		Client: client,
	}, nil
}
