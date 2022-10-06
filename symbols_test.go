package strpad

import "testing"

func TestPadSymbol(t *testing.T) {
	ans := PadSymbol("test", "X")
	if ans != "XtestX" {
		t.Errorf("PadSymbol(\"test\", \"X\") = %s; want \"XtestX\"", ans)
	}
}

func TestPadSymbolLeft(t *testing.T) {
	ans := PadSymbolLeft("test", "X")
	if ans != "Xtest" {
		t.Errorf("PadSymbolLeft(\"test\", \"X\") = %s; want \"Xtest\"", ans)
	}
}

func TestPadSymbolRight(t *testing.T) {
	ans := PadSymbolRight("test", "X")
	if ans != "testX" {
		t.Errorf("PadSymbolRight(\"test\", \"X\") = %s; want \"testX\"", ans)
	}
}

func TestPadNSymbol(t *testing.T) {
	ans := PadNSymbol("test", "X", 2)
	if ans != "XXtestXX" {
		t.Errorf("PadNSymbol(\"test\", \"X\", 2) = %s; want \"XXtestXX\"", ans)
	}
}

func TestPadNSymbolLeft(t *testing.T) {
	ans := PadNSymbolLeft("test", "X", 2)
	if ans != "XXtest" {
		t.Errorf("PadNSymbolLeft(\"test\", \"X\", 2) = %s; want \"XXtest\"", ans)
	}
}

func TestPadNSymbolRight(t *testing.T) {
	ans := PadNSymbolRight("test", "X", 2)
	if ans != "testXX" {
		t.Errorf("PadNSymbolRight(\"test\", \"X\", 2) = %s; want \"testXX\"", ans)
	}
}
