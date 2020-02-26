package provider

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/trrtly/proxy/request"
)

// Horo 蜻蜓代理
type Horo struct {
	apiURL string
}

// HoroOpts 蜻蜓代理配置
type HoroOpts struct {
	OrderID       string
	Num           string
	Format        string
	LineSeparator string
	CanRepeat     string
	UserToken     string
}

// HoroIPResp HoroIpResp
type HoroIPResp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data []*HoroIPData `json:"data"`
}

// HoroIPData HoroIpData
type HoroIPData struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	CountryCn  string `json:"country_cn"`
	ProvinceCn string `json:"province_cn"`
	CityCn     string `json:"city_cn"`
}

// HoroURL HoroURL
const HoroURL = "https://proxyapi.horocn.com/api/v2/proxies"

// NewHoro init
func NewHoro(opts *HoroOpts) *Horo {
	horoIns := new(Horo)
	horoIns.apiURL = horoIns.generalAPIURL(opts)
	return horoIns
}

// GetProxys 获取代理
func (h *Horo) GetProxys() (res []string, err error) {
	response, err := request.DefaultGet.Request(h.apiURL)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	horoResp := HoroIPResp{}
	err = json.Unmarshal(body, &horoResp)
	if err != nil {
		return
	}
	for _, v := range horoResp.Data {
		res = append(res, v.Host + ":", v.Port)
	}
	return
}

func (h *Horo) generalAPIURL(opts *HoroOpts) string {
	// Query params
	params := url.Values{}
	params.Add("order_id", opts.OrderID)
	params.Add("num", opts.Num)
	params.Add("format", opts.Format)
	params.Add("line_separator", opts.LineSeparator)
	params.Add("can_repeat", opts.CanRepeat)
	params.Add("user_token", opts.UserToken)

	return HoroURL + "?" + params.Encode()
}
