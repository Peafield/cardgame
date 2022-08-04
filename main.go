package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// TODO: FIX TEN TO NOT ADD EXTRA SPACE

func suitSwtichCase(suit string) {
	switch {
	case suit == "Hearts":
		fmt.Print("♥")
	case suit == "Diamonds":
		fmt.Print("◆")
	case suit == "Spades":
		fmt.Print("♠")
	case suit == "Clubs":
		fmt.Print("♣")
	}
}

func printCard(suit string, value string) {
	height := 9
	width := 9
	for h := 1; h <= height; h++ {
		for w := 1; w <= width; w++ {
			if h == 1 && w == 1 || h == height && w == width {
				fmt.Print("/")
			} else if h == 1 && w == width || h == height && w == 1 {
				fmt.Print("\\")
			} else if h == 1 || h == height {
				fmt.Print("-")
			} else if w == 1 || w == width {
				fmt.Print("|")
			} else {
				switch {
				case value == "Ace":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("A")
					} else if h == 5 && w == 5 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Two":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("2")
					} else if h == 3 && w == 3 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Three":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("3")
					} else if h == 3 && w == 5 || h == 5 && w == 5 || h == 7 && w == 5 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Four":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("4")
					} else if h == 3 && w == 3 || h == 7 && w == 3 || h == 3 && w == 7 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Five":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("5")
					} else if h == 3 && w == 3 || h == 7 && w == 3 || h == 5 && w == 5 || h == 3 && w == 7 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Six":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("6")
					} else if h == 3 && w == 3 || h == 5 && w == 3 || h == 7 && w == 3 || h == 3 && w == 7 || h == 5 && w == 7 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Seven":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("7")
					} else if h == 3 && w == 3 || h == 5 && w == 3 || h == 7 && w == 3 || h == 5 && w == 5 || h == 3 && w == 7 || h == 5 && w == 7 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Eight":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("8")
					} else if h == 3 && w == 3 || h == 5 && w == 3 || h == 4 && w == 5 || h == 6 && w == 5 || h == 7 && w == 3 || h == 3 && w == 7 || h == 5 && w == 7 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Nine":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("9")
					} else if h == 3 && w == 3 || h == 4 && w == 3 || h == 5 && w == 3 || h == 6 && w == 3 || h == 7 && w == 3 || h == 5 && w == 5 || h == 3 && w == 7 || h == 4 && w == 7 || h == 5 && w == 7 || h == 6 && w == 7 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else {
						fmt.Print(" ")
					}
				case value == "Ten":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("10")
					} else if h == 3 && w == 3 || h == 4 && w == 3 || h == 5 && w == 3 || h == 6 && w == 3 || h == 4 && w == 5 || h == 6 && w == 5 || h == 7 && w == 3 || h == 3 && w == 7 || h == 4 && w == 7 || h == 5 && w == 7 || h == 6 && w == 7 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else {
						if h == 8 && w == 2 || h == 2 && w == 9 || h == 2 && w == 8 {
							fmt.Print("")
						} else {
							fmt.Print(" ")
						}

					}
				case value == "Jack":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("J")
					} else if h == 3 && w == 3 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else if h == 5 && w == 5 {
						fmt.Print("♝")
					} else {
						fmt.Print(" ")
					}
				case value == "Queen":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("Q")
					} else if h == 3 && w == 3 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else if h == 5 && w == 5 {
						fmt.Print("♛")
					} else {
						fmt.Print(" ")
					}
				case value == "King":
					if h == 2 && w == 2 || w == 8 && h == 8 {
						fmt.Print("K")
					} else if h == 3 && w == 3 || h == 7 && w == 7 {
						suitSwtichCase(suit)
					} else if h == 5 && w == 5 {
						fmt.Print("♚")
					} else {
						fmt.Print(" ")
					}
				}
			}
		}
		fmt.Println()
	}
}

func getCardValue(n int) string {
	value := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	return value[n]
}

func getSuit(n int) string {
	suit := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	return suit[n]
}

func getRandomNum(n int) int {
	return rand.Intn(n)
}

func getCard(s, v int) string {
	return getCardValue(v) + " of " + getSuit(s)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ranSuit := getRandomNum(4)
	ranVal := getRandomNum(13)
	chosenCard := getCard(ranSuit, ranVal)
	splitCard := strings.Split(chosenCard, " ")
	splitSuit := splitCard[2]
	splitValue := splitCard[0]
	fmt.Println(chosenCard)
	printCard(splitSuit, splitValue)
}
