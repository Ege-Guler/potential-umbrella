package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type device struct {
	IP  string `json:"ip"`
	MAC string `json:"mac"`
}

func getDevices(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, devices)
}

var devices []device

func ReadByLine(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.Join(strings.Fields(s), " ")
		z := strings.Split(s, " ")

		devices = append(devices, device{z[1], z[0]})
	}
}

func main() {

	const filename = "ethX.dat"
	ReadByLine(filename)
	fmt.Println(devices)
	router := gin.Default()
	router.GET("/devices", getDevices)
	router.Run("localhost:9090")

}
