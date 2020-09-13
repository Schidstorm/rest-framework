package rest_framework

import (
	"errors"
	"fmt"
	"strconv"
)

type QueryArguments struct {
	arguments map[string][]string
}

func NewQueryArguments(arguments map[string][]string) *QueryArguments {
	return &QueryArguments{arguments: arguments}
}

func (q *QueryArguments) GetString(key string) (string, error) {
	if arg, ok := q.arguments[key]; ok {
		return arg[0], nil
	}
	return "", errors.New(fmt.Sprintf("argument %s not found", key))
}

func (q *QueryArguments) GetInt32(key string, base int) (int32, error) {
	arg, err := q.GetVariableInt(key, base, 32)
	if err != nil {
		return 0, err
	}
	return int32(arg), nil
}

func (q *QueryArguments) GetInt64(key string, base int) (int64, error) {
	arg, err := q.GetVariableInt(key, base, 64)
	if err != nil {
		return 0, err
	}
	return arg, nil
}

func (q *QueryArguments) GetVariableInt(key string, base, bitSize int) (int64, error) {
	arg, err := q.GetString(key)
	if err != nil {
		return 0, err
	}
	argInt, err := strconv.ParseInt(arg, base, bitSize)
	if err != nil {
		return 0, err
	}
	return argInt, nil
}
