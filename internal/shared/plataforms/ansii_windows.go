//go:build windows

package plataforms

import "golang.org/x/sys/windows"

// enable Enables ANSII Escape sequence in the current terminal window
func EnableANSII() {
	handle := windows.Handle(windows.Stdout)
	var mode uint32
	windows.GetConsoleMode(handle, &mode)
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	windows.SetConsoleMode(handle, mode)
}
