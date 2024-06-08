package api

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	// 允许跨域请求
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	user := r.Group("/user")
	{
		user.POST("/login", login)
	}

	cargo := r.Group("/cargo")
	{
		addCargo := cargo.Group("/add")
		{
			addCargo.POST("/product", addProduct)
			addCargo.POST("/brand", addBrand)
			addCargo.POST("/type", addProductType)
			addCargo.POST("/vendor", addVendor)
		}

		ordering := cargo.Group("/ordering")
		{
			ordering.GET("/get", getAutoOrderList)
			ordering.POST("/add", addAutoOrder)
			ordering.POST("/stop", stopAutoOrder)
		}

		cargo.POST("/restock", addCargoStock)
	}

	customer := r.Group("/customer")
	{
		customer.POST("/new", newMembership)
		purchase := customer.Group("/purchase")
		{
			purchase.POST("/begin", StartPurchase)
			purchase.POST("/add", AddItem)
			purchase.POST("/end", FinishPurchase)
		}
	}

	//根据name查询
	query := r.Group("/query")
	{
		query.POST("/brand", brandQuery)
		query.POST("/type", typeQuery)
		query.POST("/vendor", vendorQuery)
		query.POST("/store", storeQuery)
		query.POST("/upc", upcQuery)

		query.POST("/products", getProductsByStoreID)

		query.GET("/stores", getStores)
	}

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile("./dist/index.html")
			if err != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		}
	})

	r.Run()
}
