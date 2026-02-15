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

func readHeader(f *os.File) (*Header, error) {
	var err error
	fileSignature := make([]byte, 2)

	if _, err := f.Read(fileSignature); err != nil {
		return nil, err
	}

	if string(fileSignature) != Signature {
		return nil, errors.New("Invalid Signature")
	}

	fileSize := make([]byte, 4)

	if _, err = f.Read(fileSize); err != nil {
		return nil, err
	}

	//Skip reserved useless bytes
	f.Seek(4, io.SeekCurrent)

	dataOffset := make([]byte, 4)

	if _, err = f.Read(dataOffset); err != nil {
		return nil, err
	}

	return &Header{
		FileSize:   binary.LittleEndian.Uint32(fileSize),
		DataOffset: binary.LittleEndian.Uint32(dataOffset),
	}, nil
}

func readInfoHeader(f *os.File) (*InfoHeader, error) {
	//Skip size (is static data)
	f.Seek(4, io.SeekCurrent)

	width := make([]byte, 4)

	if _, err := f.Read(width); err != nil {
		return nil, err
	}

	height := make([]byte, 4)

	if _, err := f.Read(height); err != nil {
		return nil, err
	}

	planes := make([]byte, 2)

	if _, err := f.Read(planes); err != nil {
		return nil, err
	}

	bitsPerPixel := make([]byte, 2)

	if _, err := f.Read(bitsPerPixel); err != nil {
		return nil, err
	}

	compression := make([]byte, 4)

	if _, err := f.Read(compression); err != nil {
		return nil, err
	}

	imageSize := make([]byte, 4)

	if _, err := f.Read(imageSize); err != nil {
		return nil, err
	}

	xPixelsPerM := make([]byte, 4)

	if _, err := f.Read(xPixelsPerM); err != nil {
		return nil, err
	}

	yPixelsPerM := make([]byte, 4)

	if _, err := f.Read(yPixelsPerM); err != nil {
		return nil, err
	}

	colorsUsed := make([]byte, 4)

	if _, err := f.Read(colorsUsed); err != nil {
		return nil, err
	}

	importantcOLORS := make([]byte, 4)

	if _, err := f.Read(importantcOLORS); err != nil {
		return nil, err
	}

	return &InfoHeader{
		Width:           binary.LittleEndian.Uint32(width),
		Height:          binary.LittleEndian.Uint32(height),
		Planes:          binary.LittleEndian.Uint16(planes),
		BitsPerPixel:    binary.LittleEndian.Uint16(bitsPerPixel),
		Compression:     binary.LittleEndian.Uint32(compression),
		ImageSize:       binary.LittleEndian.Uint32(imageSize),
		XPixelsPerM:     binary.LittleEndian.Uint32(xPixelsPerM),
		YPixelsPerM:     binary.LittleEndian.Uint32(yPixelsPerM),
		ColorsUsed:      binary.LittleEndian.Uint32(colorsUsed),
		ImportantColors: binary.LittleEndian.Uint32(importantcOLORS),
	}, nil
}

func Read(f *os.File) (*BMP, error) {
	header, err := readHeader(f)

	if err != nil {
		return nil, err
	}

	infoHeader, err := readInfoHeader(f)

	if err != nil {
		return nil, err
	}

	return &BMP{
		Header:     *header,
		InfoHeader: *infoHeader,
	}, nil
}
