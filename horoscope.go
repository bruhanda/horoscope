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

func GetHoroscope()  map[string]map[string]Horoscope{
	general := make(map[string]map[string]Horoscope)
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
		//fmt.Println(content)
		horoscope := Zsigns{}

		err = xml.Unmarshal([]byte(content), &horoscope)
		if err != nil {
			log.Println("error: %v", err)
		}



		for sgn, _ := range Signs {
			switch sgn {
			case "aries":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "taurus":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "gemini":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "cancer":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "leo":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "virgo":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "libra":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "scorpio":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "sagittarius":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "capricorn":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "aquarius":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}
			case "pisces":
				general[key]=map[string]Horoscope{
					"aries":horoscope.Aries,
				}

			}
		}

	}
    return general
}

