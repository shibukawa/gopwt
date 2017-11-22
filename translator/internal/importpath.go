package internal

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

func HandleGlobalOrLocalImportPath(globalOrLocalImportPath string) (importpath, fpath string, err error) {
	if globalOrLocalImportPath == "" {
		globalOrLocalImportPath = "."
	}

	if strings.HasPrefix(globalOrLocalImportPath, ".") {
		fpath = filepath.Join(WorkingDir, globalOrLocalImportPath)
		if _, err = os.Stat(fpath); err != nil {
			return
		}

		importpath, err = findImportPathByPath(fpath)
		if err != nil {
			return
		}
	} else {
		importpath = globalOrLocalImportPath

		fpath, err = findPathByImportPath(importpath)
		if err != nil {
			return
		}
	}

	return
}

func findImportPathByPath(path string) (string, error) {
	for _, srcDir := range build.Default.SrcDirs() {
		if path, err := filepath.EvalSymlinks(path); err == nil {
			if srcDir, err := filepath.EvalSymlinks(srcDir); err == nil {
				if strings.HasPrefix(path, srcDir) {
					return strings.TrimPrefix(strings.Replace(path, srcDir, "", 1), "/"), nil
				}
			}
		}
	}

	return "", fmt.Errorf("%s is not found in $GOPATH/src(%q)", path, build.Default.SrcDirs())
}

func findPathByImportPath(importPath string) (string, error) {
	for _, srcDir := range build.Default.SrcDirs() {
		if _, err := os.Stat(filepath.Join(srcDir, importPath)); err == nil {
			return filepath.Join(srcDir, importPath), nil
		}
	}

	return "", fmt.Errorf("package %s is not found in $GOPATH/src(%q)", importPath, build.Default.SrcDirs())
}

func findDeps(importPath, srcDir string) ([]string, error) {
	deps := map[string]bool{}

	pkg, err := build.Import(importPath, srcDir, build.AllowBinary)
	if err != nil {
		return nil, err
	}

	for _, imp := range pkg.Imports {
		if imp == importPath {
			continue
		}
		deps[imp] = true
	}

	for _, imp := range pkg.TestImports {
		if imp == importPath {
			continue
		}
		deps[imp] = true
	}

	ret := make([]string, 0, len(deps))
	for d, _ := range deps {
		ret = append(ret, d)
	}
	return ret, nil
}
