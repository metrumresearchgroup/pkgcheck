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
	"github.com/dpastoor/goutils"
	"github.com/dpastoor/pkgcheck/rcmdparser"
	"github.com/sirupsen/logrus"
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
	checkDir := args[0]
	ok, _ := goutils.DirExists(fs, checkDir)
	if ok {
		checkResults, err := rcmdparser.NewCheck(fs, checkDir)
		if err != nil {
			return err
		}
		checkResults.Log(log)
		logIfNonZero(log, checkResults.Checks.Notes)
		logIfNonZero(log, checkResults.Checks.Warnings)
		logIfNonZero(log, checkResults.Checks.Errors)
	}
	return nil
}

func init() {
	RootCmd.AddCommand(summaryCmd)
}

func logIfNonZero(lg *logrus.Logger, s []string) error {
	if s == nil || len(s) == 0 {
		return nil
	}
	for _, v := range s {
		lg.Infof("%v", v)
	}
	return nil
}
