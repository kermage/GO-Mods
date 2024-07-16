package main

import (
	"testing"

	"github.com/kermage/GO-Mods/pathinfo"
)

// var testFiles = []string{"main.go", "test.go"}
var testFiles = map[string]bool{
	"main.go": true,
	"test.go": false,
}

func TestFullness(t *testing.T) {
	for filename := range testFiles {
		testFile := pathinfo.Get(filename)

		t.Logf("Get %q full path", filename)

		if testFile.FullPath == filename {
			t.Errorf("Failed resolving full path: %s == %s", testFile.FullPath, filename)
		}
	}
}

func TestExistence(t *testing.T) {
	for filename, exists := range testFiles {
		testFile := pathinfo.Get(filename)

		t.Logf("Check if %q exists", filename)

		if (exists && !testFile.Exists) || (!exists && testFile.Exists) {
			t.Errorf("%s? %s = %s --> %s", "Exist", testFile.FullPath, testFile.Mode.String(), testFile.Stats)
		}
	}
}
