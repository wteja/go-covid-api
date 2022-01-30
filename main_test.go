package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/wteja/go-covid-api/data"
	"github.com/wteja/go-covid-api/server"
)

func getMockData() data.CovidDataList {

	age1 := 51
	provinceId1 := 46
	statQuarantine1 := 5
	age2 := 51
	provinceId2 := 65
	statQuarantine2 := 8
	age3 := 79
	provinceId3 := 53
	statQuarantine3 := 1

	list := []data.CovidData{
		{
			No:             nil,
			ConfirmDate:    "2021-05-04",
			Age:            &age1,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "China",
			Province:       "Phrae",
			ProvinceId:     &provinceId1,
			District:       "",
			ProvinceEn:     "Phrae",
			StatQuarantine: &statQuarantine1,
		},
		{
			No:             nil,
			ConfirmDate:    "2021-05-01",
			Age:            &age2,
			Gender:         "ชาย",
			GenderEn:       "Male",
			Nation:         "",
			NationEn:       "India",
			Province:       "Suphan Buri",
			ProvinceId:     &provinceId2,
			District:       "",
			ProvinceEn:     "Suphan Buri",
			StatQuarantine: &statQuarantine2,
		},
		{
			No:             nil,
			ConfirmDate:    "2021-05-01",
			Age:            &age3,
			Gender:         "ชาย",
			GenderEn:       "Male",
			Nation:         "",
			NationEn:       "India",
			Province:       "Roi Et",
			ProvinceId:     &provinceId3,
			District:       "",
			ProvinceEn:     "Roi Et",
			StatQuarantine: &statQuarantine3,
		},
	}

	dataList := data.CovidDataList{
		Data: list,
	}

	return dataList
}

func TestGetAgeKeyYoung1(t *testing.T) {
	var age int = 0
	ageArgs := &age
	key := data.GetAgeKey(ageArgs)
	if key != "0-30" {
		t.Error("Get the wrong age key [0-30]")
	}
}

func TestGetAgeKeyYoung2(t *testing.T) {
	var age int = 15
	ageArgs := &age
	key := data.GetAgeKey(ageArgs)
	if key != "0-30" {
		t.Error("Get the wrong age key [0-30]")
	}
}

func TestGetAgeKeyYoung3(t *testing.T) {
	var age int = 30
	ageArgs := &age
	key := data.GetAgeKey(ageArgs)
	if key != "0-30" {
		t.Error("Get the wrong age key [0-30]")
	}
}

func TestGetAgeKeyMiddle1(t *testing.T) {
	var age int = 31
	ageArgs := &age
	key := data.GetAgeKey(ageArgs)
	if key != "31-60" {
		t.Error("Get the wrong age key [31-60]")
	}
}

func TestGetAgeKeyMiddle2(t *testing.T) {
	var age int = 45
	ageArgs := &age
	key := data.GetAgeKey(ageArgs)
	if key != "31-60" {
		t.Error("Get the wrong age key [31-60]")
	}
}

func TestGetAgeKeyMiddle3(t *testing.T) {
	var age int = 60
	ageArgs := &age
	key := data.GetAgeKey(ageArgs)
	if key != "31-60" {
		t.Error("Get the wrong age key [31-60]")
	}
}

func TestGetAgeKeyOlder1(t *testing.T) {
	var age int = 70
	ageArgs := &age
	key := data.GetAgeKey(ageArgs)
	if key != "61+" {
		t.Error("Get the wrong age key [61+]")
	}
}

func TestGetAgeKeyNull(t *testing.T) {
	var ageArgs *int = nil
	key := data.GetAgeKey(ageArgs)
	if key != "N/A" {
		t.Error("Get the wrong age key [N/A]")
	}
}

func TestProvinceNull(t *testing.T) {
	var province string = ""
	key := data.GetProvinceKey(province)
	if key != "N/A" {
		t.Error("Get the wrong province key [N/A]")
	}
}

func TestProvinceYala(t *testing.T) {
	var province string = "Yala"
	key := data.GetProvinceKey(province)
	if key != "Yala" {
		t.Error("Get the wrong province key [Yala]")
	}
}

func TestGroupData(t *testing.T) {
	dataList := getMockData()
	provinces, ages := dataList.Group()
	if provinces["Prae"] != 1 && provinces["Roi Et"] != 1 && provinces["Suphan Buri"] != 1 {
		t.Error("Province counting is invalid.")
	}
	if ages["31-60"] != 2 && ages["61+"] != 1 {
		t.Error("Age counting is invalid.")
	}
}

func TestCovidSummary(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	server.GetCovidSummary(c)
	if w.Code != 200 {
		t.Error("Covid Summary Endpoint should get status code: 200")
	}

	var list gin.H
	err := json.Unmarshal(w.Body.Bytes(), &list)
	if err != nil {
		t.Error("Covid Summary Endpoint should get response body")
	}
}
