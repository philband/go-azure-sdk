package cloudservicepublicipaddresses

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayLoadBalancerTunnelInterface struct {
	Identifier *int64                                  `json:"identifier,omitempty"`
	Port       *int64                                  `json:"port,omitempty"`
	Protocol   *GatewayLoadBalancerTunnelProtocol      `json:"protocol,omitempty"`
	Type       *GatewayLoadBalancerTunnelInterfaceType `json:"type,omitempty"`
}
