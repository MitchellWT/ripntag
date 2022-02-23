package enums

import "strings"

type SearchMethod int
type UndefinedSearchMethodError struct{}

func (usme UndefinedSearchMethodError) Error() string {
	return "Error: Undefined search method provided!"
}

const (
	SearchMethodUndefined = -1
	Barcode               = 0
	ArtistAlbum           = 1
)

func ToSearchMethod(s string) (SearchMethod, error) {
	s = strings.ToLower(s)
	switch s {
	case "barcode":
		return Barcode, nil
	case "artist-album":
		return ArtistAlbum, nil
	default:
		return SearchMethodUndefined, UndefinedSearchMethodError{}
	}
}

func (sm SearchMethod) String() string {
	switch sm {
	case Barcode:
		return "barcode"
	case ArtistAlbum:
		return "artist-album"
	default:
		return "nil"
	}
}
