package tf6muxserver

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// dynamicValueEquals performs equality checking of DynamicValue.
func dynamicValueEquals(schemaType tftypes.Type, i *tfprotov6.DynamicValue, j *tfprotov6.DynamicValue) (bool, error) {
	if i == nil {
		return j == nil, nil
	}

	if j == nil {
		return false, nil
	}

	// Upstream will panic on DynamicValue.Unmarshal with nil Type
	if schemaType == nil {
		return false, fmt.Errorf("unable to unmarshal DynamicValue: missing Type")
	}

	iValue, err := i.Unmarshal(schemaType)

	if err != nil {
		return false, fmt.Errorf("unable to unmarshal DynamicValue: %w", err)
	}

	jValue, err := j.Unmarshal(schemaType)

	if err != nil {
		return false, fmt.Errorf("unable to unmarshal DynamicValue: %w", err)
	}

	return iValue.Equal(jValue), nil
}
