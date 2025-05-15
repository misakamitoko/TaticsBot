package fetchtaticsdata

import "TaticsBot/consts"

type Hero struct {
	Name   string
	ImgUrl string // 图片地址
}

type Team struct {
	Type        string  // 强度，如op
	AverageRank float32 // 平均名次
	TopRate     string  // 登顶率
	Top4Rate    string  // 前四率
	PickRate    string  // 挑选率
	Heros       []*Hero
	ShortDes    string   // 简单描述 羁绊+主副c
	Cost        uint32   // 阵容总共花费
	DetailDes   []string //运营细节描述
	DetailLink  string   // 阵容详情页面
}

type TanticsData struct {
	DataFile string //数据文件
	Teams    []Team
}

func NewTaticsData() *TanticsData {
	return &TanticsData{
		DataFile: consts.DataFilePath,
		Teams:    make([]Team, consts.Teams),
	}
}
