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

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/viper"

	"github.com/metrumresearchgroup/pkgcheck/configlib"
	"github.com/spf13/cobra"
)

// VERSION is the current pkc version
const VERSION string = "0.2.0-alpha.3"

var log *logrus.Logger
var fs afero.Fs

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
	cobra.OnInitialize(initConfig, setGlobals)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/pkgcheck.toml)")
	RootCmd.PersistentFlags().String("loglevel", "", "level for logging")
	RootCmd.PersistentFlags().String("libpaths", "", "library paths, colon separated list")
	RootCmd.PersistentFlags().Int("threads", 0, "number of threads to execute with")
	RootCmd.PersistentFlags().Bool("preview", false, "preview action, but don't actually run command")
	RootCmd.PersistentFlags().Bool("debug", false, "use debug mode")
	RootCmd.PersistentFlags().StringSlice("whitelist", []string{}, "package whitelist")
	RootCmd.PersistentFlags().StringSlice("blacklist", []string{}, "package blacklist")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("loglevel", RootCmd.PersistentFlags().Lookup("loglevel"))
	viper.BindPFlag("libpaths", RootCmd.PersistentFlags().Lookup("libpaths"))
	viper.BindPFlag("threads", RootCmd.PersistentFlags().Lookup("threads"))
	viper.BindPFlag("preview", RootCmd.PersistentFlags().Lookup("preview"))
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("whitelist", RootCmd.PersistentFlags().Lookup("whitelist"))
	viper.BindPFlag("blacklist", RootCmd.PersistentFlags().Lookup("blacklist"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// globals

}

func setGlobals() {

	fs = afero.NewOsFs()

	log = logrus.New()

	switch logLevel := viper.GetString("loglevel"); logLevel {
	case "debug":
		log.Level = logrus.DebugLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "fatal":
		log.Level = logrus.FatalLevel
	case "panic":
		log.Level = logrus.PanicLevel
	default:
		log.Level = logrus.InfoLevel
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" { // enable ability to specify config file via flag
	// 	viper.SetConfigFile(cfgFile)
	// }
	if viper.GetString("config") == "" {
		_ = configlib.LoadGlobalConfig("pkgcheck")
	} else {
		_ = configlib.LoadConfigFromPath(viper.GetString("config"))
	}
	if viper.GetBool("debug") {
		viper.Debug()
	}
}

func expand(s string) string {
	if strings.HasPrefix(s, "~/") {
		return filepath.Join(os.Getenv("HOME"), s[1:])
	}

	return s
}
