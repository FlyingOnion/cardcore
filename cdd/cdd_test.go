package cdd

import (
	"testing"

	. "github.com/FlyingOnion/cardcore"
)

func TestStrFlush(t *testing.T) {
	isStraight, isFlush := cddCardGroup{
		Cards: []cddCard{
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "4"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "5"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "7"}},
		},
	}.isStraightOrFlush()
	if isStraight && isFlush {
		t.Log("str flush test succeeded")
		return
	}
	t.Fatal("str flush test failed")
}

func TestStraight(t *testing.T) {
	isStraight, isFlush := cddCardGroup{
		Cards: []cddCard{
			cddCard{Card: Card{CARD_COLOR_CLUB, "3"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "4"}},
			cddCard{Card: Card{CARD_COLOR_SPADE, "5"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "7"}},
		},
	}.isStraightOrFlush()
	if isStraight && !isFlush {
		t.Log("straight test succeeded")
		return
	}
	t.Fatal("straight test failed")
}

func TestFlush(t *testing.T) {
	isStraight, isFlush := cddCardGroup{
		Cards: []cddCard{
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "4"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "5"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "7"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "9"}},
		},
	}.isStraightOrFlush()
	if !isStraight && isFlush {
		t.Log("flush test succeeded")
		return
	}
	t.Fatal("flush test failed")
}

func TestKK(t *testing.T) {
	cgs := []cddCardGroup{
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "3"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "3"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "4"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "A"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "A"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "A"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "A"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "K"}},
			},
		},
	}
	for i, cg := range cgs {
		if cg.isKK() {
			continue
		}
		t.Fatal("KK test failed, id: ", i)
	}
	t.Log("KK test succeeded")
}
