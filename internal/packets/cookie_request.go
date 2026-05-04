package packets

type CookieRequestPacket struct {
	Identifier string
}

func (*CookieRequestPacket) ID() int32 { return PACKET_COOKIE_REQUEST }

func (s *CookieRequestPacket) Encode(w *Writer) error {
	err := w.WriteString(s.Identifier)
	return err
}

func (s *CookieRequestPacket) Decode(r *Reader) error {

	var err error
	s.Identifier, err = r.ReadString()
	if err != nil {
		return err
	}
	return err
}
