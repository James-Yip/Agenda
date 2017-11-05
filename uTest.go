package main

import fuck "agenda/util"
import "fmt"
func main() {
	var fucjk string
	fucjk =  fuck.Time2str("2017-10-31/23:23")
	fucjk = fuck.Str2time(fucjk)
	x := fuck.IsTimeValid(fucjk)
	fmt.Println(x)
	fmt.Println(fucjk)
}