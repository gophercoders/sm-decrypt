package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/gophercoders/simpleio"
)

func main() {
	// tell the user what the program does.
	fmt.Println("sm-decrypt decrypts a word encrypted by the Caesar Cipher.")
	fmt.Println("Please enter the word to decrypt:")

	// message is the word the user wants to decrypt.
	var message string
	// Read the enrypted word that the user typed and store it in message
	message = simpleio.ReadStringFromKeyboard()

	// messageRunes holds the characters in encrypted message
	var messageRunes []rune
	messageRunes = []rune(message)
	// numberOfMessageRunes holds the number of characters (or runes) in the
	// encrypted message
	var numberOfMessageRunes int
	numberOfMessageRunes = utf8.RuneCountInString(message)

	// This is the the plain text alphabet, as a string
	// You decrypt to this alphabet.
	var plaintext string
	plaintext = "abcdefghijklmnopqrstuvwxyz"
	// plainTextRunes holds the characters that make up the plain text alphabet
	var plaintextRunes []rune
	plaintextRunes = []rune(plaintext)
	// The length of the plain text alphabet. This is the same
	// length as the cipher text alphabet in the case of the Caesar Cipher
	var plaintextLength int
	plaintextLength = utf8.RuneCountInString(plaintext)

	// The cipher text as a string
	// The cipher text is the alpahbet you decrypt from
	var ciphertext string
	ciphertext = "xyzabcdefghijklmnopqrstuvw"
	// cipherTextRunes holds the characters that make up the cipher text.
	var ciphertextRunes []rune
	ciphertextRunes = []rune(ciphertext)

	// i is the position of the current character in the messsage
	var i int
	fmt.Print("The decrypted Message is: ")
	// loop over the entire message
	i = 0
	for i < numberOfMessageRunes {
		// look for the position of the letter in the plain text
		var j int
		j = 0
		for j < plaintextLength {
			// Your turn
			// The encryption compares the characters (or runes) in the
			// message to the characters (or runes) in the plain text
			// The decryption has to compare the message characters (or runes)
			// in the message to the characters (or runes) in the cipher text.
			// What do you need to change in the line below?
			if messageRunes[i] == plaintextRunes[j] {
				// Your Turn
				// At this point you know the position of the letter (or rune)
				// in the plain text that you want to encrypt. The position
				// is stored in the variable j.
				// The encryption encrypts by looking up what letter is
				// at position j in the cipher text and printing it out.
				// The decryption has to decrypt by looking up what the letter is
				// at position j in the plain text printing it out
				// What do you need to change on the line below?
				fmt.Print(string(ciphertextRunes[j]))
				break // required! This stops the loop early
			}
			// There was no match. So try the next character in the plain text
			// alpahbet. For decryption thisis teh enxt letter in the cipher text
			// alpahbet.
			j = j + 1
		}
		// We have encrypted (or decrypted) the current character in the message
		// Now we need to try the next one.
		i = i + 1
	}
	fmt.Println()
}
