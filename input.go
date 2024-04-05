package input

type person struct {
	name string
	age  int
}

func (p person) getName() string {
	return p.name
}

func (p *person) setName(name string) {
	p.name = name
}

/*
func main() {
	fmt.Println("Hello")
	person1 := person{"john", 25}
	person2 := person{"jane", 25}
	fmt.Println(person1.name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	person2.setName(name)
	person1.setName("lamba")
	userName := person2.getName()
	fmt.Println("person2 name:", userName)
	fmt.Println("person1 name:", person1.getName())
}
*/
