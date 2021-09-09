package tlv

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Reader decodes records from TLV format using a Codec from provided io.Reader
type Reader struct {
	codec  *Codec
	reader io.Reader
}

func NewReader(reader io.Reader, codec *Codec) *Reader {
	return &Reader{codec: codec, reader: reader}
}

// Next tries to read a single Record from the io.Reader
func (r *Reader) Next() (*Record, error) {
	// get type
	typeBytes := make([]byte, r.codec.TypeBytes)
	_, err := r.reader.Read(typeBytes)
	if err != nil {
		return nil, err
	}
	typ := readUint(typeBytes, r.codec.TypeBytes)

	// get len
	payloadLenBytes := make([]byte, r.codec.LenBytes)
	_, err = r.reader.Read(payloadLenBytes)
	if err != nil && err != io.EOF {
		return nil, err
	}
	payloadLen := readUint(payloadLenBytes, r.codec.LenBytes)

	if err == io.EOF && payloadLen != 0 {
		return nil, err
	}

	// get value
	v := make([]byte, payloadLen)
	_, err = r.reader.Read(v)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return &Record{
		Type:    typ,
		Payload: v,
	}, nil
}

func readUint(b []byte, sz ByteSize) uint {
	reader := bytes.NewReader(b)
	switch sz {
	case OneByte:
		var i uint8
		binary.Read(reader, binary.BigEndian, &i)
		return uint(i)
	case TwoBytes:
		var i uint16
		binary.Read(reader, binary.BigEndian, &i)
		return uint(i)
	case FourBytes:
		var i uint32
		binary.Read(reader, binary.BigEndian, &i)
		return uint(i)
	case EightBytes:
		var i uint64
		binary.Read(reader, binary.BigEndian, &i)
		return uint(i)
	default:
		return 0
	}
}
