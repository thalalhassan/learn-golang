package main

func main() {

	todos := Todos{}

	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	cmd := NewCommand()
	cmd.Execute(&todos)

	storage.Save(todos)

}
