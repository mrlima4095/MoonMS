package packets

const PACKET_RESET_CHAT int32 = 0x06

type ResetChatPacket struct{}

func (*ResetChatPacket) ID() int32 { return PACKET_RESET_CHAT }

func (p *ResetChatPacket) Encode(w *Writer) error { return nil }

func (p *ResetChatPacket) Decode(r *Reader) error { return nil }
