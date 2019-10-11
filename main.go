package main

import (
	"os"

	"github.com/OkanoShogo0903/web-ai-speaker/backend/controller"
	"github.com/OkanoShogo0903/web-ai-speaker/backend/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	speech_result := model.New()
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/speech", controller.SpeechPost(speech_result))

	router.Run(":" + os.Getenv("PORT"))
}
