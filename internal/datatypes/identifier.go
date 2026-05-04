package datatypes

import (
	"fmt"
	"regexp"
	"strings"
)

type Identifier struct {
	namespace string
	value     string
}

var exp = regexp.MustCompile(`[a-z0-9.-_]`)

func (i *Identifier) Namespace() string { return i.value }

func (i *Identifier) SetNamespace(s string) error {
	s = strings.ReplaceAll(s, " ", "_")

	if !exp.MatchString(s) {
		return fmt.Errorf("invalid namespace: %s", s)
	}

	i.value = s
	return nil
}

func (i *Identifier) Value() string { return i.value }

func (i *Identifier) SetValue(s string) error {
	s = strings.ReplaceAll(s, " ", "_")

	if !exp.MatchString(s) {
		return fmt.Errorf("invalid value: %s", s)
	}

	i.value = s
	return nil
}

func NewIdentifier(s string) (Identifier, error) {
	sep := strings.Split(s, ":")
	switch len(sep) {
	case 1:
		tmp := sep[0]
		sep = make([]string, 2)
		sep[0] = "minecraft"
		sep[1] = tmp
	default:
		return Identifier{}, fmt.Errorf("invalid number of separators: %d", len(sep))
	}

	var a Identifier

	if err := a.SetNamespace(sep[0]); err != nil {
		return Identifier{}, err
	}
	if err := a.SetValue(sep[1]); err != nil {
		return Identifier{}, err
	}

	return Identifier{}, nil
}

func (i *Identifier) String() string {
	return i.namespace + ":" + i.value
}
