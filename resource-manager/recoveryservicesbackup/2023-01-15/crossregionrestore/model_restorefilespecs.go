package crossregionrestore

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreFileSpecs struct {
	FileSpecType     *string `json:"fileSpecType,omitempty"`
	Path             *string `json:"path,omitempty"`
	TargetFolderPath *string `json:"targetFolderPath,omitempty"`
}
