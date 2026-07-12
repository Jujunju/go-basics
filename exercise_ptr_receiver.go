// package main

// import "fmt"

// type Cart struct {
// 	ListProducts map[string]int
// }

// func main() {

// 	product1 := &Cart{
// 		ListProducts: map[string]int{
// 			"Mie Ayam Goreng":     100,
// 			"Kue Nabati 500 Gram": 50,
// 			"Laptop Mac":          10,
// 		},
// 	}

// 	product1.addProducts("Baju Polos Putih", 30)
// 	fmt.Println(product1.getSumProducts())
// 	fmt.Println(product1)

// }

// func (c *Cart) addProducts(name string, qty int) {
// 	c.ListProducts[name] = qty
// }

// func (c *Cart) getSumProducts() int {

// 	result := 0

// 	for _, q := range c.ListProducts{
// 		result += q
// 	}

// 	return result

// }
