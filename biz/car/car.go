package car

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"Guoaking/gotools/tools"

	"github.com/PuerkitoBio/goquery"
)

/**
@description
@date: 05/31 19:51
@author Gk
**/

func (params RequestParams) GetUrlParam() string {
	formValues := url.Values{}
	if params.Price != "" {
		formValues.Set("price", params.Price)
	}

	if params.SeriesType != 0 {
		formValues.Set("series_type", strconv.Itoa(params.SeriesType))
	}

	if params.MileageRange != "" {
		formValues.Set("mileage_range", params.MileageRange)
	}
	if params.FuelForm != "" {
		formValues.Set("fuel_form", params.FuelForm)
	}
	if params.CapacityL != "" {
		formValues.Set("capacity_l", params.CapacityL)
	}
	if params.GearBoxType != "" {
		formValues.Set("gearbox_type", params.GearBoxType)
	}

	if params.AgeRange != "" {
		formValues.Set("age_range", params.AgeRange)
	}
	if params.ExpandedOrigin != 0 {
		formValues.Set("expanded_origin", strconv.Itoa(params.ExpandedOrigin))
	}
	if params.SHCityName != "" {
		formValues.Set("sh_city_name", params.SHCityName)
	}
	if params.Page != 0 {
		formValues.Set("page", strconv.Itoa(params.Page))
	}
	if params.Limit != 0 {
		formValues.Set("limit", strconv.Itoa(params.Limit))
	}
	return formValues.Encode()
}

func GetCarList(params RequestParams) *CarList {

	requestBodyString := params.GetUrlParam()
	var data io.Reader = strings.NewReader(requestBodyString)

	url := fmt.Sprintf("https://www.dongchedi.com/motor/pc/sh/sh_sku_list?aid=1839&app_name=auto_web_pc")
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Printf("output: %v\n", err)
		return nil
	}
	do, err := BaseDo[CarList](req)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return nil
	}
	return &do
}

func GetCarDetailHtml(carID string) []byte {
	//req, err := http.NewRequest("GET", "https://www.dongchedi.com/auto/params-carIds-3456", nil)
	client := &http.Client{}
	url := fmt.Sprintf("https://www.dongchedi.com/auto/params-carIds-%v", carID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	//req.Header.Set("cache-control", "no-cache")
	//req.Header.Set("dnt", "1")
	req.Header.Set("get-svc", "1")
	//req.Header.Set("pragma", "no-cache")
	//req.Header.Set("priority", "u=0, i")
	//req.Header.Set("referer", "https://www.dongchedi.com/usedcar/19417740")
	req.Header.Set("sec-ch-ua", `"Chromium";v="137", "Not/A)Brand";v="24"`)
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")
	req.Header.Set("cookie", "ttwid=1%7C4s8wbe5Ujrl7rBlFs7vdwa8JMcorFOFRGuKZljxDvKc%7C1747103408%7C054046f329601a740354b1189544e09d8e72d47b50c86c97e5988058aef8ce0e; tt_webid=7503751884554651161; tt_web_version=new; is_dev=false; is_boe=false; x-web-secsdk-uid=4467f8d9-9ddb-4e44-978b-a0365e2c255d; s_v_web_id=verify_malw9f0s_NnLgf7Hj_9sQ4_4Bgp_A9LU_cbAVBRLHl717; city_name=%E5%8C%97%E4%BA%AC; biz_trace_id=7fd8742c; rit_city=%E5%8C%97%E4%BA%AC")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return all
}

func BaseDo[T any](req *http.Request) (T, error) {
	var res T

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("origin", "https://www.dongchedi.com")
	req.Header.Set("referer", "https://www.dongchedi.com/usedcar/3,5-1-x-x-x-x-x-x-x-x-x-x-x-x-x-x-3,5-x-3-x-x-1-1-x-x-x-x-x")
	req.Header.Set("cookie", "ttwid=1%7C4s8wbe5Ujrl7rBlFs7vdwa8JMcorFOFRGuKZljxDvKc%7C1747103408%7C054046f329601a740354b1189544e09d8e72d47b50c86c97e5988058aef8ce0e; tt_webid=7503751884554651161; tt_web_version=new; is_dev=false; is_boe=false; x-web-secsdk-uid=4467f8d9-9ddb-4e44-978b-a0365e2c255d; s_v_web_id=verify_malw9f0s_NnLgf7Hj_9sQ4_4Bgp_A9LU_cbAVBRLHl717; city_name=%E5%8C%97%E4%BA%AC; biz_trace_id=7fd8742c")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, nil
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	var resp2 BaseVo[T]
	err = json.Unmarshal(bodyText, &resp2)
	if err != nil {
		return res, err
	}

	if resp2.Status != 0 {
		if strings.Contains(resp2.Message, "beops_session cookie is expired or invalid") {
			return res, fmt.Errorf("需要更新一下BeopsSession再试: err:%s ", resp2.Message)
		}
		return res, fmt.Errorf("%s", resp2.Message)
	}

	return resp2.Data, nil

}

func PageSp(Total, limit int, f func(page, limit int)) {

	var old, end int
	if Total <= limit {
		// 无须分页 直接用
		f(1, limit)
		return
	}

	count := Total / limit
	for i := 1; i < count+2; i++ {
		if i == 1 {
			continue
		}
		old = old + limit
		end = limit + old
		if end >= Total {
			end = Total
		}

		fmt.Printf("start:第%v页, %v:%v\n", i, old, end)
		f(i, limit)
		time.Sleep(time.Second * 3)
	}

}

func getC(page, limit int) *CarList {
	return GetCarList(RequestParams{
		Price:          "3,6",
		SeriesType:     1,
		AgeRange:       "2,6",
		ExpandedOrigin: 2,
		SHCityName:     "全国",
		MileageRange:   "0,6",
		FuelForm:       "1",
		CapacityL:      "1.5,2.0",
		GearBoxType:    "2",
		Page:           page,
		Limit:          limit,
	})
}

func getData(filename string) []SearchShSkuInfo {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var carlist []SearchShSkuInfo
	err = json.Unmarshal(b, &carlist)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return nil
	}
	return carlist
}

func GetCarID(sname string) map[string]string {
	csv, err := tools.ReadCsv("1.csv")
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return nil
	}

	idMap := make(map[string]string)
	for _, record := range csv[1:] {

		carId := record[3]
		carName := record[4]
		series_name := record[5]

		if sname != "" {
			if series_name != sname {
				continue
			}
		}

		v, ok := idMap[carId]
		if !ok {
			idMap[carId] = carName
		} else {
			if v != carName {
				fmt.Printf("output: 名字不一样的: %v, %v\n", v, carName)
			}
		}

	}

	return idMap
}

func GetCarInfo(carID string) *CarContent {
	reader := GetCarDetailHtml(carID)

	var car CarContent
	html, err := NewCarHtml(bytes.NewReader(reader))
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return nil
	}

	car.CarInfo = *html.getCarInfo()
	car.Tags = html.getCarGroupContent()
	return &car
}

func NewCarHtml(htmlContent io.Reader) (CarHtml, error) {
	htmlDoc, err := goquery.NewDocumentFromReader(htmlContent)
	BasicBody := htmlDoc.Find("[class^='configuration_root__']").First()
	return CarHtml{
		htmlDoc:   htmlDoc,
		BasicBody: BasicBody,
	}, err
}

func (h *CarHtml) getCarInfo() *CarInfo {

	if h.car != nil {
		return h.car
	}

	scriptContent := h.htmlDoc.Find("script#__NEXT_DATA__").Text()
	if scriptContent == "" {
		log.Fatal("未找到__NEXT_DATA__脚本")
	}

	var nextData Top
	if err := json.Unmarshal([]byte(scriptContent), &nextData); err != nil {
		log.Fatalf("JSON解析失败: %v", err)
	}

	carInfo := nextData.Props.PageProps.RawData.CarInfo
	if len(carInfo) == 1 {
		h.car = &carInfo[0]
		return h.car
	}

	return nil
}

func (h *CarHtml) getBasicBody() {
	h.BasicBody = h.htmlDoc.Find("[class^='configuration_root__']").First()
}

func (h *CarHtml) getCarTitle() string {
	return h.BasicBody.Find("[class^='table_head__'],[class^='cell_car__']").Text()
}

func (h *CarHtml) getCarGroupContent() map[string][]string {
	navM := make(map[string][]string)

	//navM2 := make(map[string]CarContent)
	// 车身 -> [aaaa]
	// 车窗 -> [玻璃: 有的, 座椅: 前 , 后 ]
	h.BasicBody.Find("[name^='config-body']").Each(func(i int, group *goquery.Selection) {
		groupTitle := getGroupTitle(group)

		group.Find("[data-row-anchor]").Each(func(j int, row *goquery.Selection) {
			value := row.Find("[class^='cell_normal']").Text()
			label, value := getLabel(row), getValue(row)

			if isBasic(groupTitle) {
				navM[label] = append(navM[label], value)
				return
			}

			if value != "" && value != "-" {
				data := label
				if value != "标配" {
					data = fmt.Sprintf("%v:%v", label, value)
				}

				navM[groupTitle] = append(navM[groupTitle], data)
			}
		})
	})

	return navM
}

func getBasicInfo() []string {

	//基本信息</a> → <a>Basic Info</a>
	//车身</a> → <a>Body</a>
	//发动机</a> → <a>Engine</a>
	//变速箱</a> → <a>Transmission</a>
	//底盘/转向</a> → <a>Chassis/Steering</a>
	//车轮/制动</a> → <a>Wheels/Braking</a>
	//主动安全</a> → <a>Active Safety</a>
	//被动安全</a> → <a>Passive Safety</a>
	//辅助/操控配置</a> → <a>Assistance/Control Configuration</a>
	//外部配置</a> → <a>Exterior Configuration</a>
	//>内部配置</a> → <a>Interior Configuration</a>
	//>舒适/防盗配置</a> → <a>Comfort/Anti-theft Configuration</a>
	//>座椅配置</a> → <a>Seat Configuration</a>
	//>智能互联</a> → <a>Smart Connectivity</a>
	//>影音娱乐</a> → <a>Multimedia</a>
	//>灯光配置</a> → <a>Lighting Configuration</a>
	//>玻璃/后视镜</a> → <a>Glass/Mirrors</a>
	//>空调/冰箱</a> → <a>AC/Refrigerator</a>
	//>智能化配置</a> → <a>Smart Features</a>
	//
	//var ActiveSafety, PassiveSafety, AssistConfig, ExterConfig, InterConfig, ComfortConfig, SeatConfig, SmartConfig []string
	//var MultimediaConfig, LightingConfig, MirrorsConfig, ACConfig, SmartFeatures []string

	return []string{"厂商", "级别", "能源类型", "上市时间", "发动机", "最大功率(kW)", "最大扭矩(N·m)", "变速箱", "长x宽x高(mm)", "车身结构", "最高车速(km/h)", "NEDC综合油耗(L/100km)", "整车保修期限", "6万公里保养总成本预估"}
}

func getBody() []string {
	return []string{"长(mm)", "宽(mm)", "高(mm)", "轴距(mm)", "前轮距(mm)", "后轮距(mm)", "最小离地间隙(mm)", "车门数(个)", "车门开启方式", "座位数(个)", "整备质量(kg)", "满载质量(kg)", "油箱容积(L)", "行李舱容积(L)"}
}

func getEngine() []string {
	return []string{"发动机型号", "排量(mL)", "排量(L)", "进气形式", "气缸排列形式", "气缸数(个)", "每缸气门数(个)", "配气机构", "最大马力(Ps)", "最大功率转速(rpm)", "最大扭矩转速(rpm)", "发动机特有技术", "燃料形式", "燃油标号", "供油方式", "缸盖材料", "缸体材料", "环保标准"}
}

func getTransmission() []string {
	return []string{"变速箱描述", "挡位数", "变速箱类型"}
}

func getChassis() []string {
	return []string{"驱动方式", "前悬挂形式", "后悬挂形式", "转向类型", "车体结构"}
}

func getNoShow(key string) bool {
	m := map[string]string{
		"厂商":         "",
		"级别":         "",
		"能源类型":       "",
		"车身结构":       "",
		"整车保修期限":     "",
		"车门数(个)":     "",
		"座位数(个)":     "",
		"前轮距(mm)":    "",
		"后轮距(mm)":    "",
		"最小离地间隙(mm)": "",
		"长x宽x高(mm)":  "",
		"高(mm)":      "",
		"车门开启方式":     "",
		"排量(mL)":     "",
		"挡位数":        "",
		"车体结构":       "",
		"变速箱":        "",
		"灯光配置":       "",
		"影音娱乐":       "",
		"选装包":        "",
	}

	_, ok := m[key]
	return ok
}

func isBasic(groupTitle string) bool {
	if groupTitle == "基本信息" ||
		groupTitle == "车身" ||
		groupTitle == "底盘/转向" ||
		groupTitle == "变速箱" ||
		groupTitle == "发动机" {
		return true
	}
	return false
}

func getValue(row *goquery.Selection) string {
	values := row.Find("[class*='cell_normal'], [class*='value']").Map(func(_ int, el *goquery.Selection) string {
		return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(el.Text(), " ", ""), "\n", ""))
	})
	value := strings.Join(values, "")
	value = strings.ReplaceAll(strings.TrimPrefix(value, "●"), "●", "+")
	return strings.TrimSpace(value)
}

func getLabel(row *goquery.Selection) string {
	label := row.Find("[class^='cell_label__']").Text()

	if strings.Contains(label, "\n") {
		label = strings.Split(label, "\n")[1]
	}

	return strings.TrimSpace(label)
}

func getGroupTitle(row *goquery.Selection) string {
	groupTitle := row.Find("[class*='table_is-title']").Text()
	//if strings.Contains(groupTitle, "影音娱乐") {
	//	fmt.Printf("output: %v\n", 11)
	//}

	groupTitle = strings.ReplaceAll(strings.ReplaceAll(groupTitle, " ", ""), "\n", "")
	groupTitle = strings.Split(groupTitle, "●")[0]

	return strings.TrimSpace(groupTitle)
}
