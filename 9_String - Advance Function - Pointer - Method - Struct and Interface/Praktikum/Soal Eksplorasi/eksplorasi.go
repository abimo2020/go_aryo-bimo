package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func chiper(name string) string {
	result := []rune{}
	for _, value := range name {
		value = 'a' + ('z' - value)
		result = append(result, value)
	}
	return string(result)
}

func (s *student) Encode() string {
	var nameEncode = ""
	nameEncode = chiper(s.name)
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	if s.nameEncode == "" {
		s.nameEncode = chiper(s.name)
	}
	nameDecode = chiper(s.nameEncode)
	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
	}
}
