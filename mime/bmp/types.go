package bmp

type Header struct {
	FileSize   uint32
	DataOffset uint32
}

type InfoHeader struct {
	Size            uint32
	Width           uint32
	Height          uint32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XPixelsPerM     uint32
	YPixelsPerM     uint32
	ColorsUsed      uint32
	ImportantColors uint32
}

type BMP struct {
	Header     Header
	InfoHeader InfoHeader
}
