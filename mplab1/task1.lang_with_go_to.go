package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	const SIZE = 15000
	type pairDic struct {
		str string
		num int
	}

	var arrayWords [SIZE]pairDic
	var banWords [SIZE]string
	banI := 0
	banFile, err := os.Open("banwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer banFile.Close()
	banScanner := bufio.NewScanner(banFile)
	banScanner.Split(bufio.ScanWords)
banInitLabel:
	banScanner.Scan()
	if banScanner.Text() == "" {
		goto readerLabel
	}
	if banI < len(banWords)-1 {
		banWords[banI] = banScanner.Text()
		banI++
		goto banInitLabel
	}

readerLabel:
	fmt.Println(banWords)
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	i := 0
	counter := 0
scannerLabel:
	scanner.Scan()
	currentWord := scanner.Text()
	out := []rune(currentWord)
	if currentWord == "" {
		goto Finish
	}
checkWordLabel:
	if out[0] >= 65 && out[0] <= 90 {
		out[0] = out[0] + 32
		currentWord = string(out)
	}
	if out[len(out)-1] >= 33 && out[len(out)-1]<=47{
		out2 := make([]rune, len(currentWord)-1)
		out[len(out)-1] = 0
		var iterator = 0
	jumpLabel:
		if out[iterator] != 0{
			out2[iterator] = out[iterator]
			iterator++
			goto jumpLabel
		}
		currentWord = string(out2)
	}
	if currentWord == banWords[counter] {
		counter = 0
		goto scannerLabel
	}
	if currentWord == arrayWords[counter].str {
		arrayWords[counter].num += 1
		counter = 0
		goto scannerLabel
	}
	if counter < len(arrayWords)-1 {
		counter++
		goto checkWordLabel
	}

	arrayWords[i].str = currentWord
	arrayWords[i].num = 1
	//fmt.Println(arrayWords[i])
	i++
	counter = 0
	goto scannerLabel

Finish:
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("######")
	m := 0
outer:
	t := 0
inner:

	if arrayWords[t].num < arrayWords[t+1].num {
		tmp := arrayWords[t]
		arrayWords[t] = arrayWords[t+1]
		arrayWords[t+1] = tmp
	}
	t += 1
	if t < len(arrayWords)-m-1 {
		goto inner
	}
	m += 1
	if m < len(arrayWords)-1 {
		goto outer
	}

	printCounter := 0
printLoopLabel:
	if printCounter < 20 && printCounter < len(arrayWords)-1 {
		fmt.Println(arrayWords[printCounter])
		printCounter++
		goto printLoopLabel
	}

}
