package cache

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheUpdateProperties struct {
	ConnectionString *string `json:"connectionString,omitempty"`
	Description      *string `json:"description,omitempty"`
	ResourceId       *string `json:"resourceId,omitempty"`
	UseFromLocation  *string `json:"useFromLocation,omitempty"`
}
