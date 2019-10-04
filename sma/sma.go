package main

import (
	"errors"
	"fmt"
	"strings"
)

const maxSize = 1000

type sma struct {
	size uint
	data []float32
	sum  float64
}

type smaOps interface {
	addData(val float32)
	getAvg() float64
	toString() string
}

func newSma(size uint) *sma {
	return &sma{
		size: size,
		data: make([]float32, 0),
		sum:  0.0,
	}
}

func (sma *sma) addData(val float32) {
	sma.sum += float64(val)
	sma.data = append(sma.data, val)

	if uint(len(sma.data)) > sma.size {
		sma.sum -= float64(sma.data[0])
		sma.data = sma.data[1:]
	}
}

func (sma *sma) getAvg() (float64, error) {
	if uint(len(sma.data)) < sma.size {
		return 0.0, errors.New("Not enough data points.")
	}

	return (sma.sum / float64(sma.size)), nil
}

func (sma *sma) toString() string {
	elems := []string{}

	for _, elem := range sma.data {
		e := fmt.Sprintf("%f", elem)
		elems = append(elems, e)
	}

	return strings.Join(elems, " ")
}

func main() {
	var size uint
	var val float32

	fmt.Printf("Enter window size: ")
	fmt.Scanf("%d", &size)
	if size < 1 || size > maxSize {
		fmt.Printf("Invalid window size!")
		return
	}

	sma := newSma(size)

	fmt.Printf("Enter numbers for calculating average - \n")
	for i := 0; i < maxSize; i++ {
		fmt.Scanf("%f", &val)
		sma.addData(val)

		avg, err := sma.getAvg()
		if err != nil {
			fmt.Printf("Not enough data points!\n\n")
			continue
		}

		fmt.Printf("\n Average: %f (%s)\n\n", avg, sma.toString())
	}
}
