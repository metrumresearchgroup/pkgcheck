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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cacheDir     string
	cacheExe     string
	saveExe      string
	cleanLvl     int
	copyLvl      int
	gitignoreLvl int
	git          bool
	oneEst       bool
	runDir       string
)

// runCmd represents the run command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check a package",
	Long: `
	pkc check <pkg tarball> 
	pkc check <pkg dir> --threads=4 //check on 4 threads
	pkc check <pkg tarball> --libDir="<somedir>"
 `,
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {

	//AppFs := afero.NewOsFs()
	// can use this to redirect log output
	// f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()

	// log.SetOutput(f)
	log.Printf("lib dir at: %s", viper.GetString("libDir"))
	log.Printf("running on %v threads", viper.GetInt("threads"))
	return nil
}

func init() {
	RootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
