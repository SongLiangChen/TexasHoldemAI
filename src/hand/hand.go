//@author: song
//@contact: 462039091@qq.com

package hand

import ("card"
		"errors"
		"fmt"
		"sort"
)


var SuitShift = []int{0, 3, 6, 9}


var StraightValue = []int{15872, 7936, 3968, 1984, 992, 496, 248, 124, 62, 31}

type Hand struct{
	flush int                    
	straight []int             
	handvalue int            
	count []int          
	initilized bool 
	cards card.Cards 
	handsize int        

	Level int 

	FinalValue int

	hole card.Cards
}

func GetHand() *Hand{
	h := new(Hand)
	h.initilized = false
	return h
}

func (h *Hand)Init(){
	h.flush = 0
	h.handvalue = 0
	h.handsize = 0
	h.initilized = true

	h.Level = -1
	h.FinalValue = -1

	h.cards = make(card.Cards, 0, 7)
	h.straight = make([]int, 4)
	h.count = make([]int, 14)
	h.hole = make(card.Cards, 0, 2)
}

func (h *Hand)SetCard(c *card.Card) error{
	if h.initilized == false{
		return errors.New("Hand must init first")
	}
	if h.handsize == 7{
		return errors.New("after a game, you should init Hand again")
	}
	if h.handsize<2 {
		cc := new(card.Card)
		cc.Suit = c.Suit
		cc.Value = c.Value
		h.hole = append(h.hole, cc)
	}
	h.cards= append(h.cards, c)
	h.handsize++
	h.analyCard(c)
	return nil
}

func (h *Hand)AnalyseHand() error{
	if h.initilized == false{
		return errors.New("Hand must init first")
	}
	if h.handsize < 7{
		return errors.New("not enough cards, must have seven!")
	}

	sort.Sort(h.cards)
	tmp := turnToValue(h.cards)

	//判断是否有皇家同花顺
	for i:=0; i<card.SUITSIZE; i++{
		if h.straight[i]&StraightValue[0] == StraightValue[0]{
			h.Level = 10
			return nil
		}
	}
	//判断是否有同花顺
	for i:=0; i<card.SUITSIZE; i++{
		for j:=1; j<len(StraightValue); j++{
			if h.straight[i]&StraightValue[j] == StraightValue[j]{
				h.Level = 9
				h.FinalValue = len(StraightValue)-j+4
				return nil
			}
		}
	}
	//判断四条
	for i:=card.CARDRANK-1; i>=0; i--{
		if h.count[i] == 4{
			h.Level = 8
			h.FinalValue = tmp
			return nil
		}
	}
	//判断葫芦，和四条同理
	for i:=card.CARDRANK-1; i>=0; i--{
		if h.count[i] == 3{
			for j:=0; j<card.CARDRANK; j++{
				if j == i{
					continue
				}
				if h.count[j] >=2{
					h.Level = 7
					h.FinalValue = tmp
					return nil
				}
			}
		}
	}
	/*判断同花*/
	for i:=0; i<card.SUITSIZE; i++{
		tmp := (h.flush>>uint(SuitShift[i])) & 5
		if tmp == 5{
			h.Level = 6
			h.FinalValue = h.straight[i]
			return nil
		}
	}

	//判断顺子
	for i:=0; i<len(StraightValue); i++{
		if h.handvalue&StraightValue[i] == StraightValue[i]{
			h.Level = 5
			h.FinalValue = len(StraightValue)-i+4
			return nil
		}
	}
	//判断三条
	for i:=card.CARDRANK-1; i>=0; i--{
		if h.count[i] == 3{
			h.Level = 4
			h.FinalValue = tmp
			return nil
		}
	}
	/*判断两对*/
	for i:=0; i<card.CARDRANK; i++{
		if h.count[i] == 2{
			for j:=i+1; j<card.CARDRANK; j++{
				if h.count[j] == 2{
					h.Level = 3
					h.FinalValue = tmp
					return nil
				}
			}
		}
	}
	//判断一对
	for i:=0; i<card.CARDRANK; i++{
		if h.count[i] == 2{
			h.Level = 2
			h.FinalValue = tmp
			return nil
		}
	}

	//判断高牌
	h.Level = 1
	h.FinalValue = tmp
	return nil
}

var SUITNAME = []string{"黑桃", "红桃", "梅花", "方块"}
func (h *Hand)ShowHand(){
	//fmt.Printf("%d %d\n",h.Level, h.FinalValue)
	for i:=0; i<len(h.cards); i++{
		fmt.Printf("%s %s, ",SUITNAME[h.cards[i].Suit], RANKNAME[h.cards[i].Value])
	}
	fmt.Println()
}

var RANKNAME = []string{"2","3","4","5","6","7","8","9","10","J","Q","K","A"}
func (h *Hand) ShowHole(){
	fmt.Printf("电脑手牌为：%s %s、%s %s\n",SUITNAME[h.hole[0].Suit], RANKNAME[h.hole[0].Value], SUITNAME[h.hole[1].Suit], RANKNAME[h.hole[1].Value])
}

func turnToValue(cards card.Cards) int{
	res := 0
	for i:=0; i<len(cards); i++{
		res *= 10
		res += cards[i].Value
	}
	return res
}

//返回二进制区间里最大的那张牌
func getHibitPos(a int) int{
	res := 0
	for a>0 {
		a /= 2
		res++
	}
	return res
}

func getCardsFromBinaly(a int) []int{
	var res []int
	count := 1
	for a>0{
		if a%2==1 {
			res = append(res, count)
		}
		a /= 2
		count++
	}
	return res
}

func (h *Hand)analyCard(c *card.Card){

	h.flush += 1<<uint(SuitShift[c.Suit])
	h.straight[c.Suit] |= 1<<uint(c.Value+1)
	h.handvalue |= 1<<uint(c.Value+1)

	if c.Value == 12{
		h.straight[c.Suit] |= 1
		h.handvalue |= 1
	}

	h.count[c.Value]++
	for i:=0; i<h.handsize; i++{
		h.cards[i].Showtime = h.count[h.cards[i].Value]
	}
}

