package packets

import "io"

const PACKET_CUSTOM_PAYLOAD int32 = 0x01

type CustomPayloadPacket struct {
	Data []byte
}

func (*CustomPayloadPacket) ID() int32 { return PACKET_CUSTOM_PAYLOAD }

func (c *CustomPayloadPacket) Encode(w *Writer) error { return w.Write(c.Data) }

func (c *CustomPayloadPacket) Decode(r *Reader) error {

	var err error
	c.Data, err = io.ReadAll(r)
	if err == io.ErrUnexpectedEOF {
		return nil
	}

	return err
}
