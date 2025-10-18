package hyprland

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type wanted interface {
	int | uint | string | bool
}

func cast[T wanted](s string) (T, error) {
	var zero T
	switch any(zero).(type) {
	case int:
		v, err := strconv.Atoi(s)
		return any(v).(T), err
	case uint:
		v, err := strconv.ParseUint(s, 10, 64)
		return any(uint(v)).(T), err
	case string:
		return any(s).(T), nil
	case bool:
		v, err := strconv.ParseBool(s)
		return any(v).(T), err
	default:
		return zero, errors.New("unsupported type")
	}
}

// cast2 splits by ',' and converts to two generic types
func cast2[T1, T2 wanted](s string) (T1, T2, error) {
	var zero1 T1
	var zero2 T2

	parts := strings.SplitN(s, ",", 2)
	if len(parts) < 2 {
		return zero1, zero2, errors.New(
			"need at least 2 comma-separated values",
		)
	}

	v1, err := cast[T1](strings.TrimSpace(parts[0]))
	if err != nil {
		return zero1, zero2, fmt.Errorf("convert 1: %w", err)
	}
	v2, err := cast[T2](strings.TrimSpace(parts[1]))
	if err != nil {
		return zero1, zero2, fmt.Errorf("convert 2: %w", err)
	}

	return v1, v2, nil
}

//revive:disable:function-result-limit

// cast3 splits by ',' and converts to three generic types
func cast3[T1, T2, T3 wanted](s string) (T1, T2, T3, error) {
	var zero1 T1
	var zero2 T2
	var zero3 T3

	parts := strings.SplitN(s, ",", 3)
	if len(parts) < 3 {
		return zero1, zero2, zero3, errors.New(
			"need at least 3 comma-separated values",
		)
	}

	v1, err := cast[T1](strings.TrimSpace(parts[0]))
	if err != nil {
		return zero1, zero2, zero3, fmt.Errorf("convert 1: %w", err)
	}
	v2, err := cast[T2](strings.TrimSpace(parts[1]))
	if err != nil {
		return zero1, zero2, zero3, fmt.Errorf("convert 2: %w", err)
	}
	v3, err := cast[T3](strings.TrimSpace(parts[2]))
	if err != nil {
		return zero1, zero2, zero3, fmt.Errorf("convert 3: %w", err)
	}

	return v1, v2, v3, nil
}

// cast4 splits by ',' and converts to three generic types
func cast4[T1, T2, T3, T4 wanted](s string) (T1, T2, T3, T4, error) {
	var zero1 T1
	var zero2 T2
	var zero3 T3
	var zero4 T4

	parts := strings.SplitN(s, ",", 3)
	if len(parts) < 3 {
		return zero1, zero2, zero3, zero4, errors.New(
			"need at least 4 comma-separated values",
		)
	}

	v1, err := cast[T1](strings.TrimSpace(parts[0]))
	if err != nil {
		return zero1, zero2, zero3, zero4, fmt.Errorf("convert 1: %w", err)
	}
	v2, err := cast[T2](strings.TrimSpace(parts[1]))
	if err != nil {
		return zero1, zero2, zero3, zero4, fmt.Errorf("convert 2: %w", err)
	}
	v3, err := cast[T3](strings.TrimSpace(parts[2]))
	if err != nil {
		return zero1, zero2, zero3, zero4, fmt.Errorf("convert 3: %w", err)
	}
	v4, err := cast[T4](strings.TrimSpace(parts[2]))
	if err != nil {
		return zero1, zero2, zero3, zero4, fmt.Errorf("convert 3: %w", err)
	}

	return v1, v2, v3, v4, nil
}

//revive:enable:function-result-limit
