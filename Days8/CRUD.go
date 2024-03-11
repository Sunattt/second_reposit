package Days8

import (
	"fmt"
	"log"
)

func CreateM() map[string]User {
	Jordon := User{
		Name:   "Mikle",
		Age:    34,
		Login:  true,
		Gender: "Male",
	}
	person := make(map[string]User)
	person["Jordon"] = Jordon
	//проверка насодержание ключа под именем Mihailб если нет то создаем её
	_, ok := person["Mihail"]
	if !ok {
		person["Mihail"] = User{Name: "Andrey", Age: 24, Login: true, Gender: "Male"}
	}
	human := person["Mihail"]
	fmt.Printf("%+v\n", human)
	fmt.Printf("%v\n", person)
	return person
}

func ReadM(person map[string]User) {
	fmt.Println("[key]: [value]")
	for n, k := range person {
		fmt.Printf("%v : %v\n", n, k)
	}
}

func UpdateM(person map[string]User) map[string]User{
	fmt.Println("Добавляем новое значение")
	Samual := User{
		Name: "Jeksn",
		Age: 56,
		Login: false,
		Gender: "Male",
	}
	Stef := User{
		Name: "Cary",
		Age: 40,
		Login: true,
		Gender: "Male",
	}
	Galina := User{
		Name: "Milana",
		Age: 22,
		Login: false,
		Gender: "Female",
	}
	person["Galina"] = Galina
	person["Samual"] = Samual
	person["Stef"] = Stef
	return person
}

func DeleteM(person map[string]User){
	delete(person , "Stef")
    fmt.Println()

	log.Println("Error and recover !!!")
}

