package main

import (
	"fmt"
	"math/rand"
)

func main() {
	thoroughBreds := []string{"Big Orange", "Our Ivanhowe", "Curren Mirotic", "Bondi Beach", "Exospheric", "Hartnell", "Who SHot Thebarman", "Wicklow Brave", "Almoonqith", "Gallant", "Grand Marshal", "Jameka", "Heartbreak City", "Sir John Hawkwood", "Excess Knowledge", "Beautiful Romance", "Almandin", "Assign", "Grey Lion", "Oceanographer", "Secret Number", "Pentathlon", "Qewy", "Rose of Virginia"}

	for i:= 0; i < 3; i++ {
		fmt.Println(thoroughBreds[rand.Intn(len(thoroughBreds))])
	}
}
