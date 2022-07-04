package main

import (
	"io/ioutil"
	"matech-backend/src/controllers"
	"matech-backend/src/services"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetMatchSection(t *testing.T) {
	mockResponse := `{"2103106_3":{"node":{"classes":[{"credit":1,"day":"อังคาร","format":"บรรยาย","location":"Online","timeEnd":"14:00","timeStart":"13:00"},{"credit":2,"day":"อังคาร","format":"ปฏิบัติ","location":"Online","timeEnd":"18:00","timeStart":"14:00"}],"code":"2103106","final":"3 ธ.ค. 64 เวลา 08:30-11:30","index":2,"midterm":"1 ต.ค. 64 เวลา 08:30-10:30","name":"ENG DRAWING","section":3,"studentCount":75},"Participants":["นายศุภครินทร์  เนียรสุพรพรรณ"]},"2301107_6":{"node":{"classes":[{"credit":3,"day":"จันทร์","format":"บรรยาย","location":"Online","timeEnd":"11:00","timeStart":"09:30"},{"credit":3,"day":"ศุกร์","format":"บรรยาย","location":"Online","timeEnd":"11:00","timeStart":"09:30"}],"code":"2301107","final":"1 ธ.ค. 64 เวลา 08:30-11:30 น.","index":5,"midterm":"28 ก.ย. 64 เวลา 13:00-16:00 น.","name":"CALCULUS I","section":6,"studentCount":150},"Participants":["นายศุภครินทร์  เนียรสุพรพรรณ"]},"2302127_2":{"node":{"classes":[{"credit":3,"day":"จันทร์","format":"บรรยาย","location":"Online","timeEnd":"12:30","timeStart":"11:00"},{"credit":3,"day":"พุธ","format":"บรรยาย","location":"Online","timeEnd":"12:30","timeStart":"11:00"}],"code":"2302127","final":"2 ธ.ค. 64 เวลา 08:30-11:30 น.","index":6,"midterm":"30 ก.ย. 64 เวลา 13:00-16:00 น.","name":"GEN CHEM","section":2,"studentCount":280},"Participants":["นายศุภครินทร์  เนียรสุพรพรรณ"]},"2302163_4":{"node":{"classes":[{"credit":1,"day":"อังคาร","format":"ปฏิบัติ","location":"MHMK 703","timeEnd":"12:00","timeStart":"09:00"}],"code":"2302163","final":"รอประกาศ","index":7,"midterm":"1 ต.ค. 64 เวลา 13:00-15:00 น.","name":"GEN CHEM LAB","section":4,"studentCount":60},"Participants":["นายศุภครินทร์  เนียรสุพรพรรณ"]},"2304103_6":{"node":{"classes":[{"credit":3,"day":"พุธ","format":"บรรยาย","location":"Online","timeEnd":"11:00","timeStart":"09:30"},{"credit":3,"day":"ศุกร์","format":"บรรยาย","location":"Online","timeEnd":"09:30","timeStart":"08:00"}],"code":"2304103","final":"30 พ.ย. 64 เวลา 08:30-11:30 น.","index":8,"midterm":"27 ก.ย. 64 เวลา 13:00-16:00 น.","name":"GEN PHYS I","section":6,"studentCount":300},"Participants":["นายศุภครินทร์  เนียรสุพรพรรณ"]},"2304183_6":{"node":{"classes":[{"credit":1,"day":"จันทร์","format":"ปฏิบัติ","location":"MHMK 501","timeEnd":"16:00","timeStart":"13:00"}],"code":"2304183","final":"7 ธ.ค. 64 เวลา 13:00-15:00 น.","index":9,"midterm":"รอประกาศ","name":"GEN PHYS LAB I","section":6,"studentCount":175},"Participants":["นายศุภครินทร์  เนียรสุพรพรรณ"]},"5500111_11":{"node":{"classes":[{"credit":3,"day":"พุธ","format":"บรรยาย/ปฏิบัติ","location":"Online","timeEnd":"16:00","timeStart":"13:00"}],"code":"5500111","final":"รอประกาศ","index":10,"midterm":"รอประกาศ","name":"EXP ENG 1","section":11,"studentCount":35},"Participants":["นายศุภครินทร์  เนียรสุพรพรรณ"]}}`
	matcherService := services.NewMatcherService()
	matcherController := controllers.NewMatcherController(matcherService)
	err := os.Setenv("DATA_URL", "https://esc.eng.chula.ac.th/or64/page-data/detail/%s/page-data.json")
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.GET("/", matcherController.GetMatchSections)
	req, _ := http.NewRequest("GET", "/?ids=6430385121", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
