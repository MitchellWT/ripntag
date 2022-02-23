package enums

import "strings"

type TagType int
type UndefinedTagTypeError struct{}

func (utte UndefinedTagTypeError) Error() string {
	return "Error: Undefined tag type provided!"
}

const (
	TagTypeUndefined = -1
	Rip              = 0
	FileName         = 1
)

func ToTagType(s string) (TagType, error) {
	s = strings.ToLower(s)
	switch s {
	case "rip":
		return Rip, nil
	case "file-name":
		return FileName, nil
	default:
		return TagTypeUndefined, UndefinedTagTypeError{}
	}
}

func (tt TagType) String() string {
	switch tt {
	case Rip:
		return "rip"
	case FileName:
		return "file-name"
	default:
		return "nil"
	}
}
