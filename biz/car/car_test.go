package car

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"Guoaking/gotools/tools"
)

/**
@description
@date: 05/31 19:53
@author Gk
**/

func TestA(t *testing.T) {

	carList := getC(1, 80)
	var res []SearchShSkuInfo

	PageSp(carList.Total, 80, func(page, limit int) {
		carList = getC(page, limit)
		res = append(res, carList.SearchShSkuInfoList...)
		fmt.Printf("output:len: %v\n", len(res))
	})

	fmt.Printf("output:all len: %v\n", len(res))
	marshal, _ := json.Marshal(res)
	fileName := fmt.Sprintf("car_%v.json", "all")
	err := os.WriteFile(fileName, marshal, 0600)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}
}

func TestJ2C(t *testing.T) {
	carlist := getData("car_all.txt")

	var updatedRecords [][]string
	header := []string{
		"title", "sku_id",
		"shop_id",
		"car_id",
		"car_name",
		//"series_id",
		"series_name",
		//"brand_id",
		"brand_name",
		"car_year", "car_source_city_name", "brand_source_city_name",
		//"spu_id", "platform_type", "transfer_cnt",
		//"group_id", "group_id_str",
		"image",
		//"related_video_thumb", "is_video",
		"sh_price", "official_price",
		"car_mileage", "car_age",
		// "sub_title", "car_source_type", "authentication_method",
		"tags", "tags_v2",
		//"official_hint_bar",
		"special_tags",
		// "is_self_trade",
	}
	updatedRecords = append(updatedRecords, header)

	for _, car := range carlist {
		var tagsStr, tagsV2Str, specialTagsStr string
		if car.Tags != nil {
			var res []string
			for _, tag := range car.Tags {
				res = append(res, tag.Text)
			}
			tagsBytes, err := json.Marshal(res)
			if err == nil {
				tagsStr = string(tagsBytes)
			}
		}

		if car.TagsV2 != nil {
			var res []string
			for _, tag := range car.TagsV2 {
				res = append(res, tag.Text)
			}
			tagsBytes, err := json.Marshal(res)
			if err == nil {
				tagsV2Str = string(tagsBytes)
			}
		}

		if car.SpecialTags != nil {
			tagsBytes, err := json.Marshal(car.SpecialTags)
			if err == nil {
				specialTagsStr = string(tagsBytes)
			}
		}

		if car.CarMileage == "" && car.CarAge == "" {
			split := strings.Split(car.SubTitle, "|")
			car.CarAge = split[0]
			car.CarMileage = split[1]
		}

		car.ShPrice = strings.ReplaceAll(car.ShPrice, "万", "")
		car.OfficialPrice = strings.ReplaceAll(car.OfficialPrice, "万", "")
		car.CarMileage = strings.ReplaceAll(car.CarMileage, "万公里", "")

		updatedRecords = append(updatedRecords, []string{
			fmt.Sprintf("%v", car.Title),
			fmt.Sprintf("%v", car.SkuId),
			fmt.Sprintf("%v", car.ShopId),
			fmt.Sprintf("%v", car.CarId),
			fmt.Sprintf("%v", car.CarName),
			//fmt.Sprintf("%v", car.SeriesId),
			fmt.Sprintf("%v", car.SeriesName),
			//fmt.Sprintf("%v", car.BrandId),
			fmt.Sprintf("%v", car.BrandName),
			fmt.Sprintf("%v", car.CarYear),
			fmt.Sprintf("%v", car.CarSourceCityName),
			fmt.Sprintf("%v", car.BrandSourceCityName),
			//fmt.Sprintf("%v", car.SpuId),
			//fmt.Sprintf("%v", car.PlatformType),
			//fmt.Sprintf("%v", car.TransferCnt),
			//fmt.Sprintf("%v", car.GroupId),
			//fmt.Sprintf("%v", car.GroupIdStr),
			fmt.Sprintf("%v", car.Image),
			//fmt.Sprintf("%v", car.RelatedVideoThumb),
			//fmt.Sprintf("%v", car.IsVideo),
			fmt.Sprintf("%v", car.ShPrice),
			fmt.Sprintf("%v", car.OfficialPrice),
			fmt.Sprintf("%v", car.CarMileage),
			fmt.Sprintf("%v", car.CarAge),

			//fmt.Sprintf("%v", car.SubTitle),
			//fmt.Sprintf("%v", car.CarSourceType),
			//fmt.Sprintf("%v", car.AuthenticationMethod),
			tagsStr,   // Serialized Tags
			tagsV2Str, // TagsV2 as string
			//fmt.Sprintf("%v", car.OfficialHintBar),
			specialTagsStr, // SpecialTags as string
			//fmt.Sprintf("%v", car.IsSelfTrade),
		})
	}
	tools.WriteCsv("1.csv", updatedRecords)
}

func TestRead(t *testing.T) {

	idMap := GetCarID("")

	var cars []CarContent
	var count int
	for id, name := range idMap {
		fmt.Printf("output:%v, %v\n", name, id)
		cars = append(cars, *GetCarInfo(id))
		time.Sleep(time.Second * 3)
		count++

		//if count == 3 {
		//	break
		//}
	}

	marshal, err := json.Marshal(cars)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	// 47556
	os.WriteFile("car.json", marshal, 0600)
}

func TestRead2(t *testing.T) {
	file, err := os.ReadFile("car.json")
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return

	}

	var cars []CarContent

	err = json.Unmarshal(file, &cars)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	var updatedRecords [][]string

	// 添加新列的标题
	header := []string{
		"car_name",
		"car_id",
		"series_name",
		"brand_name",
		"car_year",
	}

	groupT := getBasicInfo()
	groupT = append(groupT, getBody()...)
	groupT = append(groupT, getEngine()...)
	groupT = append(groupT, getTransmission()...)
	groupT = append(groupT, getChassis()...)

	groupT = append(groupT, []string{
		"车轮/制动",
		"主动安全",
		"被动安全",
		"座椅配置",
		"舒适/防盗配置",
		"辅助/操控配置",
		"内部配置",
		"外部配置",
		"玻璃/后视镜",
		//"灯光配置",
		//"影音娱乐",
		"智能互联",
		"智能化配置",
		"空调/冰箱",
		"选装包",
	}...)

	for i, car := range cars {
		elems := []string{
			fmt.Sprintf("%v", car.CarName),
			fmt.Sprintf("%v", car.CarId),
			//fmt.Sprintf("%v", car.SeriesId),
			fmt.Sprintf("%v", car.SeriesName),
			//fmt.Sprintf("%v", car.BrandId),
			fmt.Sprintf("%v", car.BrandName),
			fmt.Sprintf("%v", car.CarYear),
		}

		for _, groupTitle := range groupT {
			if getNoShow(groupTitle) {
				continue
			}

			if i == 0 {
				header = append(header, groupTitle)
			}

			v, ok := car.Tags[groupTitle]
			if ok {
				elems = append(elems, strings.Join(v, ","))
			} else {
				fmt.Printf("没有: %v %v %v\n", car.CarName, groupTitle, car.CarId)
				elems = append(elems, "")
			}
		}

		if i == 0 {
			updatedRecords = append(updatedRecords, header)
		}

		updatedRecords = append(updatedRecords, elems)
	}

	tools.WriteCsv("car.csv", updatedRecords)
}

func TestJ(t *testing.T) {
	str := `[{"text":"直营","logo":"https://p3.dcarimg.com/img/tos-cn-i-dcdx/b0f80078812d4c138709e3183ccd2994~tplv-dcdx-origin.image","text_color":"","background_color":""},{"text":"送整备保养","logo":"","text_color":"rgba(209,135,0,1)","background_color":"transparent"},{"text":"收藏飙升","logo":"","text_color":"rgba(209,135,0,1)","background_color":"transparent"},{"text":"深度质检","logo":"","text_color":"rgba(118,122,138,1)","background_color":"transparent"},{"text":"高配","logo":"","text_color":"rgba(118,122,138,1)","background_color":"transparent"}]`

	var tt []Tag

	json.Unmarshal([]byte(str), &tt)

	fmt.Printf("output: %v\n", tt)
}

func GetCarInfoBak() CarContent {
	//reader := GetCarDetailHtml("3456")
	//os.WriteFile("2.html", reader, 0600)
	//return

	htmlFilePath := "/Users/bytedance/Documents/project/go/gendev/biz/handler/car/2.html"

	reader, err := os.ReadFile(htmlFilePath)
	if err != nil {
	}

	reader2 := bytes.NewReader(reader)

	html, err := NewCarHtml(reader2)
	if err != nil {
	}

	content := html.getCarGroupContent()
	for k, v := range content {
		fmt.Printf("output: %v: %v\n", k, v)
	}

	var car CarContent
	car.CarInfo = *html.getCarInfo()
	car.Tags = content

	return car
}
