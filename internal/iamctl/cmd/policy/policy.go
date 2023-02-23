// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package policy provides functions to manage authorization policies on iam platform.
package policy

import (
	"github.com/spf13/cobra"

	cmdutil "github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/iamctl/cmd/util"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/iamctl/util/templates"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/pkg/cli/genericclioptions"
)

var policyLong = templates.LongDesc(`
	Authorization policy management commands.

	This commands allow you to manage your authorization policy on iam platform.`)

// NewCmdPolicy returns new initialized instance of 'policy' sub command.
func NewCmdPolicy(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "policy SUBCOMMAND",
		DisableFlagsInUseLine: true,
		Short:                 "Manage authorization policies on iam platform",
		Long:                  policyLong,
		Run:                   cmdutil.DefaultSubCommandRun(ioStreams.ErrOut),
	}

	cmd.AddCommand(NewCmdCreate(f, ioStreams))
	cmd.AddCommand(NewCmdGet(f, ioStreams))
	cmd.AddCommand(NewCmdList(f, ioStreams))
	cmd.AddCommand(NewCmdDelete(f, ioStreams))
	cmd.AddCommand(NewCmdUpdate(f, ioStreams))

	return cmd
}
