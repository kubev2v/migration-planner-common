package engines

import (
	"fmt"
	"math"

	"github.com/kubev2v/migration-planner-common/estimations/estimation"
	"github.com/kubev2v/migration-planner-common/estimations/estimation/calculators"
)

type parameterDefinition struct {
	Key     string
	Type    string
	Default any
	Min     *float64
	Max     *float64
}

var estimationParamDefs = func() []parameterDefinition {
	minRate := 0.1
	minHours := 0.5
	minMins := 1.0
	minEngineers := 1.0
	return []parameterDefinition{
		{Key: calculators.ParamTransferRateMbps, Type: "number", Default: calculators.DefaultTransferRateMbps, Min: &minRate},
		{Key: calculators.ParamWorkHoursPerDay, Type: "number", Default: calculators.DefaultWorkHoursPerDay, Min: &minHours},
		{Key: calculators.ParamTroubleshootMinsPerVM, Type: "number", Default: calculators.DefaultTroubleshootMinsPerVM, Min: &minMins},
		{Key: calculators.ParamPostMigrationEngineers, Type: "integer", Default: calculators.DefaultEngineerCount, Min: &minEngineers},
	}
}()

// ValidateParams validates user settings against estimationParamDefs.
// Inventory-derived inputs, such as VM counts and disk sizes, cannot be overridden.
func ValidateParams(params []estimation.Param) error {
	definitions := make(map[string]parameterDefinition, len(estimationParamDefs))
	for _, def := range estimationParamDefs {
		definitions[def.Key] = def
	}

	for _, param := range params {
		def, known := definitions[param.Key]
		if !known {
			return fmt.Errorf("unknown param key %q", param.Key)
		}

		number, err := toFloat64(param.Value)
		if err != nil {
			return fmt.Errorf("param %q: value must be numeric", param.Key)
		}
		if def.Type == "integer" && math.Trunc(number) != number {
			return fmt.Errorf("param %q: value %v must be a whole number", param.Key, number)
		}
		if def.Min != nil && number < *def.Min {
			return fmt.Errorf("param %q: value %v is below minimum %v", param.Key, number, *def.Min)
		}
		if def.Max != nil && number > *def.Max {
			return fmt.Errorf("param %q: value %v exceeds maximum %v", param.Key, number, *def.Max)
		}
	}
	return nil
}

// BuildBaseParams merges defaults with validated user settings.
func BuildBaseParams(overrides ...estimation.Param) []estimation.Param {
	merged := make(map[string]estimation.Param, len(estimationParamDefs))
	for _, def := range estimationParamDefs {
		merged[def.Key] = estimation.Param{Key: def.Key, Value: def.Default}
	}
	for _, param := range overrides {
		merged[param.Key] = param
	}

	parameters := make([]estimation.Param, 0, len(merged))
	for _, param := range merged {
		parameters = append(parameters, param)
	}
	return parameters
}

func toFloat64(value any) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	default:
		return 0, fmt.Errorf("not numeric")
	}
}
