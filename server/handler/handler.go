package handler

import (
	"fmt"
	"html"
	"net/http"

	"github.com/garciasa/machinedirectory/server/storage"
	"github.com/garciasa/machinedirectory/server/storage/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"response,omitempty"`
}

type saveCommand struct {
	IP         string `json:"ip"`
	DomainName string `json:"domainname"`
	Tags       string `json:"tags"`
}

func dBMiddleware(d database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Set("dbConn", d.Db)
		c.Next()
	}
}

// New Create gin engine
func New(d database.Database) *gin.Engine {
	r := gin.Default()
	r.Use(dBMiddleware(d))
	r.GET("/items", getAllItems)
	r.GET("/item/:id", getItem)
	r.GET("items/:tags", searchByTags)
	r.POST("/item/", createItem)
	r.PUT("/item/:id", updateItem)
	r.DELETE("/item/:id", deleteItem)

	return r
}

func getAllItems(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		c.AbortWithStatus(505)
		return
	}

	var items []storage.Item
	if err := db.Find(&items).Error; err != nil {
		c.JSON(http.StatusOK, &response{Success: false})
		return
	}
	c.JSON(http.StatusOK, &response{Success: true, Data: items})
}
func getItem(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		c.AbortWithStatus(505)
		return
	}

	id := html.EscapeString(c.Params.ByName("id"))
	var item storage.Item
	if err := db.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusOK, &response{Success: false})
		return
	}
	c.JSON(http.StatusOK, &response{Success: true, Data: item})
}

func createItem(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		c.AbortWithStatus(505)
		return
	}

	var sc saveCommand
	c.BindJSON(&sc)
	item := storage.Item{
		IP:         html.EscapeString(sc.IP),
		DomainName: html.EscapeString(sc.DomainName),
		Tags:       html.EscapeString(sc.Tags),
		Deleted:    false,
	}

	if err := db.Save(&item).Error; err != nil {
		fmt.Println(err.Error())
		c.JSON(200, &response{Success: false, Error: "Something was wrong :("})

		return
	}

	c.JSON(200, &response{Success: true, Data: &item})

}

func searchByTags(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		c.AbortWithStatus(505)
		return
	}
	tags := html.EscapeString(c.Params.ByName("tags"))

	var items []storage.Item
	err := db.Where("tags LIKE ? or domain_name LIKE ?", "%"+tags+"%", "%"+tags+"%").Find(&items).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(200, &response{Success: false, Error: "Something was wrong :("})

		return
	}

	c.JSON(200, &response{Success: true, Data: &items})

}

func deleteItem(c *gin.Context) {}
func updateItem(c *gin.Context) {}
