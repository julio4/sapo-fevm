package bridge

import (
	"errors"
)

func unpack(bytes [32]byte) string {
	var res string = ""

	for _, b := range bytes {
		if b != 0 {
			res += string(b)
		}
	}

	return res
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
