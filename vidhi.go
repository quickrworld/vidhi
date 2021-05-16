package main

import (
	"fmt"
	"strings"
	"vidhi/vidhi"
)

func Contains(args []vidhi.Arg) (bool, error) {
	if len(args) != 2 {
		return false, fmt.Errorf("args array must have 2 elements")
	}
	s, err := extractStringArg(args, 0)
	if err != nil {
		return false, err
	}
	substr, err := extractStringArg(args, 1)
	if err != nil {
		return false, err
	}
	return strings.Contains(s, substr), nil
}

func ContainsAny(args []vidhi.Arg) (bool, error) {
	if len(args) != 2 {
		return false, fmt.Errorf("args array must have 2 elements")
	}
	s, err := extractStringArg(args, 0)
	if err != nil {
		return false, err
	}
	chars, err := extractStringArg(args, 1)
	if err != nil {
		return false, err
	}
	return strings.ContainsAny(s, chars), nil
}

func HasPrefix(args []vidhi.Arg) (bool, error) {
	if len(args) != 2 {
		return false, fmt.Errorf("args array must have 2 elements")
	}
	s, err := extractStringArg(args, 0)
	if err != nil {
		return false, err
	}
	substr, err := extractStringArg(args, 1)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(s, substr), nil
}

func HasSuffix(args []vidhi.Arg) (bool, error) {
	if len(args) != 2 {
		return false, fmt.Errorf("args array must have 2 elements")
	}
	s, err := extractStringArg(args, 0)
	if err != nil {
		return false, err
	}
	substr, err := extractStringArg(args, 1)
	if err != nil {
		return false, err
	}
	return strings.HasSuffix(s, substr), nil
}

func Length(args []vidhi.Arg) (int, error) {
	if len(args) != 1 {
		return 0, fmt.Errorf("args array must have 1 element")
	}
	s, err := extractStringArg(args, 0)
	if err != nil {
		return 0, err
	}
	return len(s), nil
}

func extractStringArg(args []vidhi.Arg, i int) (string, error) {
	switch args[i].Value.(type) {
	case string:
		return args[i].Value.(string), nil
	default:
		return "", fmt.Errorf("%v arg is not a string", i)
	}
}


