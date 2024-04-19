package main

import (
	"fmt"
)

func get_item(dict map[string]float64, key string) string {
	value, ok := dict[key]
	if ok {
		return fmt.Sprint(value)
	}
	return "NotFound!"
}

func print_existing_item(dict map[float64]string, key float64) {
	if name, ok := dict[key]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("FUCKOFF")
	}
}

func show_dict_like_javascript(name string, dict map[string]map[string]string) {
	fmt.Printf("%s: {\n", name)
	for k1, v1 := range dict {
		fmt.Printf("\t%s: { \n\t\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t\t%s: %s,\n", k2, v2)
		}
		fmt.Println("\n\t},")
	}
	fmt.Println("}")
}
func main() {
	var x map[string]float64 = make(map[string]float64)
	x["k1"] = 10.2
	x["k2"] = 6.66

	y := make(map[float64]string)
	y[1.2] = "one point two"
	y[6.66] = "six point sixty six"

	fmt.Println(y[6.66])

	x["to b deleted"] = 0
	fmt.Println(x)
	delete(x, "to b deleted")
	value, ok := x["item that doesnt exist"]
	fmt.Println(value, ok)

	fmt.Println(get_item(x, "k1"), get_item(x, "k3"))

	print_existing_item(y, 1.2)
	print_existing_item(y, 6.65)

	k2 := 10

	if k2 *= 3; k2 > 10 {
		fmt.Println("Correcto!")
	}

	mandaliof_table := map[string]map[string]string{
		"H": map[string]string{ // type here can be removed also, as it seen next
			"name":  "Hydrogen",
			"state": "gas",
			"Z":     "1",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
			"Z":     "2",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
			"Z":     "3",
		},
		"Be": {
			"name":  "Berrylium",
			"state": "solid",
			"Z":     "4",
		},
	}

	show_dict_like_javascript("Mandaliof", mandaliof_table)
}
