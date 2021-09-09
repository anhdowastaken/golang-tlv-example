package tlv

import (
	"encoding/binary"
	"io"
)

// Writer encodes records into TLV format using a Codec and writes them into a provided io.Writer
type Writer struct {
	writer io.Writer
	codec  *Codec
}

func NewWriter(w io.Writer, codec *Codec) *Writer {
	return &Writer{
		codec:  codec,
		writer: w,
	}
}

// Write encodes records into TLV format using a Codec and writes them into a provided io.Writer
func (w *Writer) Write(rec *Record) error {
	err := writeUint(w.writer, w.codec.TypeBytes, rec.Type)
	if err != nil {
		return err
	}

	ulen := uint(len(rec.Payload))
	err = writeUint(w.writer, w.codec.LenBytes, ulen)
	if err != nil {
		return err
	}

	_, err = w.writer.Write(rec.Payload)
	return err
}

func writeUint(w io.Writer, b ByteSize, i uint) error {
	var num interface{}
	switch b {
	case OneByte:
		num = uint8(i)
	case TwoBytes:
		num = uint16(i)
	case FourBytes:
		num = uint32(i)
	case EightBytes:
		num = uint64(i)
	}
	return binary.Write(w, binary.BigEndian, num)
}
