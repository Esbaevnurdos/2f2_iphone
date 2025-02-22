package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// iPhone struct
type iPhone struct {
	ID    int    `json:"id"`
	Model string `json:"model"`
	Color string `json:"description"`
	Price string `json:"price"`
}

// iPhone data with IDs
var iPhones = []iPhone{
	{1, "Apple iPhone 14 Pro Max", "128GB Deep Purple", "$900"},
	{2, "Blackmagic Pocket Cinema", "Camera 6k", "$2535"},
	{3, "Apple Watch Series 9 GPS", "41mm Starlight Aluminium", "$399"},
	{4, "AirPods Max Silver", "Starlight Aluminium ", "$549"},
	{5, "Samsung Galaxy Watch6", "Classic 47mm Black", "$369"},
	{6, "Galaxy Z Fold5 Unlocked", "256GB | Phantom Black", "$1799"},
	{7, "Galaxy Buds FE", "Graphite", "$99.99"},
	{8, "Apple iPad 9 10.2 64GB Wi-Fi", "Silver (MK2L3) 2021", "$398"},
	{9, "Apple iPhone 14 Pro 512GB", "Gold (MQ233)", "$1437"},
	{10, "AirPods Max Silver", "Starlight Aluminium ", "$549"},
	{11, "Apple Watch Series 9 GPS", "41mm Starlight Aluminium", "$399"},
	{12, "Apple iPhone 14 Pro 1TB", "Gold (MQ2V3)", "$1499"},
	{13, "Apple iPhone 14 Pro", "512GB Gold (MQ233)", "$1437"},
	{14, "Apple iPhone 11", "White (MQ233)", "$550"},
	{15, "Apple iPhone 14 Pro 1TB", "Gold (MQ2V3)", "$1499"},
	{16, "Apple iPhone 14 Pro 1TB", "Gold (MQ2V3)", "$1399"},
	{17, "Apple iPhone 14 Pro", "128GB Deep Purple (MQ0G3)", "$1600"},
	{18, "Apple iPhone 13 mini", "128GB Pink (MLK23)", "$850"},
	{19, "Apple iPhone 14 Pro", "256GB Space Black (MQ0T3)", "$3499"},
	{20, "Apple iPhone 14 Pro", "256GB Silver (MQ103)", "$1399"},
}

// Enable CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// Get all iPhones
func getIPhones(c *gin.Context) {
	c.JSON(http.StatusOK, iPhones)
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/iphones", getIPhones)

	r.Run(":8080") // Run server on port 8080
}
