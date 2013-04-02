package termsize

import (
	"syscall"
	"unsafe"
)

// stolen from go-nuts post by "Ostsol":
type Size struct {
	Row, Col, Xpixel, Ypixel uint16
}

func Get() Size {
	ws := Size{}
	syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(0), uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))
	return ws
}
