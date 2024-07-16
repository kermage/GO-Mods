package pathinfo

import (
	"os"

	"github.com/kermage/GO-Mods/pathperms"
)

type PathInfo struct {
	fullpath    string
	permissions pathperms.Permissions
	stats       os.FileInfo
	mode        os.FileMode
	exists      bool
}

func (pi *PathInfo) Accessibility() (bool, bool, bool) {
	return pi.permissions.Owner().HasAccess(),
		pi.permissions.Group().HasAccess(),
		pi.permissions.Others().HasAccess()
}

func (pi *PathInfo) Readability() (bool, bool, bool) {
	return pi.permissions.Owner().HasReadAccess(),
		pi.permissions.Group().HasReadAccess(),
		pi.permissions.Others().HasReadAccess()
}

func (pi *PathInfo) Writability() (bool, bool, bool) {
	return pi.permissions.Owner().HasWriteAccess(),
		pi.permissions.Group().HasWriteAccess(),
		pi.permissions.Others().HasWriteAccess()
}

func (pi *PathInfo) Executability() (bool, bool, bool) {
	return pi.permissions.Owner().HasExecuteAccess(),
		pi.permissions.Group().HasExecuteAccess(),
		pi.permissions.Others().HasExecuteAccess()
}
