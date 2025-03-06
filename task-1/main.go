package main

import "fmt"

func main() {
	const formulaMessage = "Formula used: BMI = weight (kg) / (height (m) * height (m))"
	var name string
	var weight float64
	var height float64

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)

	fmt.Print("Enter your weight in kg: ")
	fmt.Scanf("%f", &weight)

	fmt.Print("Enter your height in meters: ")
	fmt.Scanf("%f", &height)

	fmt.Printf("Welcome, %s! Let's calculate your BMI.\n", name)
	fmt.Println(formulaMessage)

	bmi := calculateBMI(weight, height)
	fmt.Printf("Your BMI is: %.2f\n", bmi)

	fmt.Println("Health Status:", getHealthStatus(bmi))
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

func getHealthStatus(bmi float64) string {
	if bmi < 18.5 {
		return "Underweight"
	} else if bmi >= 18.5 && bmi <= 24.9 {
		return "Normal weight"
	} else if bmi >= 25 && bmi <= 29.9 {
		return "Overweight"
	} else {
		return "Obese"
	}
}
