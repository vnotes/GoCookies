package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"utils"

	"github.com/PuerkitoBio/goquery"
)

var wg sync.WaitGroup

var result []CarLogo

type CarLogo struct {
	Country string
	Brand   []CarBrand
}

type CarBrand struct {
	Name    string
	Url     string
	Comment string
}

func FetchInfo(car *utils.Request) (dataJson CarLogo) {
	body, err := car.RequestReader()
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}
	var carLogo []CarBrand
	country := doc.Find(".mark a").Text()
	country = utils.Source2DestCode(country, "gbk", "utf8")
	doc.Find(".expPicA li").Each(func(i int, selection *goquery.Selection) {
		img, _ := selection.Find(".dPic .iPic a img").Attr("src")
		img = "https:" + img
		name := selection.Find(".dTxt .iTit a").Text()
		name = utils.Source2DestCode(name, "gbk", "utf8")
		comment := selection.Find(".dTxt .iDes").Text()
		comment = utils.Source2DestCode(comment, "gbk", "utf8")
		data := CarBrand{name, img, comment}
		carLogo = append(carLogo, data)
	})
	dataJson = CarLogo{country, carLogo}
	return

}

func FetchInfoV2(car *utils.Request) {
	defer wg.Done()
	body, err := car.RequestReader()
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}
	var carLogo []CarBrand
	country := doc.Find(".mark a").Text()
	country = utils.Source2DestCode(country, "gbk", "utf8")
	doc.Find(".expPicA li").Each(func(i int, selection *goquery.Selection) {
		img, _ := selection.Find(".dPic .iPic a img").Attr("src")
		img = "https:" + img
		name := selection.Find(".dTxt .iTit a").Text()
		name = utils.Source2DestCode(name, "gbk", "utf8")
		comment := selection.Find(".dTxt .iDes").Text()
		comment = utils.Source2DestCode(comment, "gbk", "utf8")
		data := CarBrand{name, img, comment}
		carLogo = append(carLogo, data)
	})
	result = append(result, CarLogo{country, carLogo})
}

func ChebiaoMain() {
	fileName := "car.json"
	jsonFile, _ := os.Create(fileName)
	defer jsonFile.Close()
	_request := &utils.Request{
		Method:  "GET",
		Url:     "https://www.pcauto.com.cn/zt/chebiao/",
		Headers: map[string]string{"": ""},
		Body:    nil,
	}
	body, err := _request.RequestReader()
	if err != nil {
		fmt.Printf("request failed, error is %s ", err.Error())
		return
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Printf("request to document failed, error is %s ", err.Error())
		return
	}
	doc.Find("#menu ul li ").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Find("a").Attr("href")
		car := &utils.Request{
			Method:  "GET",
			Url:     "https:" + url,
			Headers: map[string]string{"": ""},
			Body:    nil,
		}
		rs := FetchInfo(car)
		result = append(result, rs)

	})
	data, _ := json.MarshalIndent(result, "", "  ")
	_, err = jsonFile.Write(data)
	if err != nil {
		fmt.Printf("wrie json file failed: %s", err.Error())
		return
	}
}

func ChebiaoMainV2() {
	fileName := "car.json"
	jsonFile, _ := os.Create(fileName)
	defer jsonFile.Close()
	_request := &utils.Request{
		Method:  "GET",
		Url:     "https://www.pcauto.com.cn/zt/chebiao/",
		Headers: map[string]string{"": ""},
		Body:    nil,
	}
	body, err := _request.RequestReader()
	if err != nil {
		fmt.Printf("request failed, error is %s ", err.Error())
		return
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Printf("request to document failed, error is %s ", err.Error())
		return
	}

	doc.Find("#menu ul li ").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Find("a").Attr("href")
		car := &utils.Request{
			Method:  "GET",
			Url:     "https:" + url,
			Headers: map[string]string{"": ""},
			Body:    nil,
		}
		wg.Add(1)
		go FetchInfoV2(car)
	})
	wg.Wait()
	data, _ := json.MarshalIndent(result, "", "  ")
	_, err = jsonFile.Write(data)
	if err != nil {
		fmt.Printf("wrie json file failed: %s", err.Error())
		return
	}
}

func main() {
	ChebiaoMain()
	ChebiaoMainV2()
}
