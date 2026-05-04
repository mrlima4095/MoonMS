package packets

import "github.com/robogg133/MoonMS/internal/datatypes"

const PAKCET_STORE_COOKIE int32 = 0x0A

type StoreCookiePacket struct {
	Key     datatypes.Identifier
	Payload []byte
}

func (*StoreCookiePacket) ID() int32 { return PAKCET_STORE_COOKIE }

func (p *StoreCookiePacket) Encode(w *Writer) error {
	if err := w.WriteString(p.Key.String()); err != nil {
		return err
	}
	return w.WritePrefixed(p.Payload)
}

func (p *StoreCookiePacket) Decode(r *Reader) error {
	str, err := r.ReadString()
	if err != nil {
		return err
	}

	p.Key, err = datatypes.NewIdentifier(str)
	if err != nil {
		return err
	}

	p.Payload, err = r.ReadPrefixed()
	return err
}
