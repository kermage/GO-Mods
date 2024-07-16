package pathperms

import (
	"os"
	"strings"
)

type Permissions struct {
	Bits os.FileMode
}

func (p Permissions) Owner() Permission {
	return Permission{Bits: (p.Bits >> OwnerShifter)}
}

func (p Permissions) Group() Permission {
	return Permission{Bits: (p.Bits >> GroupShifter) & 7}
}

func (p Permissions) Others() Permission {
	return Permission{Bits: (p.Bits >> OthersShifter) & 7}
}

func (p Permissions) String() string {
	return strings.TrimPrefix(p.Bits.String(), "-")
}

func (p Permissions) Int() int {
	return int(p.Bits)
}
