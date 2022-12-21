package ptrArit

import (
	"fmt"
	"strconv"
)

const (
	ARIT_METHOD_ADD       = "add"
	ARIT_METHOD_MULTIPLY  = "multiply"
	ARIT_METHOD_SUBSTRACT = "substract"
	ARIT_METHOD_DIVIDE    = "divide"
)

type Number interface {
	float64 | float32 | int | int8 | int16 | int32 | int64
}

func Add[NumericT Number](v1 interface{}, v2 interface{}) *NumericT {
	return pointerArit[NumericT](v1, v2, ARIT_METHOD_ADD)
}

func Multiply[NumericT Number](v1 interface{}, v2 interface{}) *NumericT {
	return pointerArit[NumericT](v1, v2, ARIT_METHOD_MULTIPLY)
}

func Substract[NumericT Number](v1 interface{}, v2 interface{}) *NumericT {
	return pointerArit[NumericT](v1, v2, ARIT_METHOD_SUBSTRACT)
}

func Divide[NumericT Number](v1 interface{}, v2 interface{}) *NumericT {
	return pointerArit[NumericT](v1, v2, ARIT_METHOD_DIVIDE)
}

// pointerArit supports ArIT_METHOD_ Constants
func pointerArit[NumericT Number](v1 interface{}, v2 interface{}, method string) (typedResult *NumericT) {
	if v1 == nil || v2 == nil {
		return nil
	}

	v1F, err := convertToFloat64(v1)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}

	v2F, err2 := convertToFloat64(v2)
	if err2 != nil {
		fmt.Errorf(err2.Error())
		return nil
	}

	var result float64
	// handles add, substract, multiply and divide with switch
	switch method {
	case ARIT_METHOD_ADD:
		result = v1F + v2F
	case ARIT_METHOD_SUBSTRACT:
		result = v1F - v2F
	case ARIT_METHOD_MULTIPLY:
		result = v1F * v2F
	case ARIT_METHOD_DIVIDE:
		result = v1F / v2F
	}

	typedResult = ptr(NumericT(result))
	return
}

func convertToFloat64(v interface{}) (float64, error) {
	// TODO why can't I compare to Numeric? -> Make it clean
	switch v := v.(type) {
	case *int:
		return float64(*v), nil
	case *int16:
		return float64(*v), nil
	case *int32:
		return float64(*v), nil
	case *int64:
		return float64(*v), nil
	case *float64:
		return float64(*v), nil
	case *float32:
		return float64(*v), nil
	case *string:
		f, err := strconv.ParseFloat(*v, 64)
		if err != nil {
			return f, err
		}
		return f, nil
	case int:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return f, err
		}
		return f, nil
	default:
		return 0, fmt.Errorf("not able to parse value [%v] as float64", v)
	}
}
func ptr[T any](v T) *T {
	return &v
}
