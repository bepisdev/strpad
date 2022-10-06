package strpad

// Pad with X symbol to both sides
func PadSymbol(inpstr string, X string) string {
    return fmt.Sprintf("%s%s%s", X, inpstr, X)
}

// Pad with X symbol to the left
func PadSymbolLeft(inpstr string, X string) string {
    return fmt.Sprintf("%s%s", X, inpstr)
}

// Pad with X symbol to the right
func PadSymbolRight(inpstr string, X string) string {
    return fmt.Sprintf("%s%s", inpstr, X)
}

// Pad with N number of X symbol to both sides
func PadNSymbol(inpstr string, X string, N int) string {
    _nsymbols := strings.Repeat(X, N)
    return fmt.Sprintf("%s%s%s", _nsymbols, inpstr, _nsymbols)
}

// Pad with N number of X symbol to the right
func PadNSymbolRight(inpstr string, X string, N int) string {
    _nsymbols := strings.Repeat(X, N)
    return fmt.Sprintf("%s%s", inpstr, _nsymbols)
}

// Pad with N number of X symbol to the left
func PadNSymbolRight(inpstr string, X string, N int) string {
    _nsymbols := strings.Repeat(X, N)
    return fmt.Sprintf("%s%s", _nsymbols, inpstr)
}