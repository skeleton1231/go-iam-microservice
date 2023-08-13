// Copyright 2020 Talhuang<talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
)

// CliOptions abstracts configuration options for reading parameters from the
// command line.
type CliOptions interface {
	// AddFlags adds flags to the specified FlagSet object.
	// AddFlags(fs *pflag.FlagSet)
	Flags() (fss cliflag.NamedFlagSets)
	Validate() []error
}

// ConfigurableOptions abstracts configuration options for reading parameters
// from a configuration file.
type ConfigurableOptions interface {
	// ApplyFlags parsing parameters from the command line or configuration file
	// to the options instance.
	ApplyFlags() []error
}

// CompleteableOptions abstracts options which can be completed.
type CompleteableOptions interface {
	Complete() error
}

// PrintableOptions abstracts options which can be printed.
type PrintableOptions interface {
	String() string
}

/*
This Go code defines several interfaces that abstract out common functionality for handling command-line options in a CLI application.

- `CliOptions` Interface: This interface is designed for reading parameters from the command line. It provides two methods:
    - `Flags()`: This returns an instance of `cliflag.NamedFlagSets`, which presumably is a collection of named sets of command-line flags. The specifics would depend on the implementation of the `cliflag.NamedFlagSets` type.
    - `Validate()`: This returns a slice of errors, potentially indicating any validation issues with the options provided in the command line.

- `ConfigurableOptions` Interface: This interface is meant for options that can be set via a configuration file. It provides one method:
    - `ApplyFlags()`: This method returns a slice of errors, indicating potential issues in parsing parameters from the command line or a configuration file and applying them to the options instance.

- `CompleteableOptions` Interface: This interface is meant for options that can be "completed". The specifics of what "completion" means would depend on the implementing type. It provides one method:
    - `Complete()`: This method returns an error, potentially indicating issues with completing the options.

- `PrintableOptions` Interface: This interface is meant for options that can be converted to a string for display purposes. It provides one method:
    - `String()`: This returns a string representation of the option.

These interfaces provide a flexible way to handle and process options in a CLI application, supporting validation, application of options from various sources (command line, configuration files), "completion" of options, and conversion to string for display.
*/
