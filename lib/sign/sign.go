package sign

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

const (
	exit = "q"
	y    = "Y"
	n    = "N"
)

// InitSignal register signals handler.
func InitSignal(version string) {
	if goos := runtime.GOOS; goos == "windows" { //兼容windows系统
		fmt.Println("enter q to exit...")

		var input string
		for {
			fmt.Scanln(&input)
			if input == exit {
				fmt.Println("did you watch exit this session(Y/N)?")
				fmt.Scanln(&input)
				if input == y {
					return
				}
				continue
			}
		}
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	for {
		s := <-c
		fmt.Println("router[%s] get a signal %s", version, s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			return
		default:
			return
		}
	}
}
