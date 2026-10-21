package main // Definiere diese Datei als package main

import (
	"fmt"
	"math/rand"
) // importiere das fmt (Format) Package für die Ausgabe

type RomanLetter rune // definiere RomanLetter als ein 32Bit UNICODE Zeichen (rune)

// Implement a function that counts the number of occurrences of roman letters in a roman numeral.
// Zähle alle Vorkommnisse der einzelnen Römischen Zählzeichen  (nur Majuskel) , z.b alle 'V' in den übergebebenen Feld, für jedes der möglichen Zeichen in dem Feld.
func task3a(input []RomanLetter) (int, int, int, int, int, int, int) {
	i := 0
	v := 0
	x := 0
	l := 0
	c := 0
	d := 0
	m := 0

	// für alle Literale in dem RomanLetter Feld
	//Zähle den Vorkommnis jedes einzelnen Literal types (switch...case)

	// wie die for mit range und switch...case funktioniert, siehe Skript

	for _, currentLetter := range input {
		switch {
		case currentLetter == 'I':
			i++
		case currentLetter == 'V':
			v++
		case currentLetter == 'X':
			x++
		case currentLetter == 'L':
			l++
		case currentLetter == 'C':
			c++
		case currentLetter == 'D':
			d++
		case currentLetter == 'M':
			m++
		}
	}
	return i, v, x, l, c, d, m
}

// Generate n roman numerals where the value of each roman numeral shall be in between low and high.
func task3b(low int, high int, n int) [][]RomanLetter {
	var allNumerals [][]RomanLetter //store all the roman numerals in here
	for i := 0; i < n; i++ {
		var randomNumber = rand.Intn(high-low+1) + low
		fmt.Println("chosen number: ", randomNumber)
		var newNumeral []RomanLetter
		// An der Stelle kann eine oder zwei  look-up-Table verwendet werden
		// TI2 Skript 4.4 Lookup Table (minb - Pech)
		// Die erste Tabelle speichert alle Kombinationen die einem Zählwert entsprechen z.B "I" und "IX" aber nicht "II" oder "XII"
		// die zweite Tabelle speichert die dazugehörenden Werte
		//so
		//  tabelleMitEinemSinvollenNamen = "I", "IV","V",....
		var romanValue = [][]rune{{'M'}, {'C', 'M'}, {'D'}, {'C', 'D'}, {'C'}, {'X', 'C'}, {'L'}, {'X', 'L'}, {'X'}, {'I', 'X'}, {'V'}, {'I', 'V'}, {'I'}}
		//  tabelleMitEinemAnderenSinvollenNamen = 1, 4, 5,....
		var arabicValue = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		// die beiden Tabellen lassen sich auch mit einer Struktur verwenden {"i",1}
		// nun muss ermittelt werden, welcher Wert in der Zahl noch vorhanden ist und entsprechend aus der Tabelle das dazugehörende Zählzeichen eingesetzt werden
		//   800 -> 1000 passt nicht rein, aber 500 also das Zeichen für 500 aus der Tabelle holen.  800 - 500 so bleibt 300 bestehen
		//   300 -> 1000 passt nicht rein, 300 -> 500 passt nicht rein, 300 -> ....
		for i, currentArabicValue := range arabicValue {
			for randomNumber >= currentArabicValue {
				var neededRomanLetter = romanValue[i]
				for _, e := range neededRomanLetter {
					newNumeral = append(newNumeral, RomanLetter(e))
				}
				randomNumber -= currentArabicValue
			}
		}
		fmt.Println("roman numeral to that number: " + romanLetterToString(newNumeral) + "\n")
		allNumerals = append(allNumerals, newNumeral)
	}
	return allNumerals
}

func romanLetterToString(romanLetters []RomanLetter) string {
	conversion := ""
	for _, currentRomanLetter := range romanLetters {
		conversion += string(currentRomanLetter)
	}
	return conversion
}

// Implement a parser for roman numerals that accepts a mix of lowercase and uppercase letters.
// Wandele den String in ein Feld vom Typ RomanLetter
// Überprüfe ob alle Zeichen in dem Übergegebenen String/Feld Römischen Zählzeichen angehören.
// Wandele dabei Zeichen die in Kleinbuchstaben vorliegen in die entsprechenden Zeichen in Großbuchstaben z.b. 'v' zu 'V'
// da Römisch Zahlen nur mit Majuskel notiert wurden, die Minuskeln sind erst im laufe des 3..4 Jahrhundert nach Christus entwickelt worden
// der zweite Rückgabewert besagt ob in der Zahl nur korrekte römische Zählzeichen befinden (true)
// oder sich dort auich Zeichen vorkommen die dort nichts zu suchen haben, z.B. 'a', 'H' etc.. (false)

func task3c(s string) ([]RomanLetter, bool) {
	convertedString := []RomanLetter(s)
	var onlyRomanLetters = true
	for index, letter := range convertedString {
		if 'a' <= letter && letter <= 'z' {
			letter = letter - ('a' - 'A')
		}
		convertedString[index] = letter

		if letter != 'I' && letter != 'V' && letter != 'X' && letter != 'L' && letter != 'C' && letter != 'D' && letter != 'M' {
			onlyRomanLetters = false
		}
	}
	// sehr ähnlich wie task3a , hier kann ein mehrfach case mit append verwendet werden (siehe Beispiel im Skript)

	if onlyRomanLetters {
		fmt.Println("output: Ist richtig, gut gemacht")
	} else {
		fmt.Println("output: Ist falsch, kys")
	}
	return convertedString, onlyRomanLetters
}

func main() {
	//task a
	fmt.Println("task a tests:")
	var testAlphabet = "MMXVIII"
	testRoman := []RomanLetter(testAlphabet)    // Test String
	var i, v, x, l, c, d, m = task3a(testRoman) // Mehrfachrückgabe der Parameter mit den gezählten Literalen
	fmt.Println("roman numeral: " + testAlphabet)
	fmt.Printf("i: %d | v: %d | x: %d | l: %d | c: %d | d: %d | m: %d \n", i, v, x, l, c, d, m)

	// task b
	fmt.Println("\ntask b tests:")
	var _ = task3b(1, 1000, 10)
	var numberToConvert = 19
	_ = task3b(numberToConvert, numberToConvert, 1)
	numberToConvert = 444
	_ = task3b(numberToConvert, numberToConvert, 1)
	numberToConvert = 999
	_ = task3b(numberToConvert, numberToConvert, 1)
	numberToConvert = 3
	_ = task3b(numberToConvert, numberToConvert, 1)

	// task c
	fmt.Println("\ntask c tests:")
	var test = "mMxvIIi"
	fmt.Println("input: " + test)
	task3c(test)
	test = "sknglng"
	fmt.Println("input: " + test)
	task3c(test)
}
