package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// forever := make(chan bool)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Height (cm): ")
	scanner.Scan()
	input := scanner.Text()
	inputHeight, err := strconv.ParseFloat(input, 32)
	if err != nil {
		fmt.Println("Please input your HEIGHT")
		return
	}

	fmt.Print("Weight (kg): ")
	scanner.Scan()
	output := scanner.Text()
	inputWeight, err := strconv.ParseFloat(output, 32)
	if err != nil {
		fmt.Println("Please input your WEIGHT")
		return
	}

	result := bmiCalc(inputWeight, inputHeight)
	fmt.Printf("You are %s\n", result)
	// for scanner.Scan() {
	// 	fmt.Print("Enter your weight (cm,kg) : ")
	// }
	// input
	// fmt.Printf("Result %q", scanner.Text())
	// <-forever
}

func bmiCalc(weight, height float64) string {
	bmi := weight / ((height * height) / 10000)

	fmt.Println(bmi)

	if bmi < 18.5 {
		return "Underweight"
	} else if (bmi >= 18.5) && (bmi <= 22.9) {
		return "Normal"
	} else if (bmi > 23) && (bmi <= 29.9) {
		return "Overweight"
	} else {
		return "Obesity"
	}
}
