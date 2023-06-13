package longtermretentionbackups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupStorageRedundancy string

const (
	BackupStorageRedundancyGeo   BackupStorageRedundancy = "Geo"
	BackupStorageRedundancyLocal BackupStorageRedundancy = "Local"
	BackupStorageRedundancyZone  BackupStorageRedundancy = "Zone"
)

func PossibleValuesForBackupStorageRedundancy() []string {
	return []string{
		string(BackupStorageRedundancyGeo),
		string(BackupStorageRedundancyLocal),
		string(BackupStorageRedundancyZone),
	}
}

func (s *BackupStorageRedundancy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupStorageRedundancy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupStorageRedundancy(input string) (*BackupStorageRedundancy, error) {
	vals := map[string]BackupStorageRedundancy{
		"geo":   BackupStorageRedundancyGeo,
		"local": BackupStorageRedundancyLocal,
		"zone":  BackupStorageRedundancyZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupStorageRedundancy(input)
	return &out, nil
}

type DatabaseState string

const (
	DatabaseStateAll     DatabaseState = "All"
	DatabaseStateDeleted DatabaseState = "Deleted"
	DatabaseStateLive    DatabaseState = "Live"
)

func PossibleValuesForDatabaseState() []string {
	return []string{
		string(DatabaseStateAll),
		string(DatabaseStateDeleted),
		string(DatabaseStateLive),
	}
}

func (s *DatabaseState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseState(input string) (*DatabaseState, error) {
	vals := map[string]DatabaseState{
		"all":     DatabaseStateAll,
		"deleted": DatabaseStateDeleted,
		"live":    DatabaseStateLive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseState(input)
	return &out, nil
}

type RequestedBackupStorageRedundancy string

const (
	RequestedBackupStorageRedundancyGeo   RequestedBackupStorageRedundancy = "Geo"
	RequestedBackupStorageRedundancyLocal RequestedBackupStorageRedundancy = "Local"
	RequestedBackupStorageRedundancyZone  RequestedBackupStorageRedundancy = "Zone"
)

func PossibleValuesForRequestedBackupStorageRedundancy() []string {
	return []string{
		string(RequestedBackupStorageRedundancyGeo),
		string(RequestedBackupStorageRedundancyLocal),
		string(RequestedBackupStorageRedundancyZone),
	}
}

func (s *RequestedBackupStorageRedundancy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRequestedBackupStorageRedundancy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRequestedBackupStorageRedundancy(input string) (*RequestedBackupStorageRedundancy, error) {
	vals := map[string]RequestedBackupStorageRedundancy{
		"geo":   RequestedBackupStorageRedundancyGeo,
		"local": RequestedBackupStorageRedundancyLocal,
		"zone":  RequestedBackupStorageRedundancyZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RequestedBackupStorageRedundancy(input)
	return &out, nil
}

type TargetBackupStorageRedundancy string

const (
	TargetBackupStorageRedundancyGeo   TargetBackupStorageRedundancy = "Geo"
	TargetBackupStorageRedundancyLocal TargetBackupStorageRedundancy = "Local"
	TargetBackupStorageRedundancyZone  TargetBackupStorageRedundancy = "Zone"
)

func PossibleValuesForTargetBackupStorageRedundancy() []string {
	return []string{
		string(TargetBackupStorageRedundancyGeo),
		string(TargetBackupStorageRedundancyLocal),
		string(TargetBackupStorageRedundancyZone),
	}
}

func (s *TargetBackupStorageRedundancy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetBackupStorageRedundancy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetBackupStorageRedundancy(input string) (*TargetBackupStorageRedundancy, error) {
	vals := map[string]TargetBackupStorageRedundancy{
		"geo":   TargetBackupStorageRedundancyGeo,
		"local": TargetBackupStorageRedundancyLocal,
		"zone":  TargetBackupStorageRedundancyZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetBackupStorageRedundancy(input)
	return &out, nil
}
