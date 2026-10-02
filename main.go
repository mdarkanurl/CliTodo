package main

func main() {
	todos := Todos{}
	todos.add("Buy milk!")
	todos.add("Get ready!")
	todos.toggle(0)
	todos.print()
}
