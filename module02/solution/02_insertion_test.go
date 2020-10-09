package module02

import (
	"algo/module02/sorttest"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, InsertionSortInt)
}

func BenchmarkInsertionSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, InsertionSortInt)
}

func TestInsertionSortInterface(t *testing.T) {
	sorttest.TestInterface(t, InsertionSort)
}

func BenchmarkInsertionSortInterface(b *testing.B) {
	sorttest.BenchmarkInterface(b, InsertionSort)
}