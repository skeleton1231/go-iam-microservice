// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"strings"

	"github.com/spf13/pflag"
)

func initFlag() {
	pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
}

// WordSepNormalizeFunc changes all flags that contain "_" separators.
func WordSepNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}

	return pflag.NormalizedName(name)
}

/*
This is a small Go code snippet that sets up command-line flag normalization for a CLI application. It replaces underscores with hyphens in command-line flag names.

Here's a brief explanation of the code:

- `initFlag`: This function sets the flag normalization function for the default `pflag` FlagSet to be `WordSepNormalizeFunc`.

- `WordSepNormalizeFunc`: This function is a flag normalization function that replaces underscores with hyphens in flag names.

Flag normalization functions are used to modify flag names before they are used. In this case, if a user provides a flag with an underscore, it is replaced with a hyphen. This can be useful in providing a consistent interface for users. For example, users can use either hyphens or underscores in flag names, and the application treats them the same way.

Note: The `pflag` package is a drop-in replacement for Go's native `flag` package but with POSIX-compliant command-line options behavior.

This code does not run on its own, it is part of a larger application where `initFlag` should be called to set up the flag normalization.
*/
