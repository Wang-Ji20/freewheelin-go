package gal

import (
	"os"
	"testing"
)

func TestOpenTB(t *testing.T) {
	// case 01: write create
	tb, err := OpenTB("./test.tb", "wc")
	if err != nil {
		t.Errorf("tb open write/create failed.")
	}
	defer tb.Close()
	t.Cleanup(func ()  {
		os.Remove("./test.tb")
	})

	// case 02: read invalid file format
	f, _ := os.OpenFile("./testbad.tb", os.O_CREATE | os.O_RDWR, 00700)
	f.Close()
	_, err = OpenTB("./testbad.tb", "r")
	if err == nil {
		t.Errorf("read invalid format, should have err but nil")
	}
	t.Cleanup(func() {
		os.Remove("./testbad.tb")
	})

}

func TestWrite(t *testing.T) {
	tb, _ := OpenTB("./write.tb", "rwc")
	msg := "阿列克塞·费奥多罗维奇·卡拉马佐夫是我县一位地主费奥多尔·巴甫洛维奇·卡拉马佐夫的第三个儿子"
	tb.WriteNormal(msg)
	tb.Close()

	tb, _ = OpenTB("./write.tb", "r")
	defer tb.Close()
	d, _ := tb.Step()
	if msg != d.Text{
		t.Errorf("expected %s len %d, result %s len %d", msg, len(msg), d.Text, len(d.Text))
	}
	t.Cleanup(
		func() {
			os.Remove("./write.tb")
		})
}