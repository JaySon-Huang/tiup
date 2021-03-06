// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"os"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

func main() {
	if err := execute(); err != nil {
		os.Exit(1)
	}
}

func execute() error {
	lang := "en"

	rootCmd := &cobra.Command{
		Use:          "tiup doc",
		Short:        "TiDB document summary page",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if lang == "en" {
				if err := open.Run("https://pingcap.com/docs/stable"); err != nil {
					return err
				}
			} else if lang == "cn" {
				if err := open.Run("https://pingcap.com/docs-cn/stable"); err != nil {
					return err
				}
			} else {
				return errors.New("lang should be en or cn")
			}
			return nil
		},
	}

	rootCmd.Flags().StringVar(&lang, "lang", lang, "The language of the document: en/cn")
	return rootCmd.Execute()
}
