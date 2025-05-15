package main

import (
	vo "TaticsBot/fetchtaticsdata"
)

func main() {
	td := vo.NewTaticsData()
	td.ExtracHeroData()
	td.ExtractShortDescAndCost()
	td.ExtractRateAndRankData()
	td.Print()
}
