package replicationrecoveryplans

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPlanProviderSpecificDetails interface {
}

// RawRecoveryPlanProviderSpecificDetailsImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRecoveryPlanProviderSpecificDetailsImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalRecoveryPlanProviderSpecificDetailsImplementation(input []byte) (RecoveryPlanProviderSpecificDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryPlanProviderSpecificDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "A2A") {
		var out RecoveryPlanA2ADetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecoveryPlanA2ADetails: %+v", err)
		}
		return out, nil
	}

	out := RawRecoveryPlanProviderSpecificDetailsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
