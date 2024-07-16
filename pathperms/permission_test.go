package pathperms

import (
	"os"
	"strings"
	"testing"
)

func TestPermission(t *testing.T) {
	for bit := range 8 {
		permission := Permission{os.FileMode(bit)}

		if (permission.Int() == 0 && permission.HasAccess()) || (permission.Int() != 0 && !permission.HasAccess()) {
			t.Errorf("Access did not match: %#o = %s\n", permission, permission)
		}

		testPermission(t, permission)
	}
}

func testPermission(t *testing.T, permission Permission) {
	for _, test := range []struct {
		label      string
		result     bool
		identifier string
	}{
		{"Readable", permission.HasReadAccess(), "r"},
		{"Writable", permission.HasWriteAccess(), "w"},
		{"Executable", permission.HasExecuteAccess(), "x"},
	} {
		t.Logf("%s: %#o = %s\n", test.label, permission, permission)

		if (test.result && !strings.Contains(permission.String(), test.identifier)) || (!test.result && strings.Contains(permission.String(), test.identifier)) {
			t.Errorf("Failed: %#o --> %s is not %s?\n", permission, permission, strings.ToLower(test.label))
		}
	}
}
