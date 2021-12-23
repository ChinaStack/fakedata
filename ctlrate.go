package main

import (
	"fmt"
	wr "github.com/mroth/weightedrand"
	"math/rand"
	"time"
)

func rate() {
	b := make(map[uint32]int32)
	b[0] = 10
	b[1] = 20
	b[2] = 30
	b[3] = 40
	h2 := make([]int, 4)
	for x := 0; x < 200000000; x++ {
		r := rand.Int31n(99)
		i, j := int32(0), int32(0)
		for k, v := range b {
			j = i + v - 1
			if r >= i && r <= j {
				h2[k]++
				break
			}
			i += v
		}
	}
	for _, b := range h2 {
		fmt.Print(float32(b)/2000000, " ")
	}
}

func RandomWeight() {

	rand.Seed(time.Now().UTC().UnixNano()) // always seed random!

	chooser, _ := wr.NewChooser(
		wr.Choice{Item: "🍒", Weight: 0},
		wr.Choice{Item: "🍋", Weight: 1},
		wr.Choice{Item: "🍊", Weight: 1},
		wr.Choice{Item: "🍉", Weight: 3},
		wr.Choice{Item: "🥑", Weight: 5},
	)
	/* The following will print 🍋 and 🍊 with 0.1 probability, 🍉 with 0.3
	   probability, and 🥑 with 0.5 probability. 🍒 will never be printed. (Note
	   the weights don't have to add up to 10, that was just done here to make the
	   example easier to read.) */
	result := chooser.Pick().(string)
	fmt.Println(result)
}
