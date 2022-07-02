package models

type StudentSectionsResponse struct {
	ComponentChunkName string `json:"componentChunkName"`
	Path               string `json:"path"`
	Result             struct {
		Data struct {
			AllGroupsJSON struct {
				Edges []struct {
					Node struct {
						GroupCode string `json:"groupCode"`
						Section   string `json:"section"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"allGroupsJson"`
			AllSectionsJSON struct {
				Edges []struct {
					Node struct {
						Classes []struct {
							Credit    int    `json:"credit"`
							Day       string `json:"day"`
							Format    string `json:"format"`
							Location  string `json:"location"`
							TimeEnd   string `json:"timeEnd"`
							TimeStart string `json:"timeStart"`
						} `json:"classes"`
						Code         string `json:"code"`
						Final        string `json:"final"`
						Index        int    `json:"index"`
						Midterm      string `json:"midterm"`
						Name         string `json:"name"`
						Section      int    `json:"section"`
						StudentCount int    `json:"studentCount"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"allSectionsJson"`
			StudentsJSON struct {
				Group               int    `json:"group"`
				GroupName           string `json:"groupName"`
				Index               int    `json:"index"`
				MeetingID           string `json:"meetingId"`
				MeetingLink         string `json:"meetingLink"`
				MeetingPassword     string `json:"meetingPassword"`
				Name                string `json:"name"`
				Teacher             string `json:"teacher"`
				TeacherAssistant    string `json:"teacherAssistant"`
				TeacherAssistantTel string `json:"teacherAssistantTel"`
				TeacherEmail        string `json:"teacherEmail"`
				UID                 string `json:"uid"`
			} `json:"studentsJson"`
		} `json:"data"`
		PageContext struct {
			Index int `json:"index"`
		} `json:"pageContext"`
	} `json:"result"`
}
