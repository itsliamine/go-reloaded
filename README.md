# Go Reloaded

This program takes two file names as arguments. The content of the first file if formatted according to the rules. The second file is created with the name provided in the second argument and the formatted content is written in it. <br>
The project is divided in steps : <br>
* The [core](core) files are the main functions that format the string
* The [utils](utils) files are mainly focused on string, array and char manipulation or types and base conversion.


### Unit Test

The **[Unit Test](main_test.go)** takes every sample in the root folder and executes every core function on them. To run it :

```bash
git clone https://zone01normandie.org/git/faoudia/go-reloaded.git
cd go-reloaded
go test -v
```

Expected output:

```bash
=== RUN   TestSample1
--- PASS: TestSample1 (0.00s)
=== RUN   TestSample2
--- PASS: TestSample2 (0.00s)
=== RUN   TestSample3
--- PASS: TestSample3 (0.00s)
=== RUN   TestSample4
--- PASS: TestSample4 (0.00s)
=== RUN   TestSample5
--- PASS: TestSample5 (0.00s)
PASS
ok  	go-reloaded	0.002s
```

### Core functions

* **[core/fileopen.go](core/fileopen.go)**

Given a filename, the FileOpen function opens it, reads its content and returns the result in a string.

```go
func FileOpen(filename string) string {
	content, err := os.ReadFile(filename)

	if err != nil {
		return ""
	}

	return string(content)
}
```
* **[core/convert.go](core/convert.go)**

Given the file content previously read, the Convert function is going to locate every instance of the word ```(hex)``` or ```(bin)```. The number located behind them will be converted from either hexadecimal or binary to decimal. <br>
```go
    words := lib.SplitWhiteSpaces(s)

	for i := 1; i < len(words); i++ {
		if lib.Contains(words[i], "(hex)") {
			words[i-1] = lib.HexConvert(words[i-1])
			words = lib.RemoveSlice(words, i)
		}

		if lib.Contains(words[i], "(bin)") {
			words[i-1] = lib.BinConvert(words[i-1])
			words = lib.RemoveSlice(words, i)
		}
	}
```

The given string is splitted using the function ```SplitWhiteSpaces(string)``` located in **[utils/stringfuncs.go](utils/stringfuncs.go)**. If we come across the word ```(hex)``` or ```(bin)```, the number before will be converted using the functions located in **[utils/baseconvert.go](utils/baseconvert.go)**.
At the end, a string is returned.

* **[core/letters.go](core/letters.go.go)**

Given the string previously returned, the Letters function is going to locate every instance of the words ```(cap)``` for capitalize, ```(up)``` for upper case or ```(low)``` for lower case. The previous word is going to be formatted accordingly. There's also the possiblity of coming across options written like ```(cap, n)```, ```n``` representing the number of words that will be affected by the formatting.  <br>

```go
if lib.Contains(words[i], "(cap,") {
	num := lib.Atoi(string(words[i+1][0]))
	for x := i; x >= i-num; x-- {
		words[x] = lib.Capitalize(words[x])
	}
	words = lib.RemoveSlice(words, i)
	words = lib.RemoveSlice(words, i)
}

if lib.Contains(words[i], "(cap)") {
	words[i-1] = lib.Capitalize(words[i-1])
	words = lib.RemoveSlice(words, i)
}
```

The functions ```Capitalize(string)```, ```ToLower(string)``` and ```ToUpper(string)``` located in **[utils/charfuncs.go](utils/charfuncs.go)** handle the formatting.

* **[core/punc.go](core/punc.go)**

Given the string previously returned, we'll take a look at each character. If we come across a comma, a dot, a colon, a semi-colon, an exclamation mark or a question mark, it is going to be attached to the previous word and a space should be left after.

The ``Ìnsert(str, i, rune)``` and ```RemoveChar(str, index)``` functions, help manipulating the string. They are located in **[utils/stringfuncs.go](utils/stringfuncs.go)**

* **[core/quotes.go](core/quotes.go)**

We'll still take a look at each character. A ```quoteBeginning```variable is created to track if we already found a quote. This way, we will be able to put the quotes correctly on the left and right side of the word(s) in the middle of them.

* **[core/vowels.go](core/vowels.go)**

We need to format any instance of the letter "a" to the word "an" if the next word starts by a vowel. While looping through our characters, if we come across the letter "a", we will check if the character that is not a space is a vowel. We'll add an "n" character (or not) accordingly. 

* **[core/fileoutput.go](core/fileoutput.go)**

This gets the formatted string and writes to the file with the given filename.
