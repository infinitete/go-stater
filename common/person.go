package common

type Person struct {
	name string
	age  int16
}

func GetName(p *Person) string {
	return p.name
}

func MakePerson(name string, age int16) Person {
	return Person{
		name: name,
		age:  age,
	}
}
