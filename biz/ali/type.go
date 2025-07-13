package ali

/**
@description
@date: 07/13 19:24
@author Gk
**/

type OcrData struct {
	AlgoVersion    string `json:"algo_version"`
	Angle          int    `json:"angle"`
	Content        string `json:"content"`
	Height         int    `json:"height"`
	OrgHeight      int    `json:"orgHeight"`
	OrgWidth       int    `json:"orgWidth"`
	PrismVersion   string `json:"prism_version"`
	PrismWnum      int    `json:"prism_wnum"`
	PrismWordsInfo []struct {
		Angle     int `json:"angle"`
		Direction int `json:"direction"`
		Height    int `json:"height"`
		Pos       []struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"pos"`
		Prob  int    `json:"prob"`
		Width int    `json:"width"`
		Word  string `json:"word"`
		X     int    `json:"x"`
		Y     int    `json:"y"`
	} `json:"prism_wordsInfo"`
	Width int `json:"width"`
}

type I18nProduct map[string]LocalProduct

type LocalProduct struct {
	Title    string   `json:"title"`
	Desc     string   `json:"desc"`
	DescAll  []string `json:"desc_all,omitempty"` // ocr出来的一组数据,喂给ai返回 title 和desc
	BaseName string   `json:"base_name"`
	Price    string   `json:"price"`
	Cate     string   `json:"cate"`
	Dir      string   `json:"dir"`
}
