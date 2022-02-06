package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	const PAGE_WORDS = 521
	const SIZE = 15000
	type pairDic struct {
		str string
		count int
		pageNumbers[SIZE]int
		//curPage int
		inArrCounter int
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
	wordsOnPageCounter := 0
	curPage := 1
scannerLabel:
	scanner.Scan()
	currentWord := scanner.Text()
	out := []rune(currentWord)
	if currentWord == "" {
		goto Finish
	}
	if wordsOnPageCounter == PAGE_WORDS {
		curPage++
		wordsOnPageCounter = 0
	}
	wordsOnPageCounter++
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
		arrayWords[counter].count += 1
		arrayWords[counter].pageNumbers[arrayWords[counter].inArrCounter] = curPage
		arrayWords[counter].inArrCounter++
		counter = 0
		goto scannerLabel
	}
	if counter < len(arrayWords)-1 {
		counter++
		goto checkWordLabel
	}

	arrayWords[i].str = currentWord
	arrayWords[i].count = 1
	arrayWords[i].pageNumbers[arrayWords[i].inArrCounter] = curPage
	arrayWords[i].inArrCounter++
	//fmt.Println(arrayWords[i])
	i++
	counter = 0
	goto scannerLabel

Finish:
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("######")
	//out := []rune(currentWord)
	m := 0
outer:
	t := 0
inner:
	/*var out1 = []rune(arrayWords[t].str)
	var out2 = []rune(arrayWords[t+1].str)*/
	if arrayWords[t].str == "" || arrayWords[t+1].str == ""{
		goto checking
	}
	if arrayWords[t].str/*out1[0]*/ > arrayWords[t+1].str /*out2[0]*/ {
		tmp := arrayWords[t]
		arrayWords[t] = arrayWords[t+1]
		arrayWords[t+1] = tmp
	}
	checking:
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
		if arrayWords[printCounter].count <= 100 {
			fmt.Println(arrayWords[printCounter])
		}
		//fmt.Println(arrayWords[printCounter])
		printCounter++
		goto printLoopLabel
	}

}
