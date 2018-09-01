// Copyright © 2018 Devin Pastoor <devin.pastoor@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"path/filepath"

	"github.com/dpastoor/goutils"
	"github.com/metrumresearchgroup/pkgcheck/rcmd"
	"github.com/metrumresearchgroup/pkgcheck/tarutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dest string

// checkCmd represents the R CMD check command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy tarballs to a suitable directory structure for checking",
	Long: `
	copy tarballs to a suitable directory structure for checking
 `,
	RunE: rCopy,
}

func rCopy(cmd *cobra.Command, args []string) error {
	// check if arg is a file or Dir
	// dirty check for if doesn't have an extension is a folder
	var totalCopied int
	fm := rcmd.CreateFilterMap(
		viper.GetStringSlice("whitelist"),
		viper.GetStringSlice("blacklist"),
	)
	for _, arg := range args {
		isDir, err := goutils.IsDir(fs, arg)
		if err != nil || !isDir {
			log.Printf("issue handling %s, if this is a run please add the extension. Err: (%s)", arg, err)
		}
		copied, err := tarutils.CopyPkgTars(fs, arg, dest, fm)
		totalCopied += len(copied)
		for _, c := range copied {
			log.Debugf("copied %s", filepath.Base(c))
		}
		absdest, _ := filepath.Abs(dest)
		log.Infof("copied %v package tarballs to %s", totalCopied, absdest)
	}

	return nil
}

func init() {
	copyCmd.Flags().StringVarP(&dest, "dest", "d", "", "destination for tarballs")
	copyCmd.MarkFlagRequired("dest")
	RootCmd.AddCommand(copyCmd)
}
