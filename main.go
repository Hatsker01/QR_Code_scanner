package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/qr_code/docs"

	"github.com/skip2/go-qrcode"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger(),
		gin.Recovery(),
	)

	api := router.Group("/v1")
	api.POST("/qrcode/:text", QrCOde)
	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	if err := router.Run(":8090"); err != nil {
		log.Fatal("failed to run http server")
		panic(err)
	}
}

// Generate QRCode
// @Summary generate qr code
// @Description This API for generating qrcode
// @Tags qrcode
// @Accept json
// @Produce json
// @Param text path string true "Enter info"
// @Router /v1/qrcode/{text} [post]
func QrCOde(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	text := c.Param("text")
	err := qrcode.WriteFile(text, qrcode.Medium, 256, "image.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("./tmp/" + filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
}
