package tmnlutils

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

// FehCmd invoces the image viewer feh with all desired flags that make the viewing experience nicer; the respective flags will be printed whenever this function is called
func FehCmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("command is", cmd)
	flag1, flag2, flag3, flag4, flag5 := "-xZGdB", "black", "--draw-tinted", "-g 1280x720", os.Args[2]
	fmt.Printf("flags: %s %s %s %s %s\n", flag1, flag2, flag3, flag4, flag5)
	out, err := exec.Command(cmd, flag1, flag2, flag3, flag4, flag5).Output() // 'out' is a slice of bytes
	if err != nil {
		fmt.Printf("error executing command: %s\n", err)
	}
	// print the output (only if feh prints an error value to stdout) to the screen
	for i := 0; i < len(out); i++ {
		fmt.Printf("%s", string(out[i]))
	}
	wg.Done()
}
