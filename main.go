package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
)

type Card struct {
	rank string
	//suit string
}

func ReadKey() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func ReadSprite(path string) []byte {
	sprite, _ := os.ReadFile(path)
	return sprite
}

func main() {
	var cmd string
	var ack string
	var fish bool
	var dealNot []Card

	const NEXT = 1
	const START = 0
	const SIZE = 52
	const SUIT = 4
	const HAND = 7
	const PLAYERS = 2
	const OVER = 2

	books := make([][]Card, PLAYERS)
	hands := make([][]Card, PLAYERS)
	deck := make([]Card, SIZE)
	player := START
	dealer := NEXT
	turn := START
	play := true
	rank := NEXT
	set := true
	dealOn := true
	over := false

	for set {
		play = true

		for card := range deck {
			if rank <= SIZE/SUIT {
				deck[card].rank = strconv.Itoa(rank)
				rank++
			} else {
				rank = NEXT
				deck[card].rank = strconv.Itoa(rank)
				rank++
			}
		}

		for card := range deck {
			random := rand.Intn(len(deck) - card)
			deck = append(deck, deck[random])
			deck = slices.Delete(deck, random, random+NEXT)
		}

		for range hands {
			hands[player] = append(hands[player], deck[:HAND]...)
			deck = slices.Delete(deck, START, HAND)

			player = NEXT
			dealer = START
			turn = NEXT
		}

		show := true

		for play {
			for show { // you wont do this status interface to be seen. you want a map or something that only looks like a map?
				var book [][]string

				lines := map[string]map[string]int{
					"cmd": {"LINE": 5, "LETTER": 16},
					"player": {"LINE": 5, "LETTER": 15},
					"dealer": {"LINE": 5, "LETTER": 14},
					"turn": {"LINE": 5, "LETTER": 13},
				}

				data := map[string]string{
					"cmd": cmd,
					"player": player,
					"dealer": dealer,
					"turn": turn,
				}

				for line, hand := range hands {
					for letter, card := range hand {
					    lines["hand"+strconv.Itoa(line)+strconv.Itoa(letter)] = map[string]int{"LINE": 4, "LETTER": letter + OVER}
        				data["hand"+strconv.Itoa(line)+strconv.Itoa(letter)] = card.rank
					}
				}
				
				for line, hand := range books {
					for letter, card := range hand {
					    lines["book"+strconv.Itoa(line)+strconv.Itoa(letter)] = map[string]int{"LINE": 3, "LETTER": letter + OVER}
        				data["book"+strconv.Itoa(line)+strconv.Itoa(letter)] = card.rank
					}
				}
				
				for letter, card := range deck {
					lines["deck"+strconv.Itoa(letter)] = map[string]int{"LINE": 2, "LETTER": letter + OVER}
        		    data["deck"+strconv.Itoa(letter)] = card.rank
				}
				
				for status := range lines {
					for coord := range lines[status] {
						if coord == "LINE" {
							for range lines[status][coord] - len(book) {
								book = append(book, []string{" "})
							}
						} else {
							for line := range book {
								for range lines[status][coord] - len(book[line]) {
									book[line] = append(book[line], " ")
								}
							}
						}
					}
				}

				for status := range lines {
					for feed := range data {
						if status == feed {
							book[lines[status]["LINE"]-NEXT][lines[status]["LETTER"]-NEXT] = data[feed]
						}
					}
				}

				for line := range book {
					fmt.Println(book[line])
				}

				break
			}

			for i, hand := range hands { // dont make the program to be shown, make it actvely to not be shown
				var deal []Card

				for _, card := range hand {
					for _, suit := range deal {
						if card.rank == suit.rank {
							dealOn = false
						}
					}

					for _, side := range hand {
						if card.rank == side.rank && dealOn == true {
							deal = append(deal, card)
						}
					}

					for len(deal)%SUIT != START {
						deal = slices.Delete(deal, len(deal)-NEXT, len(deal))
					}

					dealOn = true
				}

				if len(deal) != START {
					books[i] = append(books[i], deal...)

					for len(deal) != START {
						idx := slices.Index(hands[i], deal[START])

						hands[i] = slices.Delete(hands[i], idx, idx+NEXT)
						deal = slices.Delete(deal, START, NEXT)
					}
				}
			}

			dealOn = false

			if len(hands[player]) != START {
				for {
					if ack == "n" || ack == "" {
						cmd = ReadKey()
					}

					if cmd != "r" {
						switch cmd {
						case "0":
							cmd = "10"
						case "q":
							cmd = "11"
						case "w":
							cmd = "12"
						case "e":
							cmd = "13"
						}

						if casted, _ := strconv.Atoi(cmd); casted <= START || casted > SIZE/SUIT {
							fmt.Println(casted)
							continue
						}

						ack = ReadKey()

						if ack == "n" || ack != "y" {
							if ack == "" {
								ack = "void"
							}

							continue
						} else {
							ack = ""
							break
						}
					} else {
						ack = ReadKey()

						if ack == "n" || ack != "y" {
							if ack == "" {
								ack = "void"
							}

							continue
						} else {
							play = false
							set = false
							break
						}
					}
				}
			} else if over == true {
				for {
					cmd = ReadKey()

					if cmd == "y" {
						play = false
						deck = make([]Card, SIZE)
						books = make([][]Card, PLAYERS)

						break
					} else if cmd == "n" {
						set = false
						play = false
						break
					} else {
						continue
					}
				}
			} else {
				fish = true
			}

			if play == false {
				break
			}

			if len(deck) == START && len(hands[player]) == 0 && len(hands[dealer]) == 0 {
				over = true
				continue
			}

			for _, card := range hands[player] {
				if card.rank == cmd {
					dealOn = true
				}
			}

			if !dealOn && fish != true {
				continue
			} else if fish != true {
				for _, card := range hands[dealer] {
					if card.rank == cmd {
						hands[player] = append(hands[player], card)
					} else {
						dealNot = append(dealNot, card)
					}
				}

				if len(dealNot) == len(hands[dealer]) {
					fish = true
				}

				hands[dealer] = slices.Delete(hands[dealer], START, len(hands[dealer]))
				hands[dealer] = append(hands[dealer], dealNot...)

				dealNot = slices.Delete(dealNot, START, len(dealNot))
			}

			// remove from bellow deck is cool and codic, really aestetic with a physical eletronic deck that loads a card on a eletronic table, above is easier with paper, bellow is easier with eletrons
			if fish == true && len(deck) != START {
				hands[player] = append(hands[player], deck[START:START+NEXT]...)
				deck = slices.Delete(deck, START, START+NEXT)
				fish = false
			}

			//(about fish mainly)think like you still dont know what you are doing/want to see. dont put those above, logic first, status interface later, everything you do will fall into its places
			turn = dealer
			dealer = player
			player = turn
		}
	}
}
