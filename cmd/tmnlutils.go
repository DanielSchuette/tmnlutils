// TmnlUtils provides a suite of useful terminal commands that can be called from anywhere
// as long as your bin directory is in your $PATH and that can be customized to meet your needs
// e.g. easy opening of applications or of cheatsheets that I provide with the the GitHub
// repository in which this code lives. All utility functions are located in ../utils/ and
// are separate *.go files for easier access and maintenance, especially if you want to change certain flags.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/DanielSchuette/tmnlutils"
)

func main() {
	tmnlutils.QuteCmd()
	defer fmt.Println("done\n----------------")
	wg := new(sync.WaitGroup)
	fmt.Println("----------------")
	if len(os.Args) == 1 {
		fmt.Println("Please provide one of the following options and make sure that you have the respective binaries installed:\nfeh\n...TODO!")
		os.Exit(1)
	}
	// check if user provided the one, required command line argument to binary
	if os.Args[1] == "feh" {
		if len(os.Args) != 3 {
			fmt.Println("Please provide a valid file path to an image that you want to open with feh.\nSee '$man feh' for more information on which files feh can actually read and display.")
			os.Exit(1)
		}
		// call 'exeCmd' with a bash command
		// TODO: if user specifies a -script flag, parse a bash script from a given path and execute it
		wg.Add(1)
		var command string = "feh"
		go tmnlutils.FehCmd(command, wg)
		wg.Wait()
	} else if os.Args[1] == "ranger" || os.Args[1] == "calcurse" || os.Args[1] == "cmus" {
		fmt.Println("opening", os.Args[1], "...")
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			out, err := exec.Command("sh", "/Users/daniel/Desktop/ranger.sh").Output()
			if err != nil {
				fmt.Printf("error executing command: %s\n", err)
			}
			for i := 0; i < len(out); i++ {
				fmt.Printf("%s", string(out[i]))
			}
			wg.Done()
		}(wg)
		wg.Wait()
	} else if os.Args[1] == "net" {
		_, err := exec.Command("open", "-a", "qutebrowser").Output()
		if err != nil {
			fmt.Printf("error executing command: %s\n", err)
		}
	} else {
		fmt.Println(os.Args[1], "is not (yet) supported - maybe it'll never be!")
		os.Exit(1)
	}
}
