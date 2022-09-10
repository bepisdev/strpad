package strpad

import "fmt"

// Pad with one space to both sides
func Pad(string inpstr) string {
    return fmt.Sprintf(" %s ", inpstr)
}

// Pad with one space to the left
func PadLeft(string inpstr, N int) string {
    return fmt.Sprintf(" %s", inpstr)
}

// Pad with one space to the right
func PadRight(string inpstr, N int) string {
    return fmt.Sprintf("%s ", inpstr)
}
