package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func info(c *gin.Context) {
	file := "GeoIP/GeoLite2-City.mmdb"

	db, err := geoip2.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ip := net.IP(c.Request.RemoteAddr)
	ip := net.ParseIP("73.200.70.234")
	fmt.Println(ip)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	// address := gin.H{"ip": c.Request.RemoteAddr, "country": record.Country.Names, "location": record.Location, "timezone": record.Location.TimeZone}
	address := gin.H{"ip": ip, "details": record}
	c.JSON(200, address)
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/info", info)
	}

	router.Run(":8080")
}
