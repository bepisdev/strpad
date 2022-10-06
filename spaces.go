package strpad

// Pad with one space to both sides
func Pad(inpstr string) string {
	return fmt.Sprintf(" %s ", inpstr)
}

// Pad with one space to the left
func PadLeft(inpstr string) string {
	return fmt.Sprintf(" %s", inpstr)
}

// Pad with one space to the right
func PadRight(inpstr string) string {
	return fmt.Sprintf("%s ", inpstr)
}

// Pad with N spaces to both sides
func PadN(inpstr string, N int) string {
    _nspaces := strings.Repeat(" ", N)
    return fmt.Sprintf("%s%s%s", _nspaces, inpstr, _nspaces)
}

// Pad with N spaces to the left
func PadNLeft(inpstr string, N int) string {
    _nspaces := strings.Repeat(" ", N)
    return fmt.Sprintf("%s%s", _nspaces, inpstr)
}

// Pad with N spaces to the right
func PadNRight(inpstr string, N int) string {
    _nspaces := strings.Repeat(" ", N)
    return fmt.Sprintf("%s%s", inpstr, _nspaces)
}