package pathperms

import (
	"os"
	"strings"
)

type Permissions struct {
	bits os.FileMode
}

func (p Permissions) Owner() Permission {
	return Permission{bits: (p.bits >> OwnerShifter)}
}

func (p Permissions) Group() Permission {
	return Permission{bits: (p.bits >> GroupShifter) & 7}
}

func (p Permissions) Others() Permission {
	return Permission{bits: (p.bits >> OthersShifter) & 7}
}

func (p Permissions) String() string {
	return strings.TrimPrefix(p.bits.String(), "-")
}

func (p Permissions) Int() int {
	return int(p.bits)
}
