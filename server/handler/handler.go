package handler

import (
	"fmt"
	"machinedirectory/server/storage"
	"machinedirectory/server/storage/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"response,omitempty"`
}

func dBMiddleware(d database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		c.JSON(200, &response{Success: false})
		return
	}
	c.JSON(200, &response{Success: true, Data: items})
}
func getItem(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		c.AbortWithStatus(505)
		return
	}

	id := c.Params.ByName("id")
	var item storage.Item
	if err := db.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(200, &response{Success: false})
		return
	}
	c.JSON(200, &response{Success: true, Data: item})
}

func createItem(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		c.AbortWithStatus(505)
		return
	}
	item := storage.Item{
		IP:         c.PostForm("ip"),
		DomainName: c.PostForm("domainname"),
		Tags:       c.PostForm("tags"),
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
	tags := c.Params.ByName("tags")

	var items []storage.Item
	err := db.Where("tags LIKE ?", "%"+tags+"%").Find(&items).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(200, &response{Success: false, Error: "Something was wrong :("})

		return
	}

	c.JSON(200, &response{Success: true, Data: &items})

}

func deleteItem(c *gin.Context) {}
func updateItem(c *gin.Context) {}
