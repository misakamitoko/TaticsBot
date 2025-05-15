package fetchtaticsdata

import (
	"TaticsBot/consts"
	"TaticsBot/utils"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func (t *TanticsData) ExtractShortDescAndCost() error {
	nodes, err := t.ExtractData(consts.ShortDecPath)
	if err != nil {
		fmt.Printf("extract data failed, err = %v\n", err)
		return err
	}
	re := regexp.MustCompile(`\d+|[\p{Han}]+(?:\s+[\p{Han}]+)*`)
	for i, node := range nodes {
		text := htmlquery.InnerText(node)
		matches := re.FindAllString(text, -1)
		t.Teams[i].ShortDes = matches[0]
		cost, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Printf("parse cost failed, err = %v\n", err)
			return err
		}
		t.Teams[i].Cost = uint32(cost)
	}
	return nil
}

func (t TanticsData) ExtractData(xpath string) ([]*html.Node, error) {
	reader, err := os.Open(t.DataFile)
	fmt.Printf("open file %s\n", t.DataFile)
	if err != nil {
		fmt.Printf("read html file failed, err = %v\n", err)
		return nil, err
	}
	doc, err := htmlquery.Parse(reader)
	if err != nil {
		fmt.Printf("parse html file failed, err = %v\n", err)
		return nil, err
	}
	nodes := htmlquery.Find(doc, xpath)
	if len(nodes) == 0 {
		fmt.Printf("no match element found\n")
		return nil, fmt.Errorf("no match element found")
	}
	nodes = nodes[0:consts.Teams]
	return nodes, nil
}

func (t *TanticsData) ExtracHeroData() {
	nodes, err := t.ExtractData(consts.HeroNamePath)
	if err != nil {
		fmt.Printf("extract data failed, err = %v\n", err)
		return
	}
	for i, node := range nodes {
		ns := htmlquery.Find(node, "//div[@class='relative']/img")
		for _, n := range ns {
			heroName := htmlquery.SelectAttr(n, "alt")
			imgUrl := htmlquery.SelectAttr(n, "src")
			t.Teams[i].Heros = append(t.Teams[i].Heros, &Hero{string(heroName), string(imgUrl)})
		}
	}
}

func (t *TanticsData) ExtractRateAndRankData() {
	node, err := t.ExtractData(consts.RateAndRankPath)
	if err != nil {
		fmt.Printf("extract data failed, err = %v\n", err)
		return
	}
	for i, node := range node {
		d := utils.SplitString(htmlquery.InnerText(node))
		if len(d) != 4 {
			fmt.Printf("data format error, data = %v\n", d)
			return
		}
		t.Teams[i].TopRate = d[1]
		averageRank, err := strconv.ParseFloat(d[0], 32)
		if err != nil {
			fmt.Printf("parse average rank failed, err = %v\n", err)
			return
		}
		t.Teams[i].AverageRank = float32(averageRank)
		t.Teams[i].TopRate = d[1]
		t.Teams[i].Top4Rate = d[2]
		pickRate := d[3]
		t.Teams[i].PickRate = pickRate
	}
}

func (t *TanticsData) Print() {
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("json marshall err = %v\n", err)
	}
	fmt.Println(string(data))
}
