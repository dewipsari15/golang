package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Nama string `json:"nama" binding:"required" gorm:"column:nama"`
	Umur int    `json:"umur" binding:"required"`
}

func (x *User) TableName() string {
	return "user"
}
func PrintKata(c *gin.Context) {
	c.JSON(200, "hello world")
}

func PrintUser(c *gin.Context) {
	var orang User
	err := c.BindJSON(&orang)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "request invalid")
		return
	}
	c.JSON(200, orang)
}

func SaveUser(c *gin.Context) {
	var orang User
	err := c.BindJSON(&orang)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "request invalid")
		return
	}
	dsn := "root:@tcp(localhost:3306)/hicolleagues"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&orang)

	err = db.Create(&orang).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "failed save data")
		return
	}

	c.JSON(200, orang)
}

//	func PrintText(w http.ResponseWriter, r *http.Request) {
//		text := "hello world"
//		temp, _ := json.Marshal(text)
//		w.Write(temp)
//	}
func main() {

	// http.HandleFunc("/", PrintText)
	// //running di : localhost:9000
	// log.Fatal(http.ListenAndServe(":9000", nil))

	app := gin.Default()
	app.GET("/", PrintKata)
	app.POST("/user", PrintUser)
	app.POST("/user/save", SaveUser)

	app.Run(":9000")
}
