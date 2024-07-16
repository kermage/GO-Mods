package pathinfo

import (
	"os"
	"path/filepath"

	"github.com/kermage/GO-Mods/pathperms"
)

func Get(path string) PathInfo {
	path, _ = filepath.Abs(path)
	info, err := os.Lstat(path)

	pi := PathInfo{
		fullpath: path,
		stats:    info,
		exists:   true,
	}

	if err == nil {
		pi.mode = info.Mode()
		pi.permissions = pathperms.CreateFrom(pi.mode)
	}

	if os.IsNotExist(err) {
		pi.exists = false
	}

	return pi
}

func (pi *PathInfo) FullPath() string {
	return pi.fullpath
}

func (pi *PathInfo) Permissions() pathperms.Permissions {
	return pi.permissions
}

func (pi *PathInfo) Stats() os.FileInfo {
	return pi.stats
}

func (pi *PathInfo) Mode() os.FileMode {
	return pi.mode
}

func (pi *PathInfo) Exists() bool {
	return pi.exists
}
