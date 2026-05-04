package packets

import "github.com/robogg133/MoonMS/internal/datatypes"

const PACKET_DISONNECT_CONFIGURATION int32 = 0x02

type DisconnectPacket struct {
	Reason datatypes.TextComponent
}

func (c *DisconnectPacket) ID() int32 { return PACKET_DISONNECT_CONFIGURATION }

func (c *DisconnectPacket) Encode(w *Writer) error {
	return w.WriteString(string(c.Reason))
}

func (c *DisconnectPacket) Decode(r *Reader) error {
	s, err := r.ReadString()
	c.Reason = datatypes.TextComponent(s)

	return err
}
