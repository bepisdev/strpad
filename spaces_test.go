package strpad

import "testing"

func TestPad(t *testing.T) {
	ans := Pad("test")
	if ans != " test " {
		t.Errorf("Pad(\"test\") = %s; want \" test \"", ans)
	}
}

func TestPadLeft(t *testing.T) {
	ans := PadLeft("test")
	if ans != " test" {
		t.Errorf("PadLeft(\"test\") = %s; want \" test\"", ans)
	}
}

func TestPadRight(t *testing.T) {
	ans := PadRight("test")
	if ans != "test " {
		t.Errorf("PadLeft(\"test\") = %s; want \"test \"", ans)
	}
}

func TestPadN(t *testing.T) {
	ans := PadN("test", 2)
	if ans != "  test  " {
		t.Errorf("Pad(\"test\") = %s; want \"  test  \"", ans)
	}
}

func TestPadNLeft(t *testing.T) {
	ans := PadNLeft("test", 2)
	if ans != "  test" {
		t.Errorf("Pad(\"test\") = %s; want \"  test\"", ans)
	}
}

func TestPadNRight(t *testing.T) {
	ans := PadNRight("test", 2)
	if ans != "test  " {
		t.Errorf("Pad(\"test\") = %s; want \"test  \"", ans)
	}
}
