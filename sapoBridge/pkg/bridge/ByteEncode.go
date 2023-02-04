package bridge

import (
	"errors"
	"fmt"
)

func unpack(bytes [32]byte) string {
	return fmt.Sprintf("%x", bytes)
}

// dummy specs for debugging
func pack(str string) (res [32]byte, err error) {
	if len(str) > 32 {
		err = errors.New("String is too large to be packed in bytes32")
		return
	}

	copy([]byte(str)[:], res[:32])
	return
}
