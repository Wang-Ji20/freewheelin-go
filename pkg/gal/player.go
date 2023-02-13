package gal

import "fmt"

type Player struct {
	TB      *textbase
	Process uint64
}

func NewPlayer(gameName string) (*Player, error) {
	tb, err := OpenTB(gameName, "r")
	if err != nil {
		return nil, err
	}
	return &Player{
		TB:      tb,
		Process: 0,
	}, nil
}

func (p *Player) Quit() {
	p.TB.Close()
}

func (p *Player) Play() {
	for d, err := p.TB.Step(); err == nil; d, err = p.TB.Step() {
		fmt.Println(d)
		fmt.Scanln()
	}
}
