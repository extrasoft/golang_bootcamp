package main

import "fmt"

type Sword struct {
	name string
}

func (s Sword) detail() string {
	return "Super cold Iced sword"
}

func (s Sword) cost() int {
	return 500
}

type Bow struct {
	name string
}

func (b Bow) detail() string {
	return "Magic fire-bow"
}

func (b Bow) cost() int {
	return 200
}

type Item interface {
	cost() int
}

type Weapon interface {
	detail() string
	Item
}

func attack(w Weapon) {
	fmt.Printf("attack by : %s, %d\n", w.detail(), w.cost())
}

func main() {
	sword := Sword{name: "Icy-sword"}
	bow := Bow{name: "Fire-Bow"}

	attack(sword)
	attack(bow)
}
