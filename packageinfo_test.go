package main

import (
	"testing"
)

func TestNewPackageInfo(t *testing.T) {
	var pkg *packageInfo

	pkg, _ = newPackageInfo("./...")
	if !pkg.recursive {
		t.Error("./... is resursive")
	}
	if pkg.importPath != "github.com/ToQoz/gopwt" {
		t.Errorf("expected=%s, but got %s", "github.com/ToQoz/gopwt", pkg.importPath)
	}

	pkg, _ = newPackageInfo("./")
	if pkg.recursive {
		t.Error("./ is not resursive")
	}
	if pkg.importPath != "github.com/ToQoz/gopwt" {
		t.Errorf("expected=%s, but got %s", "github.com/ToQoz/gopwt", pkg.importPath)
	}

	pkg, _ = newPackageInfo("github.com/ToQoz/gopwt/...")
	if !pkg.recursive {
		t.Error("github.com/ToQoz/gopwt/... is resursive")
	}
	if pkg.importPath != "github.com/ToQoz/gopwt" {
		t.Errorf("expected=%s, but got %s", "github.com/ToQoz/gopwt", pkg.importPath)
	}
}

func TestPackageInfoGoTestArg(t *testing.T) {
	var pkg *packageInfo
	var got string

	pkg, _ = newPackageInfo("./...")
	got = pkg.ToGoTestArg()
	if got != "github.com/ToQoz/gopwt/..." {
		t.Errorf("expected=%s, but got %s", "github.com/ToQoz/gopwt/...", got)
	}

	pkg, _ = newPackageInfo("./")
	got = pkg.ToGoTestArg()
	if got != "github.com/ToQoz/gopwt" {
		t.Errorf("expected=%s, but got %s", "github.com/ToQoz/gopwt", got)
	}
}