package main

import (
	"bufio"
	"fmt"
	"os"
)

func main3() {
	var persons = make(map[int]Person)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("CRUD Operations")

	fmt.Println("1. Create Person")
	fmt.Println("2. Read Persons")
	fmt.Println("3. Update Person")
	fmt.Println("4. Delete Person")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")

	for {
		switch input, _, _ := reader.ReadRune(); input {
		case '1':
			fmt.Println("Creating Person")

			name, age, email := "", 0, ""

			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Age: ")
			fmt.Scan(&age)
			fmt.Print("Enter Email: ")
			fmt.Scan(&email)

			person := Person{len(persons) + 1, name, age, email}

			createPerson(person, persons)
		case '2':
			fmt.Println("Reading Persons")
			persons := readPersons(persons)
			for _, person := range persons {
				fmt.Println(person)
			}
		case '3':
			fmt.Println("Updating Person")
			fmt.Print("Enter ID: ")
			id := 0
			fmt.Scan(&id)
			for _, person := range persons {
				if person.id == id {
					newName, newAge, newEmail := "", 0, ""

					fmt.Print("Enter Name: ")
					fmt.Scan(&newName)
					fmt.Print("Enter Age: ")
					fmt.Scan(&newAge)
					fmt.Print("Enter Email: ")
					fmt.Scan(&newEmail)

					person := Person{id, newName, newAge, newEmail}
					updatePerson(id, persons, person)
				}
			}
		case '4':
			fmt.Println("Deleting Person")
			fmt.Print("Enter ID: ")
			id := 0
			fmt.Scan(&id)

			deletePerson(id, persons)
		case '0':
			fmt.Println("Exiting")
			os.Exit(0)
		}
	}

}

func createPerson(person Person, persons map[int]Person) {
	persons[person.id] = person
}

func readPersons(persons map[int]Person) map[int]Person {
	return persons
}

func updatePerson(id int, persons map[int]Person, newPerson Person) {
	for i, person := range persons {
		if person.id == id {
			persons[i] = newPerson
		}
	}
}

func deletePerson(id int, persons map[int]Person) {
	for i, person := range persons {
		if person.id == id {
			delete(persons, i)
		}
	}
}

type Person struct {
	id    int
	name  string
	age   int
	email string
}
