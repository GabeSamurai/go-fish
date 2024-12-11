package main

import (
    "fmt"
    "math/rand"
    "slices"
)

type Card struct {
    rank int
    suit string
}

//code file/device memory type
func New() [52]Card {
    var cards [52]Card
    const suitSize = 12
    
    for ord := range cards {
        cards[ord].rank = ord + 1
        
        switch {
            case ord < 13: cards[ord].suit = "spades"
            case ord < 26: cards[ord].suit = "diamonds"
            case ord < 39: cards[ord].suit = "clubs"
            
            default: cards[ord].suit = "hearts"
        }
    }
    
    return cards
}

//code file/device memory type
func Shuffle(cards [52]Card) []Card {
    var shuffled []Card
    var ind, indS []int
    
    for i := range cards {
        ind = append(ind, i)
    }
    
    for len(ind) != 0 {
        random := rand.Intn(len(ind))
        indS = append(indS, ind[random])
        
        ind = slices.Delete(ind, random, random+1)
    }
    
    for i := range indS {
        shuffled = append(shuffled, cards[indS[i]])
    }
    
    return shuffled
}

func main() {
    goDeck := New()
    shuDeck := Shuffle(goDeck)
    
    fmt.Println(goDeck)
    fmt.Println(shuDeck)
}