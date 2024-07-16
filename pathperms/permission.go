package pathperms

import (
	"os"
	"strings"
)

type Permission struct {
	Bits os.FileMode
}

func (p Permission) HasAccess() bool {
	return p.Bits > 0
}

func (p Permission) HasReadAccess() bool {
	return int(p.Bits)&ReadDigit != 0
}

func (p Permission) HasWriteAccess() bool {
	return int(p.Bits)&WriteDigit != 0
}

func (p Permission) HasExecuteAccess() bool {
	return int(p.Bits)&ExecuteDigit != 0
}

func (p Permission) String() string {
	return strings.TrimPrefix(p.Bits.String(), "-------")
}

func (p Permission) Int() int {
	return int(p.Bits)
}
