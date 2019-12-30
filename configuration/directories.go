// This file is part of arduino-cli.
//
// Copyright 2019 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package configuration

import (
	"github.com/arduino/go-paths-helper"
	"github.com/spf13/viper"
)

// HardwareDirectories returns all paths that may contains hardware packages.
func HardwareDirectories() paths.PathList {
	res := paths.PathList{}

	if IsBundledInDesktopIDE() {
		ideDir := paths.New(viper.GetString("IDE.Directory"))
		bundledHardwareDir := ideDir.Join("hardware")
		if bundledHardwareDir.IsDir() {
			res.Add(bundledHardwareDir)
		}
	}

	if viper.IsSet("directories.Data") {
		packagesDir := PackagesDir()
		if packagesDir.IsDir() {
			res.Add(packagesDir)
		}
	}

	if viper.IsSet("directories.User") {
		skDir := paths.New(viper.GetString("directories.User"))
		hwDir := skDir.Join("hardware")
		if hwDir.IsDir() {
			res.Add(hwDir)
		}
	}

	return res
}

// BundleToolsDirectories returns all paths that may contains bundled-tools.
func BundleToolsDirectories() paths.PathList {
	res := paths.PathList{}

	if IsBundledInDesktopIDE() {
		ideDir := paths.New(viper.GetString("IDE.Directory"))
		bundledToolsDir := ideDir.Join("hardware", "tools")
		if bundledToolsDir.IsDir() {
			res = append(res, bundledToolsDir)
		}
	}

	return res
}

// IDEBundledLibrariesDir returns the libraries directory bundled in
// the Arduino IDE. If there is no Arduino IDE or the directory doesn't
// exists then nil is returned
func IDEBundledLibrariesDir() *paths.Path {
	if IsBundledInDesktopIDE() {
		ideDir := paths.New(viper.GetString("IDE.Directory"))
		libDir := ideDir.Join("libraries")
		if libDir.IsDir() {
			return libDir
		}
	}

	return nil
}

// LibrariesDir returns the full path to the user directory containing
// custom libraries
func LibrariesDir() *paths.Path {
	return paths.New(viper.GetString("directories.User")).Join("libraries")
}

// PackagesDir returns the full path to the packages folder
func PackagesDir() *paths.Path {
	return paths.New(viper.GetString("directories.Data")).Join("packages")
}
