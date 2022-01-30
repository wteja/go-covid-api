package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wteja/go-covid-api/data"
)

func fetchDataSource() (*data.CovidDataList, error) {
	res, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var jsonData data.CovidDataList
	jsonError := json.Unmarshal(body, &jsonData)
	if jsonError != nil {
		return nil, err
	}

	return &jsonData, nil
}

func GetCovidSummary(c *gin.Context) {
	covidData, err := fetchDataSource()
	if err != nil {
		c.Error(err)
	}
	provinces, ages := covidData.Group()
	c.JSON(http.StatusOK, gin.H{
		"Province": provinces,
		"AgeGroup": ages,
	})
}

func CreateServer() *gin.Engine {
	r := gin.Default()
	r.GET("covid/summary", GetCovidSummary)
	return r
}
