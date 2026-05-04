package packets

import "github.com/google/uuid"

const PACKET_RESOURCE_PACK_POP int32 = 0x08

type ResourcePackPop struct {
	UUID *uuid.UUID
}

func (*ResourcePackPop) ID() int32 { return PACKET_RESOURCE_PACK_POP }

func (p *ResourcePackPop) Encode(w *Writer) error {

	if p.UUID == nil {
		return w.WriteBoolean(false)
	}

	if err := w.WriteBoolean(true); err != nil {
		return err
	}

	return w.WriteUUID(*p.UUID)
}

func (p *ResourcePackPop) Decode(r *Reader) error {
	id, err := r.ReadUUID()
	if err != nil {
		return err
	}

	p.UUID = &id
	return nil
}
