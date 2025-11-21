/**
 * @author Bianca Boo
 * @version 1.0.0
 * @date 2025-11-16
 * @fileoverview This program will determine and output age cateogory.
 */

// Declare variables
let studentName: string = "";
let age: number = 0;
let isStudent: boolean = false;

// Initialize variables
studentName = prompt("Enter your name:") || "";
age = Number(prompt("Enter your age:") || "0");
isStudent = (prompt("Are you a student? (true/false)") || "false").trim().toLowerCase() === "true";

// Output
if (isStudent) {
  if (age >= 13 && age <= 19) {
    console.log(`Student: ${studentName} is a teenager.`);
  } else if (age >= 5 && age <= 12) {
    console.log(`Student: ${studentName} is a child.`);
  } else {
    console.log(`Student: ${studentName} is in a different life stage.`);
  }
} else {
  if (age >= 20 && age <= 30) {
    console.log(`${studentName} is a young adult.`);
  } else {
    console.log(`${studentName} is in a different life stage.`);
  }
}

console.log("\nDone.");
