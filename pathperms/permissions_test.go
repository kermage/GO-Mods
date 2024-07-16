package pathperms

import (
	"testing"
)

func TestPermissions(t *testing.T) {
	for owner := range 8 {
		for group := range 8 {
			for others := range 8 {
				mode := createMode(owner, group, others)
				permissions := Permissions{mode}

				testPermissions(t, permissions)
			}
		}
	}
}

func testPermissions(t *testing.T, permissions Permissions) {
	for _, class := range []string{"Owner", "Group", "Others"} {
		var permission Permission

		switch class {
		case "Owner":
			permission = permissions.Owner()
		case "Group":
			permission = permissions.Group()
		case "Others":
			permission = permissions.Others()
		default:
			t.Errorf("Invalid access class: %s", class)
		}

		t.Logf("%s? %#o = %s --> %#o = %s\n", class, permissions, permissions, permission, permission)
		testPermission(t, permission)
	}
}
