package main

//class extend slice like oop
type deck []string 
//receiver in a function
func (d deck) print()  {
	for i, card := range d {
		println(i, card)
	}
}