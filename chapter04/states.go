package main

import "fmt"

type State struct {
	name       string
	population int
	capital    string
}

func main() {
	states := make(map[string]State, 6)

	states["RS"] = State{"Rio grande do Sul", 1129000, "Porto Alegre"}
	states["GO"] = State{"Goiás", 6434052, "Goiânia"}
	states["PR"] = State{"Paraná", 10997462, "Curitiba"}
	states["RN"] = State{"Rio Grande do Norte", 3373960, "Natal"}
	states["AM"] = State{"Amazonas", 3807923, "Manaus"}
	states["SE"] = State{"Sergipe", 2228489, "Aracaju"}

	fmt.Println(states)

	// Lookup

	lookUpState := states["RS"]

	fmt.Println(lookUpState)

	// Lookup undefined but avoid zero value { 0 }
	lookUpState, wasFound := states["SP"]
	if wasFound {
		fmt.Println(lookUpState)
	}

	delete(states, "AM")

	for uf, state := range states {
		fmt.Printf("%s (%s) has %d inhabitants \n", state.name, uf, state.population)
	}
}
