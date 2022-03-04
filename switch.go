package main

import "fmt"

func main() {

	const hr = 3                    // credit hour
	const total_sub = 2             // total subject in a course
	const total_hr = hr * total_sub // total hours
	var sub_1 float64
	var sub1_point float64

	var sub_2 float64
	var sub2_point float64


	// first subject
	fmt.Println("Enter first subject marks:")
	fmt.Scanln(&sub_1)
	if sub_1 >= 0 && sub_1 < 50 {
		fmt.Println("you are fail in first subject..! and point of first subject is ", 0.0*hr)
	} else if sub_1 >= 50 && sub_1 <= 52 {
		fmt.Println("your grade in first subject is D and point of first subject is ", 1.0*hr)
	} else if sub_1 >= 53 && sub_1 <= 56 {
		fmt.Println("your grade in first subject is D+ and point of first subject is ", 1.4*hr)
	} else if sub_1 >= 57 && sub_1 <= 60 {
		fmt.Println("your grade in first subject is C- and point of first subject is ", 1.8*hr)
	} else if sub_1 >= 61 && sub_1 <= 63 {
		fmt.Println("your grade in first subject is C and point of first subject is ", 2.0*hr)
	} else if sub_1 >= 64 && sub_1 <= 67 {
		fmt.Println("your grade in first subject is C+ and point of first subject is ", 2.4*hr)
	} else if sub_1 >= 68 && sub_1 <= 70 {
		fmt.Println("your grade in first subject is B- and point of first subject is ", 2.8*hr)
	} else if sub_1 >= 71 && sub_1 <= 74 {
		fmt.Println("your grade in first subject is B and point of first subject is ", 3.0*hr)
	} else if sub_1 >= 75 && sub_1 <= 79 {
		fmt.Println("your grade in first subject is B+ and point of first subject is ", 3.4*hr)
	} else if sub_1 >= 80 && sub_1 <= 84 {
		fmt.Println("your grade in first subject is A- and point of first subject is ", 3.8*hr)
	} else if sub_1 >= 85 && sub_1 <= 89 {
		fmt.Println("your grade in first subject is A and point of first subject is ", 4.0*hr)
	} else if sub_1 >= 90 && sub_1 <= 100 {
		fmt.Println("your grade in first subject is A+ and point of first subject is ", 4.0*hr)
	}




	// second subject
	fmt.Println("Enter second subject marks:")
	fmt.Scanln(&sub_2)
	if sub_2 >= 0 && sub_2 < 50 {
		fmt.Println("you are fail in second subject..! and point of first subject is ", 0.0*hr)
	} else if sub_2 >= 50 && sub_2 <= 52 {
		fmt.Println("your grade in second subject is D and point of first subject is ", 1.0*hr)
	} else if sub_2 >= 53 && sub_2 <= 56 {
		fmt.Println("your grade in second subject is D+ and point of first subject is ", 1.4*hr)
	} else if sub_2 >= 57 && sub_2 <= 60 {
		fmt.Println("your grade in second subject is C- and point of first subject is ", 1.8*hr)
	} else if sub_2 >= 61 && sub_2 <= 63 {
		fmt.Println("your grade in second subject is C and point of first subject is ", 2.0*hr)
	} else if sub_2 >= 64 && sub_2 <= 67 {
		fmt.Println("your grade in second subject is C+ and point of first subject is ", 2.4*hr)
	} else if sub_2 >= 68 && sub_2 <= 70 {
		fmt.Println("your grade in second subject is B- and point of first subject is ", 2.8*hr)
	} else if sub_2 >= 71 && sub_2 <= 74 {
		fmt.Println("your grade in second subject is B and point of first subject is ", 3.0*hr)
	} else if sub_2 >= 75 && sub_2 <= 79 {
		fmt.Println("your grade in second subject is B+ and point of first subject is ", 3.4*hr)
	} else if sub_2 >= 80 && sub_2 <= 84 {
		fmt.Println("your grade in second subject is A- and point of first subject is ", 3.8*hr)
	} else if sub_2 >= 85 && sub_2 <= 89 {
		fmt.Println("your grade in second subject is A and point of first subject is ", 4.0*hr)
	} else if sub_2 >= 90 && sub_2 <= 100 {
		fmt.Println("your grade in second subject is A+ and point of first subject is ", 4.0*hr)
	}

	var sum_ofGP float64 = sub1_point + sub2_point
	var cgpa float64 = sum_ofGP/total_hr
	fmt.Println("Your CGPA is:",cgpa)
}
