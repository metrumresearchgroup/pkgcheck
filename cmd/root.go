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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"

	"github.com/dpastoor/pkgcheck/configlib"
	"github.com/spf13/cobra"
)

// VERSION is the current pkc version
const VERSION string = "0.0.1"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pkc",
	Short: "package check functionality",
	Long:  fmt.Sprintf("pkc cli version %s", VERSION),
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/pkgcheck.toml)")
	RootCmd.PersistentFlags().String("libPaths", "", "library paths, colon separated list")
	RootCmd.PersistentFlags().Int("threads", 0, "number of threads to execute with")
	RootCmd.PersistentFlags().Bool("preview", false, "preview action, but don't actually run command")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("libPaths", RootCmd.PersistentFlags().Lookup("libPaths"))
	viper.BindPFlag("threads", RootCmd.PersistentFlags().Lookup("threads"))
	viper.BindPFlag("preview", RootCmd.PersistentFlags().Lookup("preview"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" { // enable ability to specify config file via flag
	// 	viper.SetConfigFile(cfgFile)
	// }
	// TODO: set config a little more flexibly
	if viper.GetString("config") == "" {
		_ = configlib.LoadGlobalConfig("pkgcheck")
	} else {
		_ = configlib.LoadConfigFromPath(viper.GetString("config"))
	}
	viper.Debug()
}

func expand(s string) string {
	if strings.HasPrefix(s, "~/") {
		return filepath.Join(os.Getenv("HOME"), s[1:])
	}

	return s
}
