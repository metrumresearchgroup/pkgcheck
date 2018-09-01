package rcmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dpastoor/goutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// Check runs RCmdCheck
func Check(
	fs afero.Fs,
	cs CheckSettings,
	rs RSettings,
	lg *logrus.Logger,
	preview bool,
) error {
	ok, err := goutils.Exists(fs, cs.TarPath)
	if err != nil {
		lg.Errorf("error in checking for TarPath at: %s", cs.TarPath)
		return err
	}
	if !ok {
		lg.Errorf("no tarball specified at %s", cs.TarPath)
		return fmt.Errorf("no tarball at %s", cs.TarPath)
	}

	if cs.OutputDir != "" {
		ok, err = goutils.DirExists(fs, cs.OutputDir)
		if err != nil {
			lg.Errorf("error in checking for TarPath at: %s", cs.TarPath)
			return err
		}
		if !ok {
			lg.Errorf("no output directory detected at %s", cs.OutputDir)
			return fmt.Errorf("no output directory at %s", cs.TarPath)
		}
	}
	cmdArgs := []string{
		"CMD",
		"check",
	}
	cmdArgs = append(cmdArgs, cs.TarPath)
	cmdFlags := cs.CmdFlags()
	cmdArgs = append(cmdArgs, cmdFlags...)

	envVars := os.Environ()
	ok, rLibsSite := rs.LibPathsEnv()
	if ok {
		envVars = append(envVars, rLibsSite)
	}

	lg.WithFields(
		logrus.Fields{
			"Package":       cs.Package().Name,
			"CheckSettings": cs,
			"RSettings":     rs,
			"env":           rLibsSite,
		}).Debug(cmdArgs)

	// --vanilla is a command for R and should be specified before the CMD, eg
	// R --vanilla CMD check
	if cs.Vanilla {
		cmdArgs = append([]string{"--vanilla"}, cmdArgs...)
	}
	cmd := exec.Command(
		rs.R(),
		cmdArgs...,
	)

	cmd.Env = envVars

	if preview {
		return nil
	}

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	}
	errReader, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdErrPipe for Cmd", err)
	}
	scanner := bufio.NewScanner(cmdReader)
	errScanner := bufio.NewScanner(errReader)
	outputFileName := filepath.Join(cs.OutputDir, fmt.Sprintf("%s_stdout.out", cs.Package().Name))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not make stdout file to pipe output to %s \n", outputFileName)
	} else {
		defer outputFile.Close()
	}
	errOutputFile, err := os.Create(filepath.Join(cs.OutputDir, fmt.Sprintf("%s_stderr.out", cs.Package().Name)))
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not make stderr file to pipe output to")
	} else {
		defer errOutputFile.Close()
	}
	outputFileWriter := bufio.NewWriter(outputFile)
	errOutputFileWriter := bufio.NewWriter(outputFile)
	// handles where to write output to
	go func() {
		for scanner.Scan() {
			//fmt.Fprintf(outputFileWriter, "%s out | %s\n", runName, scanner.Text())
			// the %s out was generally unneeded unless pushing multiple streams
			// out to the cli at once, but that is currently suppressed until
			// a multiwriter is better supported, will just output the stdout itself to
			// the output writter
			fmt.Fprintf(outputFileWriter, "%s\n", scanner.Text())
		}
	}()
	go func() {
		for errScanner.Scan() {
			fmt.Fprintf(errOutputFileWriter, "%s\n", errScanner.Text())
		}
	}()
	// I think defering these here should be reasonable to make sure they flush before
	// returning any errors from the start/wait processes. Originally had these after the
	// potential errors and think I was missing capturing the flushed lines because
	// the main thread was closing too quickly
	defer outputFileWriter.Flush()
	defer errOutputFileWriter.Flush()
	err = cmd.Start()
	if err != nil {
		lg.Errorf("Error starting Cmd: %s", err)
		return err
	}
	err = cmd.Wait()
	if err != nil {
		lg.Errorf("Cmd failed with error: %s", err)
		return err
	}
	return nil
}
