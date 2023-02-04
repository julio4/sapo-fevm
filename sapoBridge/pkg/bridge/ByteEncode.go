package bridge

import (
	"errors"
)

func unpack(bytes []byte) string {
	return string(bytes)
}

// dummy specs for debugging
func pack(str string) ([]byte, error) {
	if len(str) > 32 {
		return nil, errors.New("String is too large to be packed in bytes32")
	}

	return []byte(str), nil
}
