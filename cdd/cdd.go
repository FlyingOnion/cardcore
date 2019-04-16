package cdd

import (
	"bytes"
	"sort"
	"strings"

	"github.com/pkg/errors"

	. "github.com/FlyingOnion/cardcore"
)

var (
	orderMap = map[string]int{
		"3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13,
		"A": 14, "2": 15, "JokerB": 16, "JokerR": 17,
	}

	// key: text of card group
	// value: index of the "biggest" card
	straightMap = map[string]int{
		"345A2":  2,
		"34562":  3,
		"34567":  4,
		"45678":  4,
		"56789":  4,
		"678910": 4,
		"78910J": 4,
		"8910JQ": 4,
		"910JQK": 4,
		"10JQKA": 4,
	}
)
var (
	errHasJoker      = errors.New("joker(s) found in card group")
	errIllegal       = errors.New("illegal card group")
	errNotComparable = errors.New("the two card groups are not comparable")
	errUnknown       = errors.New("unknown error")
)

const (
	// card group type
	ILLEGAL = iota
	SINGLE
	PAIR
	TRIPLE
	STRAIGHT
	FLUSH
	SKELETON
	KK
	STRFLUSH
)

type cddCard struct {
	Card
}

func (c cddCard) LessThan(another cddCard) bool {
	result := c.CompareTextWith(another)
	return result == -1 || (result == 0 && c.Color < another.Color)
}

func (c cddCard) CompareWith(another cddCard) (result int) {
	result = c.CompareTextWith(another)
	if result == 0 {
		switch {
		case c.Color < another.Color:
			result = -1
		case c.Color > another.Color:
			result = 1
		}
	}
	return
}

func (c cddCard) CompareTextWith(another cddCard) int {
	v1 := orderMap[c.Card.Text]
	v2 := orderMap[another.Card.Text]
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return 1
	}
	return 0
}

func (c cddCard) String() string {
	return ColorMap[c.Color] + "-" + c.Text
}

type cddCardGroup struct {
	Cards []cddCard
}

// implement sort.Interface
func (cg cddCardGroup) Len() int {
	return len(cg.Cards)
}
func (cg cddCardGroup) Less(i, j int) bool {
	return cg.Cards[i].LessThan(cg.Cards[j])
}
func (cg cddCardGroup) Swap(i, j int) {
	cg.Cards[i], cg.Cards[j] = cg.Cards[j], cg.Cards[i]
}

func (cg cddCardGroup) validate() (cgType int, err error) {
	cgType, err = ILLEGAL, errIllegal
	if !sort.IsSorted(cg) {
		sort.Sort(cg)
	}

	// no jokers in cdd
	for _, card := range cg.Cards {
		if hasJoker := orderMap[card.Text] >= 16; hasJoker {
			err = errHasJoker
			return
		}
	}

	switch cg.Len() {
	case 1:
		cgType, err = SINGLE, nil
	case 2:
		if cg.isPair() {
			cgType, err = PAIR, nil
		}
	case 3:
		if cg.isTriple() {
			cgType, err = TRIPLE, nil
		}
	case 5:

		b1, b2 := cg.isStraightOrFlush()
		switch {
		case b1 && b2:
			cgType, err = STRFLUSH, nil
		case b1:
			cgType, err = STRAIGHT, nil
		case b2:
			cgType, err = FLUSH, nil
		case cg.isKK():
			cgType, err = KK, nil
		case cg.isSkeleton():
			cgType, err = SKELETON, nil
		}
	}
	return
}

func (cg cddCardGroup) LessThan(another cddCardGroup) (result bool, err error) {
	result, err = false, errIllegal
	var type1, type2 int
	if type1, err = cg.validate(); err != nil {
		return
	}
	if type2, err = another.validate(); err != nil {
		return
	}
	if type1 == type2 {
		err = nil
		switch type1 {
		case SINGLE, PAIR, TRIPLE:
			result = cg.Cards[cg.Len()-1].LessThan(another.Cards[another.Len()-1])

		case STRAIGHT, STRFLUSH:
			result = cg.Cards[straightMap[cg.Text()]].LessThan(another.Cards[straightMap[another.Text()]])

		case SKELETON, KK:
			result = cg.Cards[2].LessThan(another.Cards[2])
		}
		return
	}
	if type1 >= STRAIGHT && type2 >= STRAIGHT {
		result, err = type1 < type2, nil
		return
	}
	return
}

func (cg cddCardGroup) isPair() bool {
	return cg.Len() == 2 && strings.Compare(cg.Cards[0].Text, cg.Cards[1].Text) == 0
}

func (sortedCG cddCardGroup) isTriple() bool {
	return sortedCG.Len() == 3 && strings.Compare(sortedCG.Cards[0].Text, sortedCG.Cards[2].Text) == 0
}

func (sortedCG cddCardGroup) isQuadruple() bool {
	return sortedCG.Len() == 4 && strings.Compare(sortedCG.Cards[0].Text, sortedCG.Cards[3].Text) == 0
}

func (cg cddCardGroup) isStraightOrFlush() (bool, bool) {
	sort.Sort(cg)
	return cg.isStraight(), cg.isFlush()
}

func (sortCG cddCardGroup) isStraight() bool {
	_, ok := straightMap[sortCG.Text()]
	return ok
}

func (cg cddCardGroup) isFlush() bool {
	sum := 0
	for _, card := range cg.Cards {
		sum += card.Card.Color
	}
	switch sum {
	case 5, 40, 320, 2560:
		return true
	}
	return false
}

func (sortedCG cddCardGroup) isSkeleton() bool {
	return (cddCardGroup{sortedCG.Cards[0:3]}.isTriple() && cddCardGroup{sortedCG.Cards[3:5]}.isPair()) ||
		(cddCardGroup{sortedCG.Cards[0:2]}.isPair() && cddCardGroup{sortedCG.Cards[2:5]}.isTriple())
}

func (sortedCG cddCardGroup) isKK() bool {
	return cddCardGroup{sortedCG.Cards[0:4]}.isQuadruple() || cddCardGroup{sortedCG.Cards[1:5]}.isQuadruple()
}

// text only
func (cg cddCardGroup) Text() string {
	sort.Sort(cg)
	var buf bytes.Buffer
	for _, card := range cg.Cards {
		buf.WriteString(card.Text)
	}
	return buf.String()
}

// color and text
func (cg cddCardGroup) String() string {
	sort.Sort(cg)
	var buf bytes.Buffer
	for i, card := range cg.Cards {
		buf.WriteString(card.String())
		if i < cg.Len()-1 {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}
