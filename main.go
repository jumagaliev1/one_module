package onemodule

import "fmt"

type User struct {
	Name string
	Age  int
}

type transport interface {
	whoOwner()
	getHealth()
}

type Car struct {
	Name   string
	Health int
	Owner  User
}
type Tank struct {
	Name   string
	Health int
	Owner  User
}

func (t *Tank) Fire(c *Car) {
	c.Health--
}

func (t *Tank) whoOwner() {
	fmt.Println("Owner:", t.Owner.Name)
}

func (c *Car) Drive() {
	fmt.Println("Car moved away Tank")
}

func (c *Car) whoOwner() {
	fmt.Println("Owner:", c.Owner.Name)
}

func (c *Car) getHealth() {
	fmt.Println("Health: ", c.Health)
}

func (t *Tank) getHealth() {
	fmt.Println("Health: ", t.Health)
}

func PrintOwner(t transport) {
	t.whoOwner()
}

func PrintHealth(t transport) {
	t.getHealth()
}
