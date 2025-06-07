package main

import (
	"log"
)

type node struct {
	name string
	pos  bool
}

type signTable map[string]map[string]bool
type table struct {
	equals    signTable
	notEquals signTable
}

func (t signTable) addNode(to, node string) {
	if t[to] == nil {
		t[to] = make(map[string]bool)
		t[to][node] = true
		for n := range t[node] {
			t[to][n] = true
		}
	}
}

func (t signTable) signAdd(a, b string) {
	t.addNode(a, b)
	t.addNode(b, a)
}

func (t *table) Add(a, b string, equal bool) {
	if equal {
		t.equals.signAdd(a, b)
	} else {
		t.notEquals.signAdd(a, b)
	}
}

func (t *table) check() bool {
	for node := range t.equals {
		for other := range t.equals[node] {
			if t.equals[node][other] == t.notEquals[node][other] {
				return false
			}
		}
	}
	for node := range t.notEquals {
		if t.notEquals[node][node] == true {
			return false
		}
	}
	return true
}

func test0() {
	tab := table{make(map[string]map[string]bool), make(map[string]map[string]bool)}
	if tab.check() != true {
		log.Fatal("test0")
	}
}

func test1() {
	tab := table{make(map[string]map[string]bool), make(map[string]map[string]bool)}
	tab.Add("A", "B", true)
	tab.Add("B", "C", true)
	tab.Add("C", "D", true)
	tab.Add("D", "A", true)
	if tab.check() != true {
		log.Fatal("test1")
	}
}

func test2() {
	tab := table{make(map[string]map[string]bool), make(map[string]map[string]bool)}
	tab.Add("A", "B", true)
	tab.Add("B", "C", true)
	tab.Add("C", "D", true)
	tab.Add("D", "A", false)
	if tab.check() != false {
		log.Fatal("test2")
	}
}

func test3() {
	tab := table{make(map[string]map[string]bool), make(map[string]map[string]bool)}
	tab.Add("A", "A", true)
	if tab.check() != true {
		log.Fatal("test3")
	}
}

func test4() {
	tab := table{make(map[string]map[string]bool), make(map[string]map[string]bool)}
	tab.Add("A", "A", false)
	if tab.check() != false {
		log.Fatal("test4")
	}
}

func main() {
	test0()
	test1()
	test2()
	test3()
	test4()
}
