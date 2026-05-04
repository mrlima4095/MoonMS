package packets

import (
	"github.com/google/uuid"
	"github.com/robogg133/MoonMS/internal/datatypes"
)

const PACKET_RESOURCE_PACK_PUSH int32 = 0x09

type ResourcePackPush struct {
	UUID          uuid.UUID
	URL           string
	Hash          string
	Forced        bool
	PromptMessage datatypes.TextComponent
}

func (*ResourcePackPush) ID() int32 { return PACKET_RESOURCE_PACK_PUSH }

func (p *ResourcePackPush) Encode(w *Writer) error {
	w.WriteUUID(p.UUID)
	w.WriteString(p.URL)
	w.WriteString(p.Hash)
	w.WriteBoolean(p.Forced)

	if p.PromptMessage == "" {
		return w.WriteBoolean(false)
	}
	w.WriteBoolean(true)

	return w.WriteString(string(p.PromptMessage))
}

func (p *ResourcePackPush) Decode(r *Reader) error {
	var err error

	p.UUID, err = r.ReadUUID()
	if err != nil {
		return err
	}

	p.URL, err = r.ReadString()
	if err != nil {
		return err
	}

	p.Hash, err = r.ReadString()
	if err != nil {
		return err
	}

	p.Forced, err = r.ReadBoolean()
	if err != nil {
		return err
	}

	val, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	switch val {
	case true:
		s, err := r.ReadString()
		if err != nil {
			return err
		}
		p.PromptMessage = datatypes.TextComponent(s)
	}

	return nil
}
