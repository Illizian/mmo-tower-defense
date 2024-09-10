package terminal

const (
	// Text Colors
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Text Attributes
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Reversed  = "\033[7m"

	// Background Colors
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	// Cursor and Screen Control
	ClearScreen    = "\033[2J"
	ClearLine      = "\033[K"
	ResetCursor    = "\033[H"
	CursorHide     = "\033[?25l"
	CursorShow     = "\033[?25h"
	CursorUp       = "\033[A"
	CursorDown     = "\033[B"
	CursorForward  = "\033[C"
	CursorBackward = "\033[D"
)
