// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package set used to set specific features on objects.
package set

import (
	"github.com/spf13/cobra"

	cmdutil "github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/iamctl/cmd/util"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/iamctl/util/templates"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/pkg/cli/genericclioptions"
)

var setLong = templates.LongDesc(`
	Configure objects.

	These commands help you make changes to existing objects.`)

// NewCmdSet returns an initialized Command instance for 'set' sub command.
func NewCmdSet(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "set SUBCOMMAND",
		DisableFlagsInUseLine: true,
		Short:                 "Set specific features on objects",
		Long:                  setLong,
		Run:                   cmdutil.DefaultSubCommandRun(ioStreams.ErrOut),
	}

	// add subcommands
	// cmd.AddCommand(NewCmdDB(f, ioStreams))

	return cmd
}
