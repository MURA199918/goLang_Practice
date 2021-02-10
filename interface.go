package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal structure
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow type
type Cow struct {
	food, locomotion, noise string
}

//Bird type
type Bird struct {
	food, locomotion, noise string
}

//Snake type
type Snake struct {
	food, locomotion, noise string
}

//Eat cow
func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

//Eat bird
func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

//Eat snake
func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

//Move cow
func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

//Move bird
func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

//Move snake
func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

//Speak cow
func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

//Speak bird
func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

//Speak snake
func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}

func main() {
	var nameType map[string]string
	nameType = make(map[string]string)
	for {
		fmt.Println(">Enter request: Name and information ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()
		fmt.Println("Request is: ", request)
		if len(strings.Split(request, " ")) != 3 {
			fmt.Println("Invalid command request.")
			continue
		}
		command := strings.Split(request, " ")[0]
		fmt.Println("Command: ", command)
		switch command {
		case "newanimal":
			name := strings.Split(request, " ")[1]
			animaltype := strings.Split(request, " ")[2]
			fmt.Println("Name: ", name)
			fmt.Println("Type: ", animaltype)
			_, ok := nameType[name]
			if ok == false {
				nameType[name] = animaltype
				fmt.Println("Created!!")
			} else {
				fmt.Println("Invalid animal name, try again")
			}
		case "query":
			name := strings.Split(request, " ")[1]
			info := strings.Split(request, " ")[2]
			fmt.Println("Name: ", name)
			fmt.Println("Information: ", info)
			switch nameType[name] {
			case "cow":
				cow := Cow{"grass", "walk", "moo"}
				switch info {
				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
				default:
					fmt.Println("Invalid Information, try again.")
					continue
				}
			case "bird":
				bird := Bird{"worms", "fly", "peep"}
				switch info {
				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
				default:
					fmt.Println("Invalid information, try again.")
					continue
				}
			case "snake":
				snake := Snake{"mice", "slither", "hsss"}
				switch info {
				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
				default:
					fmt.Println("Invalid information. try again.")
					continue
				}
			default:
				fmt.Println("Invalid name, try again.")
				continue
			}
		default:
			fmt.Println("Invalid query, try again!!")
			continue
		}
	}
}
