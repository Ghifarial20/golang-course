package main

import(
	"fmt"
	"sort"
)
	


func main() {
	fruits := []string{"apel","mangga","apel","jeruk","mangga","jeruk","apel","mangga","mangga"}
	// Urutkan dari buah yang terbanyak
	// mangga 4
	// apel 3
	// jeruk 2

	mapFruit := make(map[string]int)
	for _,num := range fruits {
		mapFruit[num] = mapFruit[num]+1
	}

	values := make([]string, 0, len(mapFruit))
  
    for key := range mapFruit {
        values = append(values, key)
    }
  
    sort.SliceStable(values, func(i, j int) bool{
        return mapFruit[values[i]] > mapFruit[values[j]]
    })
  
    for _, k := range values{
        fmt.Println(k, mapFruit[k])
    }
}