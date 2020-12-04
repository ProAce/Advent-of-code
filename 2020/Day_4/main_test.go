package main

import "testing"

func TestCheckHcl(t *testing.T) {
	tests := map[string]bool{"#123abc": true, "#123abz": false, "123abc": false}

	for input, outcome := range tests {
		if checkHcl(input) != outcome {
			t.Errorf("For input: %s; Got %t, want: %t", input, !outcome, outcome)
		}
	}
}

func TestCheckPid(t *testing.T) {
	tests := map[string]bool{"000000001": true, "0123456789": false}

	for input, outcome := range tests {
		if checkPid(input) != outcome {
			t.Errorf("For input: %s; Got %t, want: %t", input, !outcome, outcome)
		}
	}
}

func TestCheckEcl(t *testing.T) {
	tests := map[string]bool{"brn": true, "wat": false, "brnb": false}

	for input, outcome := range tests {
		if checkEcl(input) != outcome {
			t.Errorf("For input: %s; Got %t, want: %t", input, !outcome, outcome)
		}
	}
}

func TestCheckHgt(t *testing.T) {
	tests := map[string]bool{"60in": true, "190cm": true, "190in": false, "190": false}

	for input, outcome := range tests {
		if checkHgt(input) != outcome {
			t.Errorf("For input: %s; Got %t, want: %t", input, !outcome, outcome)
		}
	}
}

func TestCheckByr(t *testing.T) {
	tests := map[string]bool{"2002": true, "2003": false, "1920": true, "1919": false}

	for input, outcome := range tests {
		if checkByr(input) != outcome {
			t.Errorf("For input: %s; Got %t, want: %t", input, !outcome, outcome)
		}
	}
}
