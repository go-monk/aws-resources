package main

import (
	"fmt"
	"strings"
)

type Tag struct {
	Key   string
	Value string
}

type Tags []Tag

func (t *Tags) String() string {
	return fmt.Sprint(*t)
}

func (t *Tags) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("expected key=value, got %q", value)
	}
	*t = append(*t, Tag{Key: parts[0], Value: parts[1]})
	return nil
}
