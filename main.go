/**
 * @author Bianca Boo
 * @version 1.0.0
 * @date 2025-11-16
 * @fileoverview This program will determine and output age cateogory.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Declare variables
	var studentName string
	var age int
	var isStudentStr string
	var isStudent bool

	// Initialize variables
	fmt.Print("Enter your name: ")
	fmt.Scanln(&studentName) // NOTE: spaces in the name may cause issues

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Are you a student? (true/false): ")
	fmt.Scanln(&isStudentStr)
	isStudent = strings.ToLower(isStudentStr) == "true"

	// Output
	if isStudent {
		if age >= 13 && age <= 19 {
			fmt.Printf("Student %s is a teenager.\n", studentName)
		} else if age >= 5 && age <= 12 {
			fmt.Printf("Student %s is a child.\n", studentName)
		} else {
			fmt.Printf("Student %s is in a different life stage.\n", studentName)
		}
	} else {
		if age >= 20 && age <= 30 {
			fmt.Printf("%s is a young adult.\n", studentName)
		} else {
			fmt.Printf("%s is in a different life stage.\n", studentName)
		}
	}

	fmt.Println("\nDone.")
}