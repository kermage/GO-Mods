package pathperms

import (
	"io/fs"
	"os"
)

const (
	ReadDigit    = 4
	WriteDigit   = 2
	ExecuteDigit = 1

	OwnerShifter  = 6
	GroupShifter  = 3
	OthersShifter = 0
)

func Create(owner int, group int, others int) Permissions {
	return Permissions{createMode(owner, group, others)}
}

func CreateFrom(mode fs.FileMode) Permissions {
	return Permissions{mode}
}

func createMode(owner int, group int, others int) os.FileMode {
	return os.FileMode((owner << OwnerShifter) | (group << GroupShifter) | (others << OthersShifter))
}
