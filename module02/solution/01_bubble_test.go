package module02

import (
	"algo/module02/sorttest"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	sorttest.TestInt(t, BubbleSortInt)
}

func BenchmarkBubbleSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, BubbleSortInt)
}

func TestBubbleSortString(t *testing.T) {
	sorttest.TestString(t, BubbleSortString)
}

func TestBubbleSortInterface(t *testing.T) {
//	sorttest.TestInterface(t, BubbleSort)
}

func BenchmarkBubbleSortInterface(b *testing.B) {
//	sorttest.BenchmarkInterface(b, BubbleSort)
}

func TestBubbleSortPerson(t *testing.T) {
	testPeople(t, BubbleSortPerson)
}
