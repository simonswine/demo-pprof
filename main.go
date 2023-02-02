package main

func main() {
	// spend 3 cpu cycles
	doALot()
	doLittle()
}

func prepare() {
	// spend 5 cpu cycles
}

func doALot() {
	prepare()
	// spend 20 cpu cycles
}

func doLittle() {
	prepare()
	// spend 5 cpu cycles
}
