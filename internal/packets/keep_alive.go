package packets

const PACKET_KEEP_ALIVE int32 = 0x04

type KeepAlivePacket struct {
	KeepAliveID int64
}

func (*KeepAlivePacket) ID() int32 { return PACKET_KEEP_ALIVE }

func (p *KeepAlivePacket) Encode(w *Writer) error { return w.WriteLong(p.KeepAliveID) }

func (p *KeepAlivePacket) Decode(r *Reader) error {
	var err error
	p.KeepAliveID, err = r.ReadLong()
	return err
}
