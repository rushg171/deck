package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Joker})
	fmt.Println(Card{Suit: Diamond, Rank: Two})
	fmt.Println(Card{Suit: Spade, Rank: Seven})
	fmt.Println(Card{Suit: Heart, Rank: King})
	fmt.Println(Card{Suit: Club, Rank: Ace})
	//Output:
	//Joker
	//Two of Diamonds
	//Seven of Spades
	//King of Hearts
	//Ace of Clubs
}

func TestNewDeck(t *testing.T) {
	cards := New()

	if len(cards) != 52 {
		t.Error("Wrong Number of Cards in the new Deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := []Card{
		{Diamond, Two},
		{Spade, King},
	}

	DefaultSort(cards)

	exp := Card{Suit: Spade, Rank: King}
	if cards[0] != exp {
		t.Error("Expected King of Spades in first position. Received:", cards[0])
	}
}

func TestCustomSort(t *testing.T) {
	cards := New(Sort(Less))

	exp := Card{Suit: Spade, Rank: Ace}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades in first position. Received:", cards[0])
	}

}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	exp := New()
	//[40 35 50 0
	first := exp[40]
	second := exp[35]

	cards := New(Shuffle)
	if first != cards[0] || second != cards[1] {
		t.Error("Shuffle Broken")
	}
}

func TestJokers(t *testing.T) {
	cards := New(AddJokers(5), Shuffle)
	if len(cards) != 52+5 {
		t.Error("Wrong Number of Cards in the new Deck")
	}
}

func TestFilter(t *testing.T) {
	f := func(card Card) bool {
		if card.Rank == Two {
			return true
		}
		return false
	}
	cards := New(AddJokers(5), FilterOut(f))
	if len(cards) != 52+5-4 {
		t.Error("Wrong Number of Cards in the new Deck")
	}
}

func TestNoOfDecks(t *testing.T) {
	f := func(card Card) bool {
		if card.Rank == Two || card.Rank == Three {
			return true
		}
		return false
	}
	cards := New(FilterOut(f), NoOfDecks(3), AddJokers(5))
	if len(cards) != (52-8)*3+5 {
		t.Error("Wrong Number of Cards in the new Deck")
	}
}
