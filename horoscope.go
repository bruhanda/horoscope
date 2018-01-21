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

type Zsigns struct {
	Aries       Horoscope `xml:"aries"`
	Taurus      Horoscope `xml:"taurus"`
	Gemini      Horoscope `xml:"gemini"`
	Сancer      Horoscope `xml:"cancer"`
	Leo         Horoscope `xml:"leo"`
	Virgo       Horoscope `xml:"virgo"`
	Libra       Horoscope `xml:"libra"`
	Scorpio     Horoscope `xml:"scorpio"`
	Sagittarius Horoscope `xml:"sagittarius"`
	Capricorn   Horoscope `xml:"capricorn"`
	Aquarius    Horoscope `xml:"aquarius"`
	Pisces      Horoscope `xml:"pisces"`
}

type Horoscope struct {
	Yesterday  string `xml:"yesterday"`
	Today      string `xml:"today"`
	Tomorrow   string `xml:"tomorrow"`
	Tomorrow02 string `xml:"tomorrow02"`
}

func GetHoroscope() map[string]map[string]Horoscope {
	general := map[string]map[string]Horoscope{
		"com": map[string]Horoscope{
			"aries":       Horoscope{},
			"taurus":      Horoscope{},
			"gemini":      Horoscope{},
			"cancer":      Horoscope{},
			"leo":         Horoscope{},
			"virgo":       Horoscope{},
			"libra":       Horoscope{},
			"scorpio":     Horoscope{},
			"sagittarius": Horoscope{},
			"capricorn":   Horoscope{},
			"aquarius":    Horoscope{},
			"pisces":      Horoscope{},
		},
		"ero": map[string]Horoscope{
			"aries":       Horoscope{},
			"taurus":      Horoscope{},
			"gemini":      Horoscope{},
			"cancer":      Horoscope{},
			"leo":         Horoscope{},
			"virgo":       Horoscope{},
			"libra":       Horoscope{},
			"scorpio":     Horoscope{},
			"sagittarius": Horoscope{},
			"capricorn":   Horoscope{},
			"aquarius":    Horoscope{},
			"pisces":      Horoscope{},
		},
		"anti": map[string]Horoscope{
			"aries":       Horoscope{},
			"taurus":      Horoscope{},
			"gemini":      Horoscope{},
			"cancer":      Horoscope{},
			"leo":         Horoscope{},
			"virgo":       Horoscope{},
			"libra":       Horoscope{},
			"scorpio":     Horoscope{},
			"sagittarius": Horoscope{},
			"capricorn":   Horoscope{},
			"aquarius":    Horoscope{},
			"pisces":      Horoscope{},
		},
		"bus": map[string]Horoscope{
			"aries":       Horoscope{},
			"taurus":      Horoscope{},
			"gemini":      Horoscope{},
			"cancer":      Horoscope{},
			"leo":         Horoscope{},
			"virgo":       Horoscope{},
			"libra":       Horoscope{},
			"scorpio":     Horoscope{},
			"sagittarius": Horoscope{},
			"capricorn":   Horoscope{},
			"aquarius":    Horoscope{},
			"pisces":      Horoscope{},
		},
		"hea": map[string]Horoscope{
			"aries":       Horoscope{},
			"taurus":      Horoscope{},
			"gemini":      Horoscope{},
			"cancer":      Horoscope{},
			"leo":         Horoscope{},
			"virgo":       Horoscope{},
			"libra":       Horoscope{},
			"scorpio":     Horoscope{},
			"sagittarius": Horoscope{},
			"capricorn":   Horoscope{},
			"aquarius":    Horoscope{},
			"pisces":      Horoscope{},
		},
		"cook": map[string]Horoscope{
			"aries":       Horoscope{},
			"taurus":      Horoscope{},
			"gemini":      Horoscope{},
			"cancer":      Horoscope{},
			"leo":         Horoscope{},
			"virgo":       Horoscope{},
			"libra":       Horoscope{},
			"scorpio":     Horoscope{},
			"sagittarius": Horoscope{},
			"capricorn":   Horoscope{},
			"aquarius":    Horoscope{},
			"pisces":      Horoscope{},
		},
		"mob": map[string]Horoscope{
			"aries":       Horoscope{},
			"taurus":      Horoscope{},
			"gemini":      Horoscope{},
			"cancer":      Horoscope{},
			"leo":         Horoscope{},
			"virgo":       Horoscope{},
			"libra":       Horoscope{},
			"scorpio":     Horoscope{},
			"sagittarius": Horoscope{},
			"capricorn":   Horoscope{},
			"aquarius":    Horoscope{},
			"pisces":      Horoscope{},
		},
		"lov": map[string]Horoscope{
			"aries":       Horoscope{},
			"taurus":      Horoscope{},
			"gemini":      Horoscope{},
			"cancer":      Horoscope{},
			"leo":         Horoscope{},
			"virgo":       Horoscope{},
			"libra":       Horoscope{},
			"scorpio":     Horoscope{},
			"sagittarius": Horoscope{},
			"capricorn":   Horoscope{},
			"aquarius":    Horoscope{},
			"pisces":      Horoscope{},
		},
	}

	for key, _ := range Keys {
		url := fmt.Sprintf(urlStr, key)
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
		horoscope := Zsigns{}

		err = xml.Unmarshal([]byte(content), &horoscope)
		if err != nil {
			log.Println("error: %v", err)
		}

		general[key]["aries"] = horoscope.Aries
		general[key]["taurus"] =horoscope.Taurus
		general[key]["gemini"] = horoscope.Gemini
		general[key]["cancer"] = horoscope.Сancer
		general[key]["leo"] =horoscope.Leo
		general[key]["virgo"] = horoscope.Virgo
		general[key]["libra"] = horoscope.Libra
		general[key]["scorpio"] = horoscope.Scorpio
		general[key]["sagittarius"] = horoscope.Sagittarius
		general[key]["capricorn"] = horoscope.Capricorn
		general[key]["aquarius"] = horoscope.Aquarius
		general[key]["pisces"] = horoscope.Pisces

	}

	return general
}

