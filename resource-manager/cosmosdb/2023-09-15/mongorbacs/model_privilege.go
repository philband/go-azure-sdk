package mongorbacs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Privilege struct {
	Actions  *[]string          `json:"actions,omitempty"`
	Resource *PrivilegeResource `json:"resource,omitempty"`
}
