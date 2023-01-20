package gal

import (
	"io"
	"os"
	"strings"
)

type Parser struct {
	TB     *textbase
	Script *os.File
}

func NewParser(sPath string, tbPath string) (*Parser, error) {
	f, err := os.Open(sPath)
	if err != nil {
		return nil, err
	}

	tb, err := OpenTB(tbPath, "rwc")
	if err != nil {
		return nil, err
	}

	return &Parser{
		TB:     tb,
		Script: f,
	}, nil
}

func (p *Parser) CloseParser() {
	p.Script.Close()
	p.TB.Close()
}

func (p *Parser) Parse() error {
	b, err := io.ReadAll(p.Script)
	if err != nil {
		return err
	}

	for _, s := range strings.Split(string(b), "\n\n") {
		if len(s) > 2 && s[:2] == "//" {
			continue
		}
		_, err := p.TB.WriteNormal(s)
		if err != nil {
			return err
		}
	}

	return nil
}
