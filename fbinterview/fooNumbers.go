// https://www.careercup.com/question?id=5692469780938752

//
//Question 1.
//You are given a string composed of uppercase English letters (‘A’ through ‘Z’).
//
//Set of letters (‘A’, ‘E’, ‘I’, ‘O’, ‘U’) are called vowels. Other letters are called consonants.
//
//We define foo value of a string as number of pairs of exactly same consecutive vowel letters.
//
//For example,
//Ex.1: BCDEEIOU - This has a foo value of 1 (because of EE). Note that although I is next to E, and O is next to I, and U is next to O, they aren’t exactly same neighbours, so they don’t contribute to foo value
//
//Ex.2: BCDEEEIOU - This has foo value of 2. Because of first pair of EE and immediately next pair of EE
//
//Ex.3: ABCDEFG - This has foo value of 0. There are no consecutive vowels
//
//Ex.4: ABEUUOUOO - This has foo value of 2, because of UU and OO
//
//You are given 2 inputs, N and K.
//How many strings of length N can you form such that they all have foo value of K?
//
//Let’s assume the constraints as:
//
//1<=N<=15
//0<=K<N

package fbinterview

func NumberOfFoos(n, k int) {
	// k / 2

	//  ‘A’, // ‘E’, // ‘I’, ‘O’, ‘U’
	// n = 1, k = 1
	// -- AA***** AAA*** AAAA***
	// *AAA**
	// **AAA*
	// AA*AA*AA
	// AA*AA*AA*
	// AA*AA*AA

	//2
	// 00000111111
	// 11011011011

}
