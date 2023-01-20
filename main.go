package main

import "main/gal"

func main() {
	p, err := gal.NewParser("./test.txt", "./test.tb")
	if err != nil {
		return
	}
	defer p.CloseParser()
	p.Parse()

	pl, err := gal.NewPlayer("./test.tb")
	if err != nil {
		return
	}
	defer pl.Quit()

	pl.Play()

}
