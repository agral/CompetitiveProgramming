package lc3484

import "testing"

func Test_Spreadsheet(t *testing.T) {
	s := Constructor(3)
	if actual := s.GetValue("=5+7"); actual != 12 {
		t.Errorf("Operation =5+7 failed: want 12, got %d", actual)
	}

	s.SetCell("A1", 10)
	if actual := s.GetValue("=A1+6"); actual != 16 {
		t.Errorf("Operation =A1+6 failed: want 16, got %d", actual)
	}

	s.SetCell("B2", 15)
	if actual := s.GetValue("=A1+B2"); actual != 25 {
		t.Errorf("Operation =A1+B2 failed: want 25, got %d", actual)
	}

	s.ResetCell("A1")
	if actual := s.GetValue("=A1+B2"); actual != 15 {
		t.Errorf("Operation =A1+B2 after resetting A1 failed: want 15, got %d", actual)
	}

	// Test whether not-yet-initialized values are returned properly as zero:
	if actual := s.GetValue("=C1+12345"); actual != 12345 {
		t.Errorf("Operation =C1+12345 failed: want 12345, got %d", actual)
	}
}
