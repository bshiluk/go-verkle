package verkle

import "testing"

func TestParseNodeEmptyPayload(t *testing.T) {
	_, err := ParseNode([]byte{}, 0, []byte{})
	if err != serializedPayloadTooShort {
		t.Fatalf("invalid error, got %v, expected %v", err, "unexpected EOF")
	}
}

func TestLeafStemLength(t *testing.T) {
	// Serialize a leaf with no values, but whose stem is 32 bytes. The
	// serialization should trim the extra byte.
	toolong := [32]byte{}
	leaf := NewLeafNode(toolong[:], make([][]byte, NodeWidth))
	ser, err := leaf.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	if len(ser) != 64 {
		t.Fatalf("invalid serialization when the stem is longer than 31 bytes: %x (%d bytes)", ser, len(ser))
	}
}
