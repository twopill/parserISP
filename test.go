package main

import (
	"fmt"
	//"reflect"
	"regexp"
	"strings"
	// "unicode"
	// "unicode/utf8"
	//"github.com/webview/webview"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"log"
)

func main() {
	//var pattern string

	a := app.New()
	w := a.NewWindow("RegExp Programm")

	widget.NewLabel("Hi, transform your shit pattern in a fantastic RegExp")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	content := widget.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
		regExpControll(input.Text)
	}))

	w.SetContent(content)
	w.ShowAndRun()

	//fmt.Println("Enter Your Pattern: ")
	//fmt.Scanln(&pattern)
	//regExpControll(pattern)

}

func regExpControll(pattern string) {
	if len(pattern) <= 5 {
		isValid, _ := regexp.MatchString("([0-9]+[!]+[x])", pattern)
		if isValid {
			testParse(pattern, len(pattern))
		}
		fmt.Println("finish")
	}
}

func testParse(pattern string, lenght int) {
	arrayForLetters := make([]interface{}, 0)
	//newArray := make([]interface{}, len(pattern))
	str := pattern
	splitPattern := strings.Split(str, "")

	for _, v := range splitPattern {
		arrayForLetters = append(arrayForLetters, v) // inserisco le lettere nell'array, con annessi spazzi
	}

	fmt.Println("Non modificato, puro: ", arrayForLetters)

	/*for i := 0; i <= len(arrayForLetters); i++ {
		newArray[i] = append(arrayForLetters, i)
	}
	fmt.Println("Nuovo array: ", newArray)*/

	/* for i, v := range arrayForLetters {
		fmt.Println(i, v)
		/*if i%2 == 0 {
		newArray[i] = append(arrayForLetters, v) // inserisco le lettere nell'array, con annessi spazzi
		}
	 } */

	// fmt.Println("Nuovo array: ", newArray)

	/* for _, v := range arrayForLetters {
		if v == reflect.Int {
			theNumber = v.(int)
			print(theNumber)
		} else if v == hasSymbol("!") {
			//exclamationPoint = v.(string)
		}
	} */
}

/*func hasSymbol(str string) bool {
for _, letter := range str {
	if unicode.IsSymbol(letter) {
		return true
	}
}
return false}
*/
