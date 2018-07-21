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
	"path"
	"strings"

	"github.com/dpastoor/goutils"
	"github.com/r-infra/pkgcheck/rcmdparser"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// checkCmd represents the R CMD check command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "summarize a check directory",
	Long: `
	summarize the output of a Rcheck directory
 `,
	RunE: rSummary,
}

func rSummary(cmd *cobra.Command, args []string) error {
	var checkDirs []string
	for _, arg := range args {
		// check if arg is a file or Dir
		// dirty check for if doesn't have an extension is a folder
		isDir, err := goutils.IsDir(fs, arg)
		if err != nil || !isDir {
			log.Printf("issue handling %s, if this is a run please add the extension. Err: (%s)", arg, err)
			continue
		}
		if isDir {
			// trim trailing slashes for that could arise on unix/windows
			arg = strings.TrimSuffix(arg, "/")
			arg = strings.TrimSuffix(arg, "\\\\")
			// is it an Rcheckdir or declaring a dir of Rcheck outputs
			if strings.HasSuffix(arg, ".Rcheck") {
				checkDirs = append(checkDirs, arg)
			} else {
				dirInfo, _ := afero.ReadDir(fs, arg)
				subcheckDirs := goutils.ListFilesByExt(goutils.ListDirNames(dirInfo), ".Rcheck")
				for _, d := range subcheckDirs {
					checkDirs = append(checkDirs, path.Join(arg, d))
				}
			}
		}
	}

	if len(checkDirs) == 0 {
		log.Fatalf("no log directories found given args: %s", args)
	} else {
		log.Infof("%v Rcheck directories found, summaries: ", len(checkDirs))
	}

	for _, checkDir := range checkDirs {
		ok, _ := goutils.DirExists(fs, checkDir)
		if ok {
			checkResults, err := rcmdparser.NewCheck(fs, checkDir)
			if err != nil {
				log.Errorf("error checking dir %s, err: %s", checkDir, err)
			}
			checkResults.Log(log)
			logIfNonZero(log, checkResults.Checks.Notes, "info")
			logIfNonZero(log, checkResults.Checks.Warnings, "warn")
			logIfNonZero(log, checkResults.Checks.Errors, "error")
		}
	}
	return nil
}

func init() {
	RootCmd.AddCommand(summaryCmd)
}

func logIfNonZero(lg *logrus.Logger, s []string, lvl string) error {
	lf := lg.Errorf
	switch logLevel := lvl; logLevel {
	case "debug":
		lf = lg.Debugf
	case "info":
		lf = lg.Infof
	case "warn":
		lf = lg.Warnf
	case "error":
		lf = lg.Errorf
	case "fatal":
		lf = lg.Fatalf
	case "panic":
		lf = lg.Panicf
	default:
	}
	if s == nil || len(s) == 0 {
		return nil
	}
	for _, v := range s {
		lf("%v", v)
	}
	return nil
}
