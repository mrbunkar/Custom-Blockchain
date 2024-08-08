package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func (a Address) ToSlice() []byte {
	sl := make([]byte, 20)
	for i := 0; i < 20; i++ {
		sl[i] = a[i]
	}
	return sl
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("Given bytes with lenght %d shoudl be of lenghth 20\n", len(b))
		panic(msg)
	}

	var value [20]uint8

	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}

	return Address(value)
}

func (a Address) ToString() string {
	return hex.EncodeToString(a.ToSlice())
}
