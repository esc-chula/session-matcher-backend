package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"matech-backend/src/models"
	"net/http"
	"os"
	"sync"
)

type MatcherService interface {
	GetStudentsSecitons(ids []string, result *[]models.StudentSections) error
	MatchSections(rawSections []models.StudentSections, section *map[string]models.Section)
}

type matcherService struct {
}

func NewMatcherService() MatcherService {
	return &matcherService{}
}

func getSections(id string, data *models.StudentSections) error {
	reqUrl := fmt.Sprintf(os.Getenv("DATA_URL"), id)
	response, err := http.Get(reqUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	tmpData := models.StudentSectionsResponse{}
	err = json.Unmarshal(responseBody, &tmpData)
	if err != nil {
		return err
	}

	data.AllSectionsJSON = tmpData.Result.Data.AllSectionsJSON
	data.StudentsJSON = tmpData.Result.Data.StudentsJSON
	return nil
}

func (m *matcherService) GetStudentsSecitons(ids []string, result *[]models.StudentSections) error {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			select {
			case <-ctx.Done():
				return
			default:
				data := models.StudentSections{}
				if err := getSections(id, &data); err != nil {
					cancel()
					wg.Done()
					return
				}
				*result = append(*result, data)
				wg.Done()
			}
		}(id)
	}
	wg.Wait()
	return ctx.Err()
}

func (m *matcherService) MatchSections(rawSections []models.StudentSections, section *map[string]models.Section) {
	for _, sectionData := range rawSections {
		for _, class := range sectionData.AllSectionsJSON.Edges {
			code := fmt.Sprintf("%s_%d", class.Node.Code, class.Node.Section)
			if _, ok := (*section)[code]; !ok {
				tmp := models.Section{}
				tmp.Node = class.Node
				(*section)[code] = tmp
			}
			tmpS := (*section)[code]
			tmpS.Participants = append(tmpS.Participants, sectionData.StudentsJSON.Name)
			(*section)[code] = tmpS
		}
	}
}
