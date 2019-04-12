package core

const (
	CARD_COLOR_DIAMOND = 1
	CARD_COLOR_CLUB    = 8
	CARD_COLOR_HEART   = 64
	CARD_COLOR_SPADE   = 512
)

type Card struct {
	Color int
	Text  string
}
