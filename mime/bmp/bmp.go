package bmp

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

const (
	HeaderSize     = 14
	InfoHeaderSize = 40
	Signature      = "BM" //0x424D
)

type Header struct {
	FileSize   uint32
	DataOffset uint32
}

type BMP struct {
	Header Header
}

func Read(f *os.File) (*BMP, error) {
	fileSignature := make([]byte, 2)
	_, err := f.Read(fileSignature)

	if err != nil {
		return nil, err
	}

	if string(fileSignature) != Signature {
		return nil, errors.New("Invalid Signature")
	}

	fileSize := make([]byte, 4)

	_, err = f.Read(fileSize)

	if err != nil {
		return nil, err
	}

	//Skip reserved useless bytes
	f.Seek(4, io.SeekCurrent)

	dataOffset := make([]byte, 4)

	_, err = f.Read(dataOffset)

	if err != nil {
		return nil, err
	}

	header := Header{
		FileSize:   binary.BigEndian.Uint32(fileSize),
		DataOffset: binary.BigEndian.Uint32(dataOffset),
	}

	return &BMP{
		Header: header,
	}, nil
}
