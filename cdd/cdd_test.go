package cdd

import (
	"sort"
	"strings"
	"testing"

	. "github.com/FlyingOnion/cardcore"
)

func TestSort(t *testing.T) {
	cg := cddCardGroup{
		Cards: []cddCard{
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
			cddCard{Card: Card{CARD_COLOR_HEART, "4"}},
			cddCard{Card: Card{CARD_COLOR_SPADE, "5"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "2"}},
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
		},
	}
	sort.Sort(cg)
	if strings.Compare(cg.Text(), "34562") != 0 {
		t.Fatal("sort test failed")
	}
	t.Log("sort test passed")
}

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
		t.Log("str flush test passed")
		return
	}
	t.Fatal("str flush test failed")
}

func TestStraight(t *testing.T) {
	cgs := []cddCardGroup{
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "4"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "5"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "2"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "4"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "5"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "7"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "A"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "5"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "4"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "2"}},
			},
		},
	}
	for i, cg := range cgs {
		isStraight, isFlush := cg.isStraightOrFlush()
		if isStraight && !isFlush {
			continue
		}
		t.Fatal("straight test failed, id: ", i)
	}
	t.Log("straight test passed")
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
		t.Log("flush test passed")
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
	t.Log("KK test passed")
}

func TestSkeleton(t *testing.T) {
	cgs := []cddCardGroup{
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "3"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "4"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "4"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "A"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "A"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "A"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "K"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "K"}},
			},
		},
	}
	for i, cg := range cgs {
		if cg.isSkeleton() {
			continue
		}
		t.Fatal("Skeleton test failed, id: ", i)
	}
	t.Log("Skeleton test passed")
}

func TestLessThan(t *testing.T) {
	cgs := []cddCardGroup{
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_HEART, "3"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_SPADE, "K"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "3"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_CLUB, "3"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "3"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "3"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_CLUB, "10"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "10"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "3"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "8"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "8"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "8"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "3"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "4"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "4"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "A"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "A"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "A"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "K"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "K"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "4"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "5"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "7"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "9"}},
			},
		},
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
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "4"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "5"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "2"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "4"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "5"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "7"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
			},
		},
	}
	for i := 0; i < len(cgs); i += 2 {
		result, err := cgs[i].LessThan(cgs[i+1])
		if err != nil {
			t.Fatal(err)
		}
		if result {
			continue
		}
		t.Fatal("Comparison test failed, id: ", i, i+1)
	}
	t.Log("Comparison test passed")
}

func TestValidate(t *testing.T) {
	cgs := []cddCardGroup{
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_HEART, "4"}},
				cddCard{Card: Card{CARD_COLOR_SPADE, "5"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "2"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
			},
		},
		cddCardGroup{
			Cards: []cddCard{
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
				cddCard{Card: Card{CARD_COLOR_CLUB, "4"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "5"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "7"}},
				cddCard{Card: Card{CARD_COLOR_DIAMOND, "6"}},
			},
		},
	}
	cgTypes := []int{STRAIGHT, STRAIGHT}
	for i, cg := range cgs {
		cgType, err := cg.validate()
		if err == nil && cgType == cgTypes[i] {
			continue
		}
		t.Fatalf("Validate test failed, expected: %d, got: %d, id: %d", cgTypes[i], cgType, i)
	}
	t.Log("Validate test passed")
}

func TestString(t *testing.T) {
	cg := cddCardGroup{
		Cards: []cddCard{
			cddCard{Card: Card{CARD_COLOR_DIAMOND, "3"}},
			cddCard{Card: Card{CARD_COLOR_HEART, "3"}},
		},
	}

	if cg.Cards[0].String() != "Diamond-3" {
		t.Fatal("String test failed, expected: Diamond-3, got: ", cg.Cards[0].String())
	}

	if cg.Text() != "33" {
		t.Fatal("String test failed, expected: 33, got: ", cg.Text())
	}

	if cg.String() != "Diamond-3 Heart-3" {
		t.Fatal("String test failed, expected: Diamond-3 Heart-3, got: ", cg.String())
	}
	t.Log("String test passed")
}
