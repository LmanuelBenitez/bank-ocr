package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	getDigits()
	fmt.Print("\nAccount numbers written in the Resuts file successfully...\n")
}

// Function to get the digits from the escaner file
func getDigits() {

	fileData, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Error reading file with error: %v", err)
	}

	//Create the result file to save the account numbers with the status
	fileResults, err := os.Create("account_numbers.txt")
	if err != nil {
		fmt.Printf("Error creating results file with error: %v", err)
	}

	defer fileData.Close()
	defer fileResults.Close()

	//Map that sets the digits patterns
	digitPatterns := setDigitPatterns()

	// EXAMPLE OF THE  segmentDigits MATRIX using 501507028
	// [[" _ "," _ ","   "," _ "," _ "," _ "," _ "," _ "," _ "],
	// ["|_ ","| |","  |","|_ ","| |","  |","| |"," _|","|_|"],
	// [" _|","|_|","  |"," _|","|_|","  |","|_|","|_ ","|_|"]]
	segmentDigits := [][]string{} //Matrix for capturing the segments of the digits
	counterLines := 1             //Variable to control the lines with valid data

	scanner := bufio.NewScanner(fileData)

	for scanner.Scan() {
		line := scanner.Text() //Read the line from the file

		digitsInLine := getNumbersSegment(line, counterLines)

		//If the digits has content then it is saved in the matrix and the counterLine is incremented
		if digitsInLine != nil {
			segmentDigits = append(segmentDigits, digitsInLine)
			counterLines++

			//The operation indicates that all the digits have been read
			if counterLines%4 == 0 {
				getAccountNumbers(fileResults, digitPatterns, segmentDigits) //Get the digits using the digitPatterns and segmentDigits
				segmentDigits = [][]string{}
			}

		} else {
			//The variable is incremented to get the next digits line
			counterLines++
		}

	}

}

// Function to get number segments
// In this case we divide all the line in segments of 3 characters to have a part of each digit
// The return value is the array with the 3 characters that are saved in the segmentDigits matrix
func getNumbersSegment(line string, counterLines int) []string {

	//If the line contains characters or if the module operation is not equal to 0 the characters are gotten
	if strings.TrimSpace(line) != "" || counterLines%4 != 0 {
		characters := strings.Split(line, "")

		digitsInLine := []string{} //Array that stores groups of 3 characters

		//Iterate through the characters in the line to get only 3 characters
		for i := 0; i < len(characters); i += 3 {
			digitsInLine = append(digitsInLine, strings.Join(characters[i:i+3], ""))
		}

		return digitsInLine

	}
	return nil
}

// Function to convert characters in the matrix to the corresponding digit just if the characteres forms a valid digit
func getAccountNumbers(fileResults *os.File, digitPatterns map[string]int, segmentDigits [][]string) {
	var accountNumbers = [9]int{} //Store the account numbers

	//Iterate through the colums of the matrix concatenating the upperLine, middleLine and lowerLine characters to form a possible digit
	for j := 0; j < 9; j++ {
		possibleNumber := segmentDigits[0][j] + segmentDigits[1][j] + segmentDigits[2][j]

		//If the digit exists it's stored in the accountNumbers array, otherwise it's stored with -1
		number, exists := digitPatterns[possibleNumber]
		if exists {
			accountNumbers[j] = number
		} else {
			accountNumbers[j] = -1
		}
	}

	result := setResults(accountNumbers[:])
	_, err := fileResults.WriteString(result) //The return value from setResults is written in the fileResults
	if err != nil {
		fmt.Printf("Error writing in the file, with error: %v", err)
	}

}

// Function to set the status of the account numbers
func setResults(accountNumbers []int) string {
	//Iterate through the account numbers to verify that the values are valid
	accountNumbersString := ""
	for _, number := range accountNumbers {
		//If the digit is equal to -1 we replace it with '?'
		if number == -1 {
			accountNumbersString += "?"
		} else {
			accountNumbersString += strconv.Itoa(number)
		}
	}

	var result string = ""

	//Using the checkSumCalculation function to establish the status of the account
	if strings.Contains(accountNumbersString, "?") {
		result = fmt.Sprintf("%s ILL\n", accountNumbersString)
		fmt.Printf("Unrecognizable account digits: %s\n", result)
		return result

	} else if checkSumCalculation(accountNumbers) {
		result = fmt.Sprintf("%s OK\n", accountNumbersString)
		fmt.Printf("Valid account digits: %s\n", result)
		return result

	} else {
		result = fmt.Sprintf("%s ERR\n", accountNumbersString)
		fmt.Printf("Invalid account digits: %s\n", result)
		return result
	}

}

func checkSumCalculation(accountNumbers []int) bool {
	//Map to set the account number and the position name
	numberPosition := map[string]int{
		"d9": 3,
		"d8": 4,
		"d7": 5,
		"d6": 8,
		"d5": 8,
		"d4": 2,
		"d3": 8,
		"d2": 6,
		"d1": 5,
	}

	var checkSum int = 0

	//Iterate over the accountNumbers to calculate the sum using (n1*d1 + n2*d2 + n2*d3 + ...) % 11 == 0
	for index, number := range accountNumbers {
		numberByPositionName := numberPosition["d"+strconv.Itoa(index+1)]
		checkSum += number * numberByPositionName
	}

	return checkSum%11 == 0 //Verify if the checkSum is equal to 0
}

// Set the digit patterns
func setDigitPatterns() map[string]int {

	return map[string]int{
		" _ " +
			"| |" +
			"|_|": 0,

		"   " +
			"  |" +
			"  |": 1,

		" _ " +
			" _|" +
			"|_ ": 2,

		" _ " +
			" _|" +
			" _|": 3,

		"   " +
			"|_|" +
			"  |": 4,

		" _ " +
			"|_ " +
			" _|": 5,

		" _ " +
			"|_ " +
			"|_|": 6,

		" _ " +
			"  |" +
			"  |": 7,

		" _ " +
			"|_|" +
			"|_|": 8,

		" _ " +
			"|_|" +
			" _|": 9,
	}
}
