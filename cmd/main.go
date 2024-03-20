package main

func main() {
	if err := commands.Execute(); err != nil {
		panic(err)
	}
}
