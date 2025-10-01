package fixtures

import (
	"os"
	"path/filepath"
)

var (
	Root = filepath.Join(repoRoot(), "fixtures", "metron")
)

func Path(name string) string {
	return filepath.Join(Root, name)
}

func repoRoot() string {
	wd, _ := os.Getwd()
	return filepath.Join(wd, "../../..")
}
