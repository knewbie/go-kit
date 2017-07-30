package terminal

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"syscall"
	"unsafe"
)

// method 1

func size() (string, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	return string(out), err
}

func parse(input string) (uint, uint, error) {
	parts := strings.Split(input, " ")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", 1))
	if err != nil {
		return 0, 0, err
	}
	return uint(x), uint(y), nil
}

// Width return the width of the terminal.
func Width() (uint, error) {
	output, err := size()
	if err != nil {
		return 0, err
	}
	_, width, err := parse(output)
	return width, err
}

// Height returns the height of the terminal.
func Height() (uint, error) {
	output, err := size()
	if err != nil {
		return 0, err
	}
	height, _, err := parse(output)
	return height, err
}

// method 2

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

var _winsize *winsize

func GetWidth() uint {
	return uint(_winsize.Col)
}

func GetHeight() uint {
	return uint(_winsize.Row)
}

func init() {
	_winsize = &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(_winsize)))

	if int(retCode) == -1 {
		panic(errno)
	}
}
