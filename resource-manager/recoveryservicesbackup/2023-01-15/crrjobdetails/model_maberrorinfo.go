package crrjobdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MabErrorInfo struct {
	ErrorString     *string   `json:"errorString,omitempty"`
	Recommendations *[]string `json:"recommendations,omitempty"`
}
