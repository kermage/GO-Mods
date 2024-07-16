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
		FullPath: path,
		Stats:    info,
		Exists:   true,
	}

	if err == nil {
		pi.Mode = info.Mode()
		pi.Permissions = pathperms.CreateFrom(pi.Mode)
	}

	if os.IsNotExist(err) {
		pi.Exists = false
	}

	return pi
}
