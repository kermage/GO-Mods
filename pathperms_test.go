package main

import (
	"os"
	"testing"

	"github.com/kermage/GO-Mods/pathperms"
)

func TestEntries(t *testing.T) {
	testFile, err := os.CreateTemp("", "permission_test")

	if err != nil {
		t.Fatalf("Could not create temporary test file: %s", err)
	}

	t.Logf("Temporary test file: %s", testFile.Name())

	defer func() {
		_ = os.Remove(testFile.Name())
	}()

	for owner := range 8 {
		for group := range 8 {
			for others := range 8 {
				access := pathperms.Create(owner, group, others)
				mode := access.Bits

				if testFile.Chmod(mode) != nil {
					t.Errorf("Failed to change the mode of test file to %#o", mode)
				}

				info, err := os.Lstat(testFile.Name())

				if err != nil {
					t.Errorf("Failed to stat test file: %s", err)
				}

				if mode != info.Mode() {
					t.Errorf("Mode did not match from actual file mode: %s != %s\n", mode, info.Mode())
				}

				permissions := pathperms.CreateFrom(info.Mode())

				if permissions != access {
					t.Errorf("Failed to create permissions from actual file mode: %s != %s\n", permissions, access)
				}
			}
		}
	}
}
