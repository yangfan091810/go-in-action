package main

import "fmt"

func main(){
	colors := map[string]string{
		"AliceBlue":"#f0f8ff",
		"Coral":"#ff7F50",
		"DarkGray":"#a9a9a9",
		"ForestGreen":"#228b22",
	}
	for key,val := range colors {
		fmt.Printf("key %s, val %s \n", key, val)
	}
	value,exists := colors["Red"]
	if !exists {
		colors["Red"] = "red"
	} else {
		fmt.Println(value)
	}
	removeColor(colors, "Coral")
	for key,val := range colors {
		fmt.Printf("key %s, val %s \n", key, val)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
