package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func FormatPhone(phone string) (string, error) {
	ethioPhone := `^(09|9|\+2519|2519)[0-9]{8}$`
	re := regexp.MustCompile(ethioPhone)

	fmt.Println(re.MatchString(phone))

	if !re.MatchString(phone) {
		return "", errors.New("improper format of phone")
	} else if strings.HasPrefix(phone, "+251") {
		phone = strings.Replace(phone, "+251", "251", -1)
	} else if strings.HasPrefix(phone, "0") {
		phone = strings.Replace(phone, "0", "251", -1)
	} else if strings.HasPrefix(phone, "9") {
		phone = strings.Replace(phone, "9", "2519", -1)
	} else {
		return phone, nil
	}

	return phone, nil
}

func main() {
	phone1 := "+251912345678"
	phone2 := "251912345678"
	phone3 := "0912345678"
	phone4 := "912345678"

	val1, _ := FormatPhone(phone1)
	val2, _ := FormatPhone(phone2)
	val3, _ := FormatPhone(phone3)
	val4, _ := FormatPhone(phone4)

	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(val3)
	fmt.Println(val4)

}
