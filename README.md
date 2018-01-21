# horoscope
golang русский гороскоп; Horoscope package

```sh
package main

import (
	"fmt"
	"github.com/bruhanda/horoscope"
)

func main()  {
	
	res:=horoscope.GetHoroscope()

	for key, name:=range horoscope.Keys{
		fmt.Println(name)
		for signKey, signName:=range horoscope.Signs{
			fmt.Println(signName)
			fmt.Println(res[key][signKey].Today)
			fmt.Println("------------------------------------")
		}
	}

}
```
