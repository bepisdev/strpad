package strpad

import "fmt"

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
