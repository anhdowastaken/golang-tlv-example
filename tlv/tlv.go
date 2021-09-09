package tlv // ByteSize is the size of a field in bytes. Used to define the size of the type and length field in a message.

type ByteSize int

const (
	OneByte    ByteSize = 1
	TwoBytes   ByteSize = 2
	FourBytes  ByteSize = 4
	EightBytes ByteSize = 8
)

// Record represents a record of data encoded in the TLV message.
type Record struct {
	Payload []byte
	Type    uint
}

// Codec is the configuration for a specific TLV encoding/decoding tasks.
type Codec struct {
	// TypeBytes defines the size in bytes of the message type field.
	TypeBytes ByteSize

	// LenBytes defines the size in bytes of the message length field.
	LenBytes ByteSize
}
