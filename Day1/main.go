package main

import (
	"fmt"
	"time"
)

var input = [...]int{
	95423,
	142796,
	88137,
	105610,
	79299,
	110633,
	136792,
	112578,
	75168,
	115615,
	147584,
	72145,
	108822,
	57753,
	96827,
	69117,
	131220,
	111193,
	120295,
	56240,
	111190,
	80740,
	137267,
	113183,
	126821,
	58966,
	63556,
	110977,
	100328,
	75367,
	57371,
	88235,
	134475,
	109071,
	92653,
	73347,
	135186,
	64534,
	81198,
	55423,
	100060,
	149555,
	110905,
	102826,
	129023,
	112618,
	146542,
	102579,
	67193,
	84258,
	60679,
	86674,
	124720,
	68719,
	55259,
	76421,
	70397,
	67998,
	73366,
	106401,
	59402,
	112481,
	131113,
	142606,
	107732,
	69291,
	61575,
	131019,
	51510,
	101215,
	116973,
	63530,
	146179,
	132427,
	127777,
	127040,
	143964,
	120340,
	144404,
	72156,
	96412,
	140554,
	60228,
	52590,
	128157,
	120444,
	125649,
	111641,
	117476,
	139326,
	149188,
	133599,
	55273,
	83773,
	50458,
	105166,
	76469,
	66681,
	84288,
	103708,
}

func main() {
	// start tracking time to run
	start := time.Now()

	// response := 0
	fmt.Println("Hello, World!")
	jobs := make(chan int, len(input))
	results := make(chan int, len(input))
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	for _, item := range input {
		fmt.Println(item)
		jobs <- item
		// response += CalculateFuel(item)

		// fmt.Println(response)
	}
	close(jobs)
	for j := 0; j < len(input); j++ {
		fmt.Println(<-results)
	}
	// fmt.Println(response)

	// tracking time to run
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}

// worker processes items in order
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- CalculateFuel(n)
	}
}

// CalculateFuel is to calculate the mass for Santa's sleigh
func CalculateFuel(mass int) int {
	response := int((mass / 3) - 2)
	return response
}
