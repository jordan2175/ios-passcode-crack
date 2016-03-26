// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Lets prompt the user for the encrypte iOS Restrictions Passcode
	// and remove the trailing newline character if it is there
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter encrypted passcode from iOS device: ")
	rawiospasscode, _ := reader.ReadString('\n')
	rawiospasscode = strings.TrimSuffix(rawiospasscode, "\n")

	// Lets prompt the user for the salt that was used in the hashing function
	// and one again, remote the newline character if there.
	fmt.Print("Enter salt from iOS device: ")
	rawiossalt, _ := reader.ReadString('\n')
	rawiossalt = strings.TrimSuffix(rawiossalt, "\n")

	// Lets start a timer so we know how long it takes to run this tool
	startTime := time.Now()

	// The salt is stored in base64, lets decode that to a string
	salt, _ := base64.StdEncoding.DecodeString(rawiossalt)

	// We will loop through every possible 4 digit passcode to find the
	// matching passcode
	iCounter := 0
	fmt.Println("\nEach '.' represents 10 attempted passcodes...")
	for iPasscode := 0; iPasscode < 9999; iPasscode++ {

		// The hashing functions needs a byte array so lets convert the integer
		// from the for loop to a strong
		sPasscode := strconv.Itoa(iPasscode)

		// Computer the hash based on the same salt and the index of the for loop
		hashedPasscode := pbkdf2.Key([]byte(sPasscode), []byte(salt), 1000, 20, sha1.New)

		// Convert the results of the hashing function to base64 so we can
		// compare it with what the user originally typed in
		b64HashedPasscode := base64.StdEncoding.EncodeToString(hashedPasscode)

		// Write some output to the screen so users know we are doing something
		iCounter++
		if iCounter%10 == 0 {
			fmt.Print(".")
			if iCounter%800 == 0 {
				fmt.Print("\n")
			}
		} // screen feedback

		// Check to see if we have a match
		if b64HashedPasscode == rawiospasscode {
			fmt.Println("\nMatch Found! Your passcode is: " + sPasscode)
			elapsedTime := time.Since(startTime)
			fmt.Println("\niOS Passcode Cracker took: ", elapsedTime)
			os.Exit(0)
		}
	}
	fmt.Println("\nNo match was found!")
	fmt.Println("Please check that you copied the passcode and salt correctly from your iOS backup.")
}
