package gal

import (
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"strings"
)

const magicNumber uint8 = 0xAB

type textbase struct {
	Path string
	Perm int
	File *os.File
	Next uint64
}

type dialogue struct {
	Length uint64
	Text string
	Next uint64
}

func newTB(path string, perm int) (*textbase, error) {
	f, err := os.OpenFile(path, perm, 00700)

	if err != nil {
		return nil, fmt.Errorf("newTB failed: %w", err)
	}

	if perm&os.O_CREATE != 0 {
		b := []byte{magicNumber}
		n, err := f.Write(b)
		if err != nil || n != 1 {
			f.Close()
			return nil, err
		}
	} else {
		b := make([]byte, 1)
		n, err := f.Read(b)
		if err != nil || n != 1{
			f.Close()
			return nil, fmt.Errorf("the file is corrupted or not gal format: %w", err)
		}
	}

	return &textbase{
		Path: path,
		File: f,
	}, nil
}


func OpenTB(path string, mode string) (*textbase, error) {
	var perm int
	for _, m := range mode {
		switch m {
		case 'r':
			perm |= os.O_RDONLY
		case 'w':
			perm |= os.O_WRONLY
		case 'c':
			perm |= os.O_CREATE
		}
	}
	
	t, err := newTB(path, perm)
	if err != nil {
		return nil, fmt.Errorf("OpenTB failed: %w", err)
	}
	
	return t, nil
}


func (t *textbase) Close() {
	defer t.File.Close()
}


func (t *textbase) write(b []byte, offset uint64) (int, error) {
	l := uint64(len(b))
	buf := make([]byte,8)
	
	binary.BigEndian.PutUint64(buf, l)
	n1, err := t.File.Write(buf)
	if err != nil {
		return 0, err
	}

	n2, err := t.File.Write(b)
	if err != nil {
		return 0, err
	}

	binary.BigEndian.PutUint64(buf, offset)
	n3, err := t.File.Write(buf)
	if err != nil {
		return 0, err
	}

	return (n1+n2+n3), nil
}

func (t *textbase) WriteNormal(s string) (int, error) {
	b, err := syscall.ByteSliceFromString(s)
	if err != nil {
		return 0, err
	}
	return t.write(b, 0)
}

func (t *textbase) readnum() (uint64, error) {
	buf := make([]byte, 8)
	n, err := t.File.Read(buf)
	if err != nil || n != 8{
		return 0, err
	}
	return binary.BigEndian.Uint64(buf), nil
}

func (t *textbase) Step() (d dialogue, err error) {
	l, err := t.readnum()
	if err != nil {
		return d, err
	}
	d.Length = l
	buf := make([]byte, l)
	t.File.Read(buf)
	d.Text = strings.TrimRight(string(buf), "\x00\r\n")
	t.Next, _ = t.readnum()
	d.Next = 0
	return d, nil
}