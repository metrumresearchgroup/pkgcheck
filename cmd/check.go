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
	"strings"
	"sync"
	"time"

	"github.com/dpastoor/goutils"

	"github.com/spf13/afero"
	"github.com/spf13/viper"

	"github.com/dpastoor/pkgcheck/rcmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// checkCmd represents the R CMD check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check a package",
	Long: `
	pkc check <pkg tarball> 
	pkc check . // check all tarballs in current directory
	pkc check <pkg dir> --threads=4 //check on 4 threads
	pkc check <pkg tarball> --libpaths=<somedir(s)>
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

	// Log as JSON instead of the default ASCII formatter.
	log := logrus.New()
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example

	// Only log the warning severity or above.
	log.Level = logrus.DebugLevel

	fs := afero.NewOsFs()
	csTemplate := rcmd.CheckSettings{
		OutputDir: viper.GetString("output"),
		Vanilla:   true,
		Cran:      false,
	}

	libPaths := strings.Split(viper.GetString("libpaths"), ":")
	ok, err := validateLibPaths(fs, libPaths)
	if err != nil || !ok {
		log.Fatalf("error checking libPaths: %s", err)
	}

	rs := rcmd.RSettings{
		LibPaths: libPaths,
		Rpath:    viper.GetString("rpath"),
	}
	var wg sync.WaitGroup
	queue := make(chan struct{}, viper.GetInt("threads"))
	start := time.Now()

	log.Debugf("setting up a work queue with %v workers", viper.GetInt("threads"))

	defer close(queue)

	for _, arg := range args {

		// check if arg is a file or Dir
		// dirty check for if doesn't have an extension is a folder
		_, ext := goutils.FileAndExt(arg)
		if ext == "" || arg == "." {
			// could be directory, will need to be careful about the waitgroup as don't want to
			// keep waiting forever since it
			isDir, err := goutils.IsDir(fs, arg)
			if err != nil || !isDir {
				log.Printf("issue handling %s, if this is a run please add the extension. Err: (%s)", arg, err)
				continue
			}
			dirInfo, _ := afero.ReadDir(fs, arg)

			tarsInDir := goutils.ListFilesByExt(goutils.ListFiles(dirInfo), "tar.gz")
			// TODO: add black/white list of tars to check against

			if err != nil {
				log.Printf("issue getting models in dir %s, if this is a run please add the extension. Err: (%s)", arg, err)
				continue
			}
			log.Debugf("adding %v model files in directory %s to queue", len(tarsInDir), arg)

			wg.Add(len(tarsInDir))
			for _, tar := range tarsInDir {
				log.Infof("adding tarball %s to queue\n", tar)
				cs := csTemplate
				cs.TarPath = tar
				go runCheck(fs, queue, &wg, cs, rs, log, viper.GetBool("preview"))
			}

		} else {
			// figure out if need to do expansion, or run as-is
			log.Infof("adding tarball %s to queue\n", arg)
			wg.Add(1)
			cs := csTemplate
			cs.TarPath = arg
			go runCheck(fs, queue, &wg, cs, rs, log, viper.GetBool("preview"))
		}
	}

	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("running on %v threads took %s", viper.GetInt("threads"), elapsed)
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

// validateLibPaths validates all libPaths exist
func validateLibPaths(fs afero.Fs, libPaths []string) (bool, error) {
	if len(libPaths) == 1 {
		// could be either "" or a single libPath
		if libPaths[0] != "" {
			ok, err := goutils.DirExists(fs, libPaths[0])
			if !ok || err != nil {
				return false, fmt.Errorf("libPath does not exist at: %s, err: %s", libPaths[0], err)
			}
		}
	} else {
		for _, lp := range libPaths {
			ok, err := goutils.DirExists(fs, lp)
			if !ok || err != nil {
				return false, fmt.Errorf("libPath does not exist at: %s, err: %s", libPaths[0], err)
			}
		}
	}

	return true, nil
}

func runCheck(
	fs afero.Fs,
	queue chan struct{},
	wg *sync.WaitGroup,
	cs rcmd.CheckSettings,
	rs rcmd.RSettings,
	log *logrus.Logger,
	preview bool,
) error {
	queue <- struct{}{}
	log.Infof("running check for package: %s", cs.Package().Name)
	defer func() {
		wg.Done()
		<-queue
	}()
	err := rcmd.Check(fs, cs, rs, log, preview)

	if err != nil {
		log.Warnf("failed check for package: %s", cs.Package().Name)
		return err
	}
	log.Infof("completed check for package: %s", cs.Package().Name)
	return nil
}
