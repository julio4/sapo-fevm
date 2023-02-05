package bridge

import (
	"errors"
	"fmt"
)

// Unpack a byte array into a string (ignore \x00)
func Unpack(bytes [32]byte) string {
	var res string = ""

	for _, b := range bytes {
		if b != 0 {
			res += string(b)
		}
	}

	return res
}

// Pack a string in a byte array. Returns an error
// if the string is longer than 32 characters
func Pack(str string) (res [32]byte, err error) {
	if len(str) > 32 {
		err = errors.New("String is too large to be packed in bytes32")
		return
	}

	bytes := []byte(str)
	copy(res[:32], bytes[:])
	return
}

func PackToHex(str string) (string, error) {
	bytes, err := Pack(str)

	if err != nil {
		return "", err
	}

	res := "0x"

	for _, b := range bytes {
		res += fmt.Sprintf("%02x", b)
	}

	return res, nil
}
