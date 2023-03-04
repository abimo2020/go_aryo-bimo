package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name  [5]string
	score [5]int
}

func (s Student) averageScore() int {
	var result int
	for _, value := range s.score {
		result += value
	}
	result = result / len(s.score)
	return result
}

func (s Student) minMaxScore() (minValue, maxValue int, nameMin, nameMax string) {
	minValue = s.score[0]
	maxValue = s.score[0]
	for key, value := range s.score {
		if value < minValue {
			minValue = value
			nameMin = s.name[key]
		}
		if value > maxValue {
			maxValue = value
			nameMax = s.name[key]
		}
	}
	return
}

func main() {
	students := new(Student)
	fmt.Println("Sample Test Case \nInput:")

	/* Inisiasi Student melalui scanner */
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		fmt.Print("Input ", i+1, " Student's Name ")
		scanner.Scan()
		students.name[i] = scanner.Text()
		fmt.Print("Input ", i+1, " Student's Score ")
		scanner.Scan()
		value, _ := strconv.Atoi(scanner.Text())
		students.score[i] = value
	}

	/* Inisiasi Student langsung */
	// students.name = [5]string{"Rizky","Kobar","Ismail","Umam","Yopan"}
	// students.score = [5]int{80, 75, 70, 75, 60}

	avg := students.averageScore()
	minValue, maxValue, nameMin, nameMax := students.minMaxScore()
	fmt.Println("Output:")
	fmt.Println("Average Score :", avg)
	fmt.Printf("Min Score of Students : %s (%d)\n", nameMin, minValue)
	fmt.Printf("Max Score of Students : %s (%d)\n", nameMax, maxValue)
}
