package main

import (
	"fmt"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*type form interface {
	func () Name() form {
		return f.name
	}
}*/

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	var form Login

	router.GET("/login", GetLogin)
	router.POST("/loginJSON", PostJSONForm)
	router.POST("/loginXML", PostXMLForm)
	router.POST("/loginForm", PostHTMLForm(form))

	/*
		sudo lsof -n -i :8080
		kill -9 <PID>
	*/
	router.Run(":8080")
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

// Example for binding JSON ({"user": "manu", "password": "123"})
// curl -v -X POST http://localhost:8080/loginJSON -H 'content-type: application/json' '{ "user": "manu", "password"="123" }'
func PostJSONForm(c *gin.Context) {
	//var json Login
	var json interface{}
	//var form map[string]interface{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	*/

	c.JSON(http.StatusOK, "json")
	c.JSON(http.StatusOK, json)
}

// Example for binding XML (
//	<?xml version="1.0" encoding="UTF-8"?>
//	<root>
//		<user>user</user>
//		<password>123</password>
//	</root>)
// curl -v -X POST http://localhost:8080/loginXML -H 'content-type: application/json' -d '{ "user": "manu", "password"="123" }'
func PostXMLForm(c *gin.Context) {
	//var xml Login
	var xml interface{}
	//var form map[string]interface{}

	if err := c.ShouldBindXML(&xml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	*/

	c.JSON(http.StatusOK, "xml")
	c.JSON(http.StatusOK, xml)
}

// Example for binding a HTML form (user=manu&password=123)
// curl -v -X POST http://localhost:8080/loginForm -H 'content-type: application/json' -d '{ "user": "manu", "password":"123" }'
func PostHTMLForm2(c *gin.Context, form map[string]interface{}) {
	//var form Login
	//var form interface{}
	//var form map[string]interface{}

	form = make(map[string]interface{})

	/*t := time.Now()
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	fmt.Printf("Form: %+v\n", c.GetRawData)*/

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	*/

	c.JSON(http.StatusOK, "html")
	c.JSON(http.StatusOK, form)
	//c.JSON(http.StatusOK, c)

	/*body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	c.JSON(http.StatusOK, string(x))
	fmt.Printf("lalala: %s \n", string(x))*/

}

func PostHTMLForm(form interface{}) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		/*
			if form.User != "manu" || form.Password != "123" {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		*/

		fmt.Println(form)
		c.String(http.StatusOK, "HTML \n")
		c.JSON(http.StatusOK, form)
		c.String(http.StatusOK, "\n")
	}

	return gin.HandlerFunc(fn)
}
