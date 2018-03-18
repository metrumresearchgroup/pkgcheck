package rcmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dpastoor/goutils"
	"github.com/spf13/afero"
)

// Check runs RCmdCheck
// fs filesystem abstraction
// libDir library directory
func Check(
	fs afero.Fs,
	libDir string,
) error {
	ok, err := goutils.DirExists(fs, libDir)
	if !ok || err != nil {
		//TODO: change these exits to instead just return an error probably
		log.Printf("could not find directory to run model %s, ERR: %s, ok: %v", libDir, err, ok)
		return err
	}
	cmdArgs := []string{}
	cmd := exec.Command("R CMD", cmdArgs...)
	// set directory for the shell to relevant directory
	cmd.Dir = modelDir
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
	outputFile, err := os.Create(filepath.Join(modelDir, "stdout.out"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not make stdout file to pipe output to")
	} else {
		defer outputFile.Close()
	}
	errOutputFile, err := os.Create(filepath.Join(modelDir, "stderr.out"))
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
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error attempting to run model, check the lst file in the run directory for more details", err)
		return err
	}
	return nil
}
