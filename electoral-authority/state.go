package electoralauthority

import "strings"

func SplitKey(key string) []string {
	return strings.Split(key, ":")
}

func MakeKey(keyParts ...string) string {
	return strings.Join(keyParts, ":")
}

type StateInterface interface {
	GetSplitKey() []string
	Serialize() ([]byte, error)
}
