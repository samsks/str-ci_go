package main

import "testing"



func TestSum(t *testing.T){
	total := sumInt(12, 10)

	if total != 22 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 22)
	}
}


