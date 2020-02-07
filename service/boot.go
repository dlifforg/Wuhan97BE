package service

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	sources []source
	rc      *redis.Client
)

func Redis() *redis.Client{
	return rc
}
func register(ss ...source) {
	sources = append(sources, ss...)
}

func init() {
	// init
	rc = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
func Init() {
	ss := []source{
		{
			url:      "https://file1.dxycdn.com/2020/0127/794/3393185296027391740-115.json",
			dataType: "News",
		},
		{
			url:      "https://file1.dxycdn.com/2020/0127/908/3393185296027391755-115.json",
			dataType: "Care",
		},
		{
			url:      "https://file1.dxycdn.com/2020/0127/797/3393185293879908067-115.json",
			dataType: "FakeNews",
		},
		{
			url:      "https://service-f9fjwngp-1252021671.bj.apigw.tencentcs.com/release/pneumonia",
			dataType: "Index",
		},
	}
	register(ss...)
	Load()

}

func Load() {
	for _, item := range sources {

		log.Printf("starting request :%s",item.dataType)

		resp, err := http.Get(item.url)

		if err != nil {
			return
		}
		defer resp.Body.Close()
		result, err := ioutil.ReadAll(resp.Body)

		_ = json.Unmarshal(result, &item.data)

		var Z []redis.Z
		var key string

		switch item.dataType {
		case "News":
			var data []newsItem
			key = "News"
			_ = json.Unmarshal(item.data.Data, &data)
			for _, i := range data {
				value, _ := json.Marshal(i)
				Z = append(Z, redis.Z{
					Score:  float64(i.ID),
					Member: value,
				})
			}
			rc.ZAdd(key, Z...)
			break
		case "Care":
			var data []careItem
			key = "Care"
			_ = json.Unmarshal(item.data.Data, &data)
			for _, i := range data {
				value, _ := json.Marshal(i)
				Z = append(Z, redis.Z{
					Score:  float64(i.ID),
					Member: value,
				})
			}
			rc.ZAdd(key, Z...)
			break
		case "FakeNews":
			var data []fakeNewsItem
			key = "FakeNews"
			_ = json.Unmarshal(item.data.Data, &data)
			for _, i := range data {
				value, _ := json.Marshal(i)
				Z = append(Z, redis.Z{
					Score:  float64(i.ID),
					Member: value,
				})
			}
			rc.ZAdd(key, Z...)
			break
		case "Index":
			key = "Index"
			value,_:=json.Marshal(item.data.Data)
			rc.Set(key,value,0)
			break
		default:
		}
		log.Printf("request finished :%s",item.dataType)

	}
}
