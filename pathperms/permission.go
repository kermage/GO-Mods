package pathperms

import (
	"os"
	"strings"
)

type Permission struct {
	bits os.FileMode
}

func (p Permission) HasAccess() bool {
	return p.bits > 0
}

func (p Permission) HasReadAccess() bool {
	return int(p.bits)&ReadDigit != 0
}

func (p Permission) HasWriteAccess() bool {
	return int(p.bits)&WriteDigit != 0
}

func (p Permission) HasExecuteAccess() bool {
	return int(p.bits)&ExecuteDigit != 0
}

func (p Permission) String() string {
	return strings.TrimPrefix(p.bits.String(), "-------")
}

func (p Permission) Int() int {
	return int(p.bits)
}
