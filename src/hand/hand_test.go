package hand_test

import ("hand"
		"card"
		"testing"
		"fmt"
		"math"
)
/*
func TestAnalyseCards(t *testing.T){
	//test have not Init
	h := hand.GetHand()
	err := h.SetCard(&card.Card{Suit:2,Value:12})
	if err.Error() != "Hand must init first"{
		t.Errorf("test have not Init fail")
	}

	err = h.AnalyseHand()
	if err.Error() != "Hand must init first"{
		t.Errorf("test have not Init fail")
	}

	//test RoyalFlush : 黑桃10 J Q K A  红桃A  梅花A
	h.Init()
	
	h.SetCard(&card.Card{Suit:0,Value:12})
	h.SetCard(&card.Card{Suit:0,Value:11})
	h.SetCard(&card.Card{Suit:0,Value:10})
	h.SetCard(&card.Card{Suit:0,Value:9})
	h.SetCard(&card.Card{Suit:0,Value:8})
	h.SetCard(&card.Card{Suit:1,Value:12})
	h.SetCard(&card.Card{Suit:2,Value:12})

	h.AnalyseHand()

	if h.Level != 10 || h.FinalValue != -1{
		t.Errorf("test RoyalFlush fail")
	}

	//test 8 cards
	err = h.SetCard(&card.Card{Suit:2,Value:12})
	if err.Error() != "after a game, you should init Hand again"{
		t.Errorf("test 8 cards fail")
	}

	//test straight flush
	h.Init()
	h.SetCard(&card.Card{Suit:1,Value:12})
	h.SetCard(&card.Card{Suit:0,Value:11})
	h.SetCard(&card.Card{Suit:0,Value:10})
	h.SetCard(&card.Card{Suit:0,Value:9})
	h.SetCard(&card.Card{Suit:0,Value:8})
	h.SetCard(&card.Card{Suit:0,Value:7})
	h.SetCard(&card.Card{Suit:2,Value:12})

	h.AnalyseHand()

	if h.Level != 9 || h.FinalValue != 13{
		t.Errorf("test straight flush fail")
	}

	//test four of a kind
	h.Init()
	h.SetCard(&card.Card{Suit:0,Value:2})
	h.SetCard(&card.Card{Suit:1,Value:2})
	h.SetCard(&card.Card{Suit:2,Value:2})
	h.SetCard(&card.Card{Suit:3,Value:2})
	h.SetCard(&card.Card{Suit:0,Value:12})
	h.SetCard(&card.Card{Suit:1,Value:12})
	h.SetCard(&card.Card{Suit:2,Value:12})

	h.AnalyseHand()
	
	if h.Level != 8 || h.FinalValue != 2223332{
		t.Errorf("test four of a kind fail")
	}

	//test full house
	h.Init()
	h.SetCard(&card.Card{Suit:0,Value:4})
	h.SetCard(&card.Card{Suit:1,Value:3})
	h.SetCard(&card.Card{Suit:2,Value:2})
	h.SetCard(&card.Card{Suit:3,Value:2})
	h.SetCard(&card.Card{Suit:0,Value:12})
	h.SetCard(&card.Card{Suit:1,Value:12})
	h.SetCard(&card.Card{Suit:2,Value:12})

	h.AnalyseHand()
	
	if h.Level != 7 || h.FinalValue != 13322243{
		t.Errorf("test full house fail")
	}

	//test flush
	h.Init()
	h.SetCard(&card.Card{Suit:0,Value:4})
	h.SetCard(&card.Card{Suit:0,Value:3})
	h.SetCard(&card.Card{Suit:0,Value:2})
	h.SetCard(&card.Card{Suit:0,Value:5})
	h.SetCard(&card.Card{Suit:0,Value:7})
	h.SetCard(&card.Card{Suit:1,Value:12})
	h.SetCard(&card.Card{Suit:2,Value:6})

	h.AnalyseHand()
	
	if h.Level != 6 || h.FinalValue != 376{
		t.Errorf("test flush fail")
	}

	//test straight
	h.Init()
	h.SetCard(&card.Card{Suit:0,Value:2})
	h.SetCard(&card.Card{Suit:1,Value:3})
	h.SetCard(&card.Card{Suit:2,Value:4})
	h.SetCard(&card.Card{Suit:3,Value:5})
	h.SetCard(&card.Card{Suit:0,Value:6})
	h.SetCard(&card.Card{Suit:1,Value:6})
	h.SetCard(&card.Card{Suit:2,Value:6})

	h.AnalyseHand()
	
	if h.Level != 5 || h.FinalValue != 8{
		t.Errorf("test straight fail")
	}

	//test three of a kind
	h.Init()
	h.SetCard(&card.Card{Suit:0,Value:2})
	h.SetCard(&card.Card{Suit:1,Value:3})
	h.SetCard(&card.Card{Suit:2,Value:4})
	h.SetCard(&card.Card{Suit:3,Value:7})
	h.SetCard(&card.Card{Suit:0,Value:6})
	h.SetCard(&card.Card{Suit:1,Value:6})
	h.SetCard(&card.Card{Suit:2,Value:6})

	h.AnalyseHand()
	
	if h.Level != 4 || h.FinalValue != 6667432{
		t.Errorf("test three of a kind fail")
	}

	//test two pairs
	h.Init()
	h.SetCard(&card.Card{Suit:0,Value:2})
	h.SetCard(&card.Card{Suit:1,Value:2})
	h.SetCard(&card.Card{Suit:2,Value:4})
	h.SetCard(&card.Card{Suit:3,Value:4})
	h.SetCard(&card.Card{Suit:0,Value:6})
	h.SetCard(&card.Card{Suit:1,Value:7})
	h.SetCard(&card.Card{Suit:2,Value:8})

	h.AnalyseHand()
	
	if h.Level != 3 || h.FinalValue != 4422876{
		t.Errorf("test two pairs fail")
	}

	//test one pair
	h.Init()
	h.SetCard(&card.Card{Suit:0,Value:2})
	h.SetCard(&card.Card{Suit:1,Value:2})
	h.SetCard(&card.Card{Suit:2,Value:4})
	h.SetCard(&card.Card{Suit:3,Value:6})
	h.SetCard(&card.Card{Suit:0,Value:7})
	h.SetCard(&card.Card{Suit:1,Value:8})
	h.SetCard(&card.Card{Suit:2,Value:10})

	h.AnalyseHand()
	
	if h.Level != 2 || h.FinalValue != 2308764{
		t.Errorf("test one pair fail")
	}

	//test high card
	h.Init()
	h.SetCard(&card.Card{Suit:0,Value:2})
	h.SetCard(&card.Card{Suit:1,Value:3})
	h.SetCard(&card.Card{Suit:2,Value:4})
	h.SetCard(&card.Card{Suit:3,Value:6})
	h.SetCard(&card.Card{Suit:0,Value:9})
	h.SetCard(&card.Card{Suit:1,Value:11})
	h.SetCard(&card.Card{Suit:2,Value:12})

	h.AnalyseHand()
	
	if h.Level != 1 || h.FinalValue != 13196432{
		t.Errorf("test high card fail")
	}
}
*/

func TestEstimateHand(t *testing.T){
	h := hand.GetHand()
	
	h.Init()
	
	h.SetCard(&card.Card{Suit:0,Value:9})
	h.SetCard(&card.Card{Suit:0,Value:11})
	h.SetCard(&card.Card{Suit:0,Value:10})
	h.SetCard(&card.Card{Suit:0,Value:8})
	h.SetCard(&card.Card{Suit:0,Value:6})

	v := h.EstimateHand()
	fmt.Println(v)
	if math.Abs(v-2197.2778)>0.001{
		t.Errorf("estimate royalflush fail")
	}

	
}