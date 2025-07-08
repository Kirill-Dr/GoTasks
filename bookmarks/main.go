package main

import "fmt"

func main() {
	m := map[string]string{
		"First": "https://first.com",
	}
	fmt.Println(m)
	fmt.Println(m["First"])
	m["First"] = "https://firstUpdated.com"
	fmt.Println(m)
	m["Second"] = "https://second.com"
	m["Third"] = "https://third.com"
	fmt.Println(m)
	delete(m, "Third")
	fmt.Println(m)
	for key, value := range m {
		fmt.Println(key, value)
	}
}
