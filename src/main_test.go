package main

import (
	"encoding/json"
	"fmt"
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
	mockResponse1 := map[string]interface{}{
		"sections": map[string]interface{}{
			"2103106_3": map[string]interface{}{
				"node": map[string]interface{}{
					"classes": []interface{}{
						map[string]interface{}{
							"credit":    1,
							"day":       "อังคาร",
							"format":    "บรรยาย",
							"location":  "Online",
							"timeEnd":   "14:00",
							"timeStart": "13:00",
						},
						map[string]interface{}{
							"credit":    2,
							"day":       "อังคาร",
							"format":    "ปฏิบัติ",
							"location":  "Online",
							"timeEnd":   "18:00",
							"timeStart": "14:00",
						},
					},
					"code":         "2103106",
					"final":        "3 ธ.ค. 64 เวลา 08:30-11:30",
					"index":        2,
					"midterm":      "1 ต.ค. 64 เวลา 08:30-10:30",
					"name":         "ENG DRAWING",
					"section":      3,
					"studentCount": 75,
				},
				"Participants": []interface{}{
					"นายศุภครินทร์  เนียรสุพรพรรณ",
				},
				"studentsJson": nil,
			},
			"2301107_6": map[string]interface{}{
				"node": map[string]interface{}{
					"classes": []interface{}{
						map[string]interface{}{
							"credit":    3,
							"day":       "จันทร์",
							"format":    "บรรยาย",
							"location":  "Online",
							"timeEnd":   "11:00",
							"timeStart": "09:30",
						},
						map[string]interface{}{
							"credit":    3,
							"day":       "ศุกร์",
							"format":    "บรรยาย",
							"location":  "Online",
							"timeEnd":   "11:00",
							"timeStart": "09:30",
						},
					},
					"code":         "2301107",
					"final":        "1 ธ.ค. 64 เวลา 08:30-11:30 น.",
					"index":        5,
					"midterm":      "28 ก.ย. 64 เวลา 13:00-16:00 น.",
					"name":         "CALCULUS I",
					"section":      6,
					"studentCount": 150,
				},
				"Participants": []interface{}{
					"นายศุภครินทร์  เนียรสุพรพรรณ",
				},
				"studentsJson": nil,
			},
			"2302127_2": map[string]interface{}{
				"node": map[string]interface{}{
					"classes": []interface{}{
						map[string]interface{}{
							"credit":    3,
							"day":       "จันทร์",
							"format":    "บรรยาย",
							"location":  "Online",
							"timeEnd":   "12:30",
							"timeStart": "11:00",
						},
						map[string]interface{}{
							"credit":    3,
							"day":       "พุธ",
							"format":    "บรรยาย",
							"location":  "Online",
							"timeEnd":   "12:30",
							"timeStart": "11:00",
						},
					},
					"code":         "2302127",
					"final":        "2 ธ.ค. 64 เวลา 08:30-11:30 น.",
					"index":        6,
					"midterm":      "30 ก.ย. 64 เวลา 13:00-16:00 น.",
					"name":         "GEN CHEM",
					"section":      2,
					"studentCount": 280,
				},
				"Participants": []interface{}{
					"นายศุภครินทร์  เนียรสุพรพรรณ",
				},
				"studentsJson": nil,
			},
			"2302163_4": map[string]interface{}{
				"node": map[string]interface{}{
					"classes": []interface{}{
						map[string]interface{}{
							"credit":    1,
							"day":       "อังคาร",
							"format":    "ปฏิบัติ",
							"location":  "MHMK 703",
							"timeEnd":   "12:00",
							"timeStart": "09:00",
						},
					},
					"code":         "2302163",
					"final":        "รอประกาศ",
					"index":        7,
					"midterm":      "1 ต.ค. 64 เวลา 13:00-15:00 น.",
					"name":         "GEN CHEM LAB",
					"section":      4,
					"studentCount": 60,
				},
				"Participants": []interface{}{
					"นายศุภครินทร์  เนียรสุพรพรรณ",
				},
				"studentsJson": nil,
			},
			"2304103_6": map[string]interface{}{
				"node": map[string]interface{}{
					"classes": []interface{}{
						map[string]interface{}{
							"credit":    3,
							"day":       "พุธ",
							"format":    "บรรยาย",
							"location":  "Online",
							"timeEnd":   "11:00",
							"timeStart": "09:30",
						},
						map[string]interface{}{
							"credit":    3,
							"day":       "ศุกร์",
							"format":    "บรรยาย",
							"location":  "Online",
							"timeEnd":   "09:30",
							"timeStart": "08:00",
						},
					},
					"code":         "2304103",
					"final":        "30 พ.ย. 64 เวลา 08:30-11:30 น.",
					"index":        8,
					"midterm":      "27 ก.ย. 64 เวลา 13:00-16:00 น.",
					"name":         "GEN PHYS I",
					"section":      6,
					"studentCount": 300,
				},
				"Participants": []interface{}{
					"นายศุภครินทร์  เนียรสุพรพรรณ",
				},
				"studentsJson": nil,
			},
			"2304183_6": map[string]interface{}{
				"node": map[string]interface{}{
					"classes": []interface{}{
						map[string]interface{}{
							"credit":    1,
							"day":       "จันทร์",
							"format":    "ปฏิบัติ",
							"location":  "MHMK 501",
							"timeEnd":   "16:00",
							"timeStart": "13:00",
						},
					},
					"code":         "2304183",
					"final":        "7 ธ.ค. 64 เวลา 13:00-15:00 น.",
					"index":        9,
					"midterm":      "รอประกาศ",
					"name":         "GEN PHYS LAB I",
					"section":      6,
					"studentCount": 175,
				},
				"Participants": []interface{}{
					"นายศุภครินทร์  เนียรสุพรพรรณ",
				},
				"studentsJson": nil,
			},
			"5500111_11": map[string]interface{}{
				"node": map[string]interface{}{
					"classes": []interface{}{
						map[string]interface{}{
							"credit":    3,
							"day":       "พุธ",
							"format":    "บรรยาย/ปฏิบัติ",
							"location":  "Online",
							"timeEnd":   "16:00",
							"timeStart": "13:00",
						},
					},
					"code":         "5500111",
					"final":        "รอประกาศ",
					"index":        10,
					"midterm":      "รอประกาศ",
					"name":         "EXP ENG 1",
					"section":      11,
					"studentCount": 35,
				},
				"Participants": []interface{}{
					"นายศุภครินทร์  เนียรสุพรพรรณ",
				},
				"studentsJson": nil,
			},
		},
		"students": []interface{}{
			map[string]interface{}{
				"group":               46,
				"groupName":           "วิศวกรรมศาสตร์ทั่วไป",
				"index":               670,
				"meetingId":           "947 6959 3055",
				"meetingLink":         "https://chula.zoom.us/j/94769593055?pwd=cFdFWWJaQnFjcS9HelFOMmpIa3AwUT09",
				"meetingPassword":     "834643",
				"name":                "นายศุภครินทร์  เนียรสุพรพรรณ",
				"teacher":             "ผู้ช่วยศาสตราจารย์ ดร.โศรดา กนกพานนท์",
				"teacherAssistant":    "ภูมินทร์ ภัยวงศ์",
				"teacherAssistantTel": "062-091-4500",
				"teacherEmail":        "",
				"uid":                 "6430385121",
			},
		},
	}

	mockResponse2 := map[string]interface{}{
		"error": "too many ids",
	}

	matcherService := services.NewMatcherService()
	matcherController := controllers.NewMatcherController(matcherService)
	err := os.Setenv("DATA_URL", "https://esc.eng.chula.ac.th/or64/page-data/detail/%s/page-data.json")
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.GET("/", matcherController.GetMatchSections)

	testCases := []struct {
		Method   string
		Url      string
		Response map[string]interface{}
		Status   int
	}{
		{
			Method:   "GET",
			Url:      "/?ids=6430385121",
			Response: mockResponse1,
			Status:   http.StatusOK,
		},
		{
			Method:   "GET",
			Url:      "/?ids=6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121,6430385121",
			Response: mockResponse2,
			Status:   http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		req, _ := http.NewRequest(testCase.Method, testCase.Url, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responseData, _ := ioutil.ReadAll(w.Body)
		var response map[string]interface{}
		if err := json.Unmarshal(responseData, &response); err != nil {
			panic(err)
		}
		if fmt.Sprint(response) != fmt.Sprint(testCase.Response) {
			t.Errorf("Expected %v, got %v", testCase.Response, response)
		}
		assert.Equal(t, testCase.Status, w.Code)
	}

}
