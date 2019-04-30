package core

const (
	CARD_COLOR_DIAMOND = 1
	CARD_COLOR_CLUB    = 8
	CARD_COLOR_HEART   = 64
	CARD_COLOR_SPADE   = 512
)

var (
	ColorMap = map[int]string{
		CARD_COLOR_DIAMOND: "Diamond",
		CARD_COLOR_CLUB:    "Club",
		CARD_COLOR_HEART:   "Heart",
		CARD_COLOR_SPADE:   "Spade",
	}
)

type Card struct {
	Color int
	Text  string
}
