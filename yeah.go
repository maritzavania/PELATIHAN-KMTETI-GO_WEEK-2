package main

import (
	"errors"
	"fmt"
)

func cekUmur(umur int) (string, error) {
	if umur < 18 {
		return "", errors.New("umur tidak mencukupi")
	}
	return "Umur cukup", nil
}
func main() {
	if hasil, err := cekUmur(16); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(hasil)
	}
}
