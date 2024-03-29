package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {

	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func performLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if isUserValid(username, password) {

		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Berhasil Login"}, "login-berhasil.html")

	} else {

		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Gagal",
			"ErrorMessage": "Kredensial tidak sesuai"})
	}
}

func generateSessionToken() string {

	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func showRegistrationPage(c *gin.Context) {

	render(c, gin.H{
		"title": "Daftar"}, "daftar.html")
}

func register(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {

		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Daftar & Login Sukses"}, "login-berhasil.html")

	} else {

		c.HTML(http.StatusBadRequest, "daftar.html", gin.H{
			"ErrorTitle":   "Pendaftaran Gagal",
			"ErrorMessage": err.Error()})

	}
}
