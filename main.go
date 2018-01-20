package horoscope

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/xml"
)

var urlStr string = "http://ignio.com/r/export/utf/xml/daily/%s.xml"

var Keys = map[string]string{
	"com":  "Общий",       //общий
	"ero":  "Эротический", //эротический
	"anti": "Анти",        //анти
	"bus":  "Бизнес",      //бизнес
	"hea":  "Здоровья",    //здоровья
	"cook": "Кулинарный",  //кулинарный
	"mob":  "Мобильный",   //мобильный
	"lov":  "Любовный",    //любовный
}

var Signs = map[string]string{
	"aries":       "Овен",
	"taurus":      "Телец",
	"gemini":      "Близнецы",
	"cancer":      "Рак",
	"leo":         "Лев",
	"virgo":       "Дева",
	"libra":       "Весы",
	"scorpio":     "Скорпион",
	"sagittarius": "Стрелец",
	"capricorn":   "Козерог",
	"aquarius":    "Водолей",
	"pisces":      "Рыбы",
}

type Horoscope struct {
	Aries       string `xml:"aries>today"`
	Taurus      string `xml:"taurus>today"`
	Gemini      string `xml:"gemini>today"`
	Сancer      string `xml:"cancer>today"`
	Leo         string `xml:"leo>today"`
	Virgo       string `xml:"virgo>today"`
	Libra       string `xml:"libra>today"`
	Scorpio     string `xml:"scorpio>today"`
	Sagittarius string `xml:"sagittarius>today"`
	Capricorn   string `xml:"capricorn>today"`
	Aquarius    string `xml:"aquarius>today"`
	Pisces      string `xml:"pisces>today"`
}

func GetHoroscope() (map[string]Horoscope) {
	general:=make(map[string]Horoscope)
	for key, _ := range Keys {
		url:= fmt.Sprintf(urlStr, key)
		resp, err := http.Get(url);
		if err != nil {
			log.Println("Can not read body")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Can not read body")
		}
		content := string(body)
		//fmt.Println(content)
		horoscope := Horoscope{}

		err = xml.Unmarshal([]byte(content), &horoscope)
		if err != nil {
			log.Println("error: %v", err)
		}
		general[key]=horoscope
	}
	return general
}

