package gal

import "testing"

func TestParse(t *testing.T) {
	p, err := NewParser("./test.txt", "./parse.tb")
	if err != nil {
		t.Errorf("cannot new parser")
	}
	err = p.Parse()
	if err != nil {
		t.Errorf("cannot parse text file")
	}
}
