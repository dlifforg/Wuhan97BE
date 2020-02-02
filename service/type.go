package service

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	tokenExpiredDate = 3600 * 1 * time.Second
)

type fakeNewsItem struct {
	Body        string `json:"body"`
	ID          int    `json:"id"`
	MainSummary string `json:"mainSummary"`
	RumorType   int    `json:"rumorType"`
	Score       int    `json:"score"`
	SourceURL   string `json:"sourceUrl"`
	Summary     string `json:"summary"`
	Title       string `json:"title"`
}
type careItem struct {
	ContentType  int    `json:"contentType"`
	CreateTime   int64  `json:"createTime"`
	Deleted      bool   `json:"deleted"`
	ID           int    `json:"id"`
	ImgURL       string `json:"imgUrl"`
	LinkURL      string `json:"linkUrl"`
	ModifyTime   int64  `json:"modifyTime"`
	Operator     string `json:"operator"`
	RecordStatus int    `json:"recordStatus"`
	Sort         int    `json:"sort"`
	Title        string `json:"title"`
}
type newsItem struct {
	AdoptType        int    `json:"adoptType"`
	CreateTime       int64  `json:"createTime"`
	DataInfoOperator string `json:"dataInfoOperator"`
	DataInfoState    int    `json:"dataInfoState"`
	DataInfoTime     int64  `json:"dataInfoTime"`
	EntryWay         int    `json:"entryWay"`
	ID               int    `json:"id"`
	InfoSource       string `json:"infoSource"`
	InfoType         int    `json:"infoType"`
	ModifyTime       int64  `json:"modifyTime"`
	ProvinceID       string `json:"provinceId"`
	ProvinceName     string `json:"provinceName"`
	PubDate          int64  `json:"pubDate"`
	PubDateStr       string `json:"pubDateStr"`
	SourceURL        string `json:"sourceUrl"`
	Summary          string `json:"summary"`
	Title            string `json:"title"`
}

func save(key string, value interface{}) {
	rc.Set(key, value, tokenExpiredDate)

}

func (news newsItem) saveToRedis() {
	key := "news." + strconv.Itoa(news.ID)
	value, _ := json.Marshal(news)
	save(key, value)
}

type source struct {
	url      string
	dataType string
	data     struct {
		Code string          `json:"code"`
		Data json.RawMessage `json:"data"`
	}
}
