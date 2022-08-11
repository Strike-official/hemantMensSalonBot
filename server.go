package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Strike-official/hemantMensSalonBot/configmanager"
	"github.com/Strike-official/hemantMensSalonBot/internal/model"
	"github.com/gin-gonic/gin"
)

func main() {
	// Read Config
	err := configmanager.InitAppConfig("configs/config.json")
	if err != nil {
		log.Fatal("[startAPIs] Failed to start APIs. Error: ", err)
	}
	model.Conf = configmanager.GetAppConfig()

	// Init LogFile
	logFile := initLogger(model.Conf.LogFilePath)

	// Init DB Connection
	// mysql.ConnectToRDS()
	// defer mysql.ConnClose()

	// Init slotDetails
	model.SlotDetail = make(map[string]map[string]map[int]model.SlotTime)
	model.SalonOpeningTime, _ = time.ParseInLocation(model.TimeLatout, "2022-08-10 09:00:00", time.Local)
	model.SalonClosingTime, _ = time.ParseInLocation(model.TimeLatout, "2022-08-10 21:00:00", time.Local)
	fmt.Printf("Salon OpeningTime = %+v and ClosingTime= %+v", model.SalonOpeningTime, model.SalonClosingTime)

	// Init Routes
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	router := gin.Default()
	initializeRoutes(router)

	// Start serving the application
	err = router.Run(model.Conf.Port)
	if err != nil {
		log.Fatal("[startAPIs] Failed to start APIs. Error: ", err)
	}
}

func initLogger(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return file
}
