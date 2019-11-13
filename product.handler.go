package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	products := getAllProducts()

	render(c, gin.H{
		"title":   "Agritek Gin",
		"payload": products}, "index.html")
}

func showProductCreationPage(c *gin.Context) {

	render(c, gin.H{
		"title": "Tambah Produk Baru"}, "tambah-produk.html")
}

func getProduct(c *gin.Context) {

	if productID, err := strconv.Atoi(c.Param("product_id")); err == nil {

		if product, err := getProductByID(productID); err == nil {

			render(c, gin.H{
				"title":   product.Title,
				"payload": product}, "produk.html")

		} else {

			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {

		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createProduct(c *gin.Context) {

	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewProduct(title, content); err == nil {

		render(c, gin.H{
			"title":   "Sukses Menambahkan",
			"payload": a}, "sukses-menambahkan.html")
	} else {

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
