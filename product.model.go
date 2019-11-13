package main

import "errors"

type product struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var productList = []product{}

func getAllProducts() []product {
	return productList
}

func getProductByID(id int) (*product, error) {
	for _, a := range productList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Produk tidak ditemukan")
}

func createNewProduct(title, content string) (*product, error) {
	a := product{ID: len(productList) + 1, Title: title, Content: content}

	productList = append(productList, a)

	return &a, nil
}
