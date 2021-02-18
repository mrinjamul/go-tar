package tarwarper

import (
	"os"
	"testing"
)

func TestArchive(t *testing.T) {
	paths := []string{
		"tarwarper.go",
		"tarwarper_test.go",
		"testdata/",
	}

	err := CreateTar(paths, "output_test.tar.gz")
	if err != nil {
		t.Errorf("Failed: %s\n", err)
		return
	}
}

func TestOpen(t *testing.T) {
	if _, err := os.Stat("output_test.tar.gz"); os.IsNotExist(err) {
		t.Error("No file for untaring dected")
		return
	}
	err := os.Mkdir("testing", 0755)
	if err != nil {
		os.RemoveAll("testing")
		os.Mkdir("testing", 0755)
	}

	err = ExtarctTar("testing", "output_test.tar.gz")
	if err != nil {
		t.Errorf("Failed untaring: %s\n", err)
		return
	}

	// files, err := ioutil.ReadDir("testing")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// if len(files) != 3 {
	// 	t.Errorf("The directory wasn't actually written")
	// }
}
