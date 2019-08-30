package cdd

import (
	"bytes"
	"sort"
	"strings"

	"github.com/pkg/errors"

	. "github.com/FlyingOnion/cardcore"
)

var (
	OrderMap = map[string]int{
		"3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13,
		"A": 14, "2": 15,
	}

	// key: text of card group
	// value: index of the "biggest" card
	StraightMap = map[string]int{
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
	errIllegalCG       = errors.New("Illegal card group")
	errCGNotComparable = errors.New("The two card groups are not comparable")
	errUnknown         = errors.New("Unknown error")
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
	return c.CompareWith(another) == -1
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
	v1 := OrderMap[c.Card.Text]
	v2 := OrderMap[another.Card.Text]
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
	cgType, err = ILLEGAL, errUnknown
	if !sort.IsSorted(cg) {
		sort.Sort(cg)
	}

	for _, card := range cg.Cards {
		_, isValidColor := ColorMap[card.Color]
		_, isValidText := OrderMap[card.Text]
		if isValidColor && isValidText {
			continue
		}
		err = errors.Errorf("<%s> is not a valid card", card.String())
		return
	}

	switch cg.Len() {
	case 1:
		cgType, err = SINGLE, nil
	case 2:
		if cg.isPair() {
			cgType, err = PAIR, nil
			break
		}
		err = errors.Errorf("Card group <%s> is not a valid pair", cg.String())
	case 3:
		if cg.isTriple() {
			cgType, err = TRIPLE, nil
			break
		}
		err = errors.Errorf("Card group <%s> is not a valid triple", cg.String())
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
		default:
			err = errors.Errorf("Card group <%s> is not a valid pentuple", cg.String())
		}
	}
	return
}

func (cg cddCardGroup) LessThan(another cddCardGroup) (bool, error) {
	var type1, type2 int
	var err error
	if type1, err = cg.validate(); err != nil {
		return false, errors.Wrap(err, "Illegal card group")
	}
	if type2, err = another.validate(); err != nil {
		return false, errors.Wrap(err, "Illegal card group")
	}
	if type1 == type2 {
		var result bool
		switch type1 {
		case SINGLE, PAIR, TRIPLE, FLUSH:
			result = cg.Cards[cg.Len()-1].LessThan(another.Cards[another.Len()-1])

		case STRAIGHT, STRFLUSH:
			result = cg.Cards[StraightMap[cg.Text()]].LessThan(another.Cards[StraightMap[another.Text()]])

		case SKELETON, KK:
			result = cg.Cards[2].LessThan(another.Cards[2])
		}
		return result, nil
	}
	if type1 >= STRAIGHT && type2 >= STRAIGHT {
		return type1 < type2, nil
	}
	return false, errCGNotComparable
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

func (sortedCG cddCardGroup) isStraight() bool {
	_, ok := StraightMap[sortedCG.Text()]
	return ok
}

func (cg cddCardGroup) isFlush() bool {
	sum := 0
	for _, card := range cg.Cards {
		sum += card.Card.Color
	}
	switch sum {
	case 5 * CARD_COLOR_DIAMOND,
		5 * CARD_COLOR_CLUB,
		5 * CARD_COLOR_HEART,
		5 * CARD_COLOR_SPADE:
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
