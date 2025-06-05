package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

var bitCount [256]uint

func init() {
	for i := range bitCount {
		bitCount[i] = bitCount[i/2] + uint(i&1)
	}
}

func uint64BitCount(x uint64) (count int) {
	for i := range 8 {
		b := uint8(x >> (i * 8))
		count += int(bitCount[b])
	}
	return
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	s.forEachWord(t, func(a, b uint64) uint64 {
		return a | b
	})
}

func (s *IntSet) IntersectWith(t *IntSet) {
	s.forEachWord(t, func(a, b uint64) uint64 {
		return a & b
	})
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	s.forEachWord(t, func(a, b uint64) uint64 {
		return a & ^b
	})
}

func (s *IntSet) SymDifferenceWith(t *IntSet) {
	s.forEachWord(t, func(a, b uint64) uint64 {
		return a ^ b
	})
}

func (s *IntSet) forEachWord(t *IntSet, op func(a, b uint64) uint64) {
	for i, sword := range s.words {
		if i < len(t.words) {
			s.words[i] = op(sword, t.words[i])
		} else {
			break
		}
	}
	for i := len(s.words); i < len(t.words); i++ {
		s.words = append(s.words, op(0, t.words[i]))
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := range 64 {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() uint {
	len := uint(0)
	for _, word := range s.words {
		len += bitCount[uint64BitCount(word)]
	}
	return len
}

func (s *IntSet) Remove(x int) {
	s.words[x/64] &= ^(1 << (uint64(x) % 64))
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var newSet IntSet
	for _, word := range s.words {
		newSet.words = append(newSet.words, word)
	}
	return &newSet
}

func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i := range 64 * len(s.words) {
		if s.Has(i) {
			elems = append(elems, i)
		}
	}
	return elems
}

func main() {
	var set IntSet
	set.Add(2)
	set.Add(8)
	set.Add(200)
	fmt.Println(&set)
	set.Remove(8)
	fmt.Println("set after remove", &set)
	fmt.Println("set size:", set.Len())
	set2 := set.Copy()

	fmt.Println("new fresh set:", set2)
	set2.Remove(200)
	fmt.Println("old set:", &set)
	fmt.Println("new set:", set2)
	set.AddAll(1, 3, 4, 5, 6)
	fmt.Println("after add all:", &set)
	var set3, set4 IntSet
	set3.AddAll(1, 2, 3)
	set4.AddAll(1, 4, 5)
	set3.SymDifferenceWith(&set4)
	fmt.Println(&set3)
	var set5 IntSet
	set5.AddAll(100, 200, 300)
	fmt.Println(&set5)
	set3.UnionWith(&set5)
	fmt.Println(&set3)

	for _, num := range set3.Elems() {
		fmt.Println(num)
	}
}
