package main

import "fmt"

func main() {
	var animal Animal
	fmt.Println("Welcome to Animal methods program.")

	for {
		fmt.Print(">Enter name and information: ")
		var animalName string
		var information string
		fmt.Scanln(&animalName)
		fmt.Scanln(&information)

		switch animalName {
		case "cow":
			animal = Animal{"grass", "walk", "moo"}
		case "bird":
			animal = Animal{"worms", "fly", "peep"}
		case "snake":
			animal = Animal{"mice", "slither", "hsss"}
		default:
			fmt.Println(animalName, " is not an animal")
			return
		}

		switch information {
		case "eat":
			fmt.Print(animalName, " eats")
			animal.Eat()
		case "move":
			fmt.Print(animalName, " eats")
			animal.Move()
		case "speak":
			fmt.Print(animalName, " eats")
			animal.Speak()
		default:
			fmt.Println(information, " is not an information")
			return
		}
	}
}

//Animal struct
type Animal struct {
	food, locomotion, noise string
}

//Eat method
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

//Move method
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

//Speak method
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}
