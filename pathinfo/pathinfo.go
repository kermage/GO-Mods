package pathinfo

import (
	"os"

	"github.com/kermage/GO-Mods/pathperms"
)

type PathInfo struct {
	FullPath    string
	Permissions pathperms.Permissions
	Stats       os.FileInfo
	Mode        os.FileMode
	Exists      bool
}

func (pi *PathInfo) Accessibility() (bool, bool, bool) {
	return pi.Permissions.Owner().HasAccess(),
		pi.Permissions.Group().HasAccess(),
		pi.Permissions.Others().HasAccess()
}

func (pi *PathInfo) Readability() (bool, bool, bool) {
	return pi.Permissions.Owner().HasReadAccess(),
		pi.Permissions.Group().HasReadAccess(),
		pi.Permissions.Others().HasReadAccess()
}

func (pi *PathInfo) Writability() (bool, bool, bool) {
	return pi.Permissions.Owner().HasWriteAccess(),
		pi.Permissions.Group().HasWriteAccess(),
		pi.Permissions.Others().HasWriteAccess()
}

func (pi *PathInfo) Executability() (bool, bool, bool) {
	return pi.Permissions.Owner().HasExecuteAccess(),
		pi.Permissions.Group().HasExecuteAccess(),
		pi.Permissions.Others().HasExecuteAccess()
}
