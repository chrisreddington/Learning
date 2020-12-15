package main

func main() {
	/*var i int
	for i < 5 {
		println(i)
		i++

		if i == 3 {
			continue
		}
		println("continuing...")
	}*/

	/*var i int
	for ; i < 5; i++ {
		println(i)
	}

	println(i)*/

	/*var i int
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}*/

	//Ugly way to access collections
	/*slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}*/

	//Clean way to access collections
	// e.g. slice
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	// e.g. map
	wellKnownPorts := map[string]int{"http": 80, "https": 43}
	for k, v := range wellKnownPorts {
		println(k, v)
	}

	// e.g. map keys
	for k := range wellKnownPorts {
		println(k)
	}

	// e.g. map values
	for _, v := range wellKnownPorts {
		println(v)
	}
}
