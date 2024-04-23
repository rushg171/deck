//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Heart
	Club
	Joker
)

var suits = [...]Suit{Spade, Heart, Diamond, Club}

type Rank uint8

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

type Deck []Card

func New(opts ...func(cards []Card) []Card) []Card {
	var cards Deck
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			cards = append(cards, Card{s, r})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func NoOfDecks(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

func FilterOut(f func(card Card) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

func AddJokers(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))

	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

func Sort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absValue(cards[i]) < absValue(cards[j])
	}
}

func absValue(card Card) int {
	return int(card.Suit)*int(maxRank) + int(card.Rank)
}
