package packets

const PACKET_PING_PONG_CONFIGURATION int32 = 0x05

type PingPongConfiguration struct {
	N int32
}

func (*PingPongConfiguration) ID() int32 { return PACKET_PING_PONG_CONFIGURATION }

func (p *PingPongConfiguration) Encode(w *Writer) error { return w.WriteInt(p.N) }

func (p *PingPongConfiguration) Decode(r *Reader) error {
	var err error
	p.N, err = r.ReadInt()
	return err
}
