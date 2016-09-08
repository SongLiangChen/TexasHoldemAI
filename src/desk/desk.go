package desk

import ("fmt"
		"ai"
		"dealmachine"
		//"card"
		"player"
		//"hand"
)


var CARDTYPE = []string{"", "高牌", "一对", "两对", "三条", "顺子", "同花", "葫芦", "四条", "同花顺", "皇家同花顺"}
var gambpool int
var index int
var dm *dealmachine.DealMachine
var playerAllIn bool
var aiAllin bool
var current int

func PlayGame(){

	dm = dealmachine.GetDealMachine()
	dm.Init()

	aai := ai.GetAI()
	p := player.GetPlayer()

	current = 0
	
	for {
		if p.GetChip()<=0 || aai.GetChip()<=0{
			break
		}

		fmt.Printf("\n\n")
		fmt.Println("新局开始！")
		dm.Shuffle()
		aai.Init()
		p.Init()

		gambpool = 0
		index = 0
		aiAllin = false
		playerAllIn = false

		//盲注
		p.Blind(100)
		aai.Blind(100)
		gambpool += 200

		//hole下注
		p.SetHole(dm.Deal(), dm.Deal())
		aai.SetHole(dm.Deal(), dm.Deal())
		if p.GetChip()==0 || aai.GetChip()==0{
			index = 2
			allin(p, aai)
			continue
		}

		showTable(p, aai)
		bet := chipIn(p, aai, current)
		if bet == 0{
			continue
		}
		gambpool += 2*bet
		p.Call(bet)
		aai.Call(bet)
		index += 2
		if playerAllIn==true || aiAllin==true{
			allin(p, aai)
			continue
		}

		//flop下注
		c1 := dm.Deal()
		c2 := dm.Deal()
		c3 := dm.Deal()

		p.SetFlop(c1, c2, c3)
		aai.SetFlop(c1, c2, c3)
		showTable(p, aai)
		bet = chipIn(p, aai, current)
		if bet == 0{
			continue
		}
		gambpool += 2*bet
		p.Call(bet)
		aai.Call(bet)
		index += 3
		if playerAllIn==true || aiAllin==true{
			allin(p, aai)
			continue
		}

		//turn下注
		c1 = dm.Deal()
		p.SetTurn(c1)
		aai.SetTurn(c1)
		showTable(p, aai)
		bet = chipIn(p, aai, current)
		if bet == 0{
			continue
		}
		gambpool += 2*bet
		p.Call(bet)
		aai.Call(bet)
		index++
		if playerAllIn==true || aiAllin==true{
			allin(p, aai)
			continue
		}

		//river下注
		c1 = dm.Deal()
		p.SetRiver(c1)
		aai.SetRiver(c1)
		showTable(p, aai)
		bet = chipIn(p, aai, current)
		if bet == 0{
			continue
		}
		gambpool += 2*bet
		p.Call(bet)
		aai.Call(bet)
		index++

		showdown(p, aai)
	}

	if p.GetChip()>0 {
		fmt.Println("你真牛")
	}else{
		fmt.Println("你太弱了")
	}

	fmt.Println("请按任意键退出")
	var tt string
	fmt.Scan(&tt)
}

func chipIn(p *player.Player, aai *ai.AI, mod int) int{

	var bet int
	var fcr int
	if mod == 0{
		fmt.Println("请下注，弃牌请输入0")
		fmt.Scan(&bet)
		if bet == 0{
			current = 1
			aai.Take(gambpool)
			fmt.Println("玩家弃牌，电脑获胜")
			return 0
		}
		if bet >= p.GetChip(){
			bet = p.GetChip()
			playerAllIn = true
		}
		fcr = aai.FCR(bet, gambpool, dm)
		if fcr == 0{
			current = 0
			p.Take(gambpool)
			fmt.Println("电脑弃牌，玩家获胜")
			return 0
		}
		if fcr >= aai.GetChip(){
			fmt.Println("来吧！ALL IN 我跟你拼了")
			fcr = aai.GetChip()
			aiAllin = true
		}
		if playerAllIn == false && fcr > bet{
			current = current^1
			var yn string
			for{
				fmt.Printf("电脑加注为：%d,是否跟注？y/n\n",fcr)
				fmt.Scan(&yn)
				if yn == "y"{
					break
				}else if yn == "n"{
					current = 1
					aai.Take(gambpool)
					fmt.Println("玩家弃牌，电脑获胜")
					return 0
				}
			}
			if fcr >= p.GetChip(){
				bet = p.GetChip()
				playerAllIn = true
			}else{
				bet = fcr
			}
		}
	}else{
		fcr = aai.FCR(0, gambpool, dm)
		if fcr == 0{
			current = 0
			p.Take(gambpool)
			fmt.Println("电脑弃牌，玩家获胜")
			return 0
		}
		if fcr >= aai.GetChip(){
			fmt.Println("来吧！ALL IN 我跟你拼了")
			fcr = aai.GetChip()
			aiAllin = true
		}
		fmt.Printf("电脑下注为：%d,请跟注或者加注,跟注为0表示弃牌\n",fcr)
		for {
			fmt.Scan(&bet)
			if bet==0 {
				current = 1
				aai.Take(gambpool)
				fmt.Println("玩家弃牌，电脑获胜")
				return 0
			}
			if bet >= p.GetChip(){
				break
			}
			if bet < fcr{
				fmt.Println("跟注不能小于电脑下注")
			}else{
				break
			}
		}
		if bet >= p.GetChip(){
			bet = p.GetChip()
			playerAllIn = true
		}
		if fcr < bet && aiAllin == false{
			tmp := aai.FCR(bet, gambpool+bet, dm)
			if tmp==0 {
				p.Take(gambpool)
				current = 0
				fmt.Println("电脑弃牌，玩家获胜")
				return 0
			}
			if bet>=aai.GetChip(){
				fmt.Println("来吧！ALL IN 我跟你拼了")
				fcr = aai.GetChip()
				aiAllin = true
			}else{
				fcr = bet
			}
		}
	}
	if aiAllin == true || playerAllIn == true{
		if bet < fcr{
			return bet
		}
		return fcr
	}
	return bet
}

func allin(p *player.Player, aai *ai.AI){
	fmt.Println("*****************All in********************")
	if index == 2{
		c1 := dm.Deal()
		c2 := dm.Deal()
		c3 := dm.Deal()
		p.SetFlop(c1, c2, c3)
		aai.SetFlop(c1, c2, c3)
		index += 3
	}
	if index == 5{
		c := dm.Deal()
		p.SetTurn(c)
		aai.SetTurn(c)
		index++
	}
	if index == 6 {
		c := dm.Deal()
		p.SetRiver(c)
		aai.SetRiver(c)
	}
	showdown(p, aai)
}

func showdown(p *player.Player, aai *ai.AI){
	p.DealOver()
	aai.DealOver()

	if p.GetLevel() > aai.GetLevel(){
		current = 0
		p.Take(gambpool)
		fmt.Println("玩家获胜")
	}else if p.GetLevel() == aai.GetLevel() && p.GetFinalValue() > aai.GetFinalValue(){
		current = 0
		p.Take(gambpool)
		fmt.Println("玩家获胜")
	}else if p.GetLevel() == aai.GetLevel() && p.GetFinalValue() == aai.GetFinalValue(){
		p.Take(gambpool/2)
		aai.Take(gambpool/2)
		fmt.Println("平局")
	}else{
		current = 1
		aai.Take(gambpool)
		fmt.Println("电脑获胜")
	}
	aai.ShowComminityCards()
	aai.ShowHole()
	fmt.Printf("电脑牌型为：%s\n",CARDTYPE[aai.GetLevel()])
}

func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}

func showTable(p *player.Player, aai *ai.AI){
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("***********************************************************\n")
	aai.ShowChip()
	fmt.Println()
	fmt.Println()
	fmt.Printf("当前赌池金额为：%d\n",gambpool)
	aai.ShowComminityCards()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	p.ShowChip()
	p.ShowHole()
	fmt.Printf("***********************************************************\n")
}