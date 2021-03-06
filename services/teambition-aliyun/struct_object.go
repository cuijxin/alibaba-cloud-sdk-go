package teambition_aliyun

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Object is a nested struct in teambition_aliyun response
type Object struct {
	StoryPoint            string   `json:"StoryPoint" xml:"StoryPoint"`
	Logo                  string   `json:"Logo" xml:"Logo"`
	Source                string   `json:"Source" xml:"Source"`
	AvatarUrl             string   `json:"AvatarUrl" xml:"AvatarUrl"`
	TaskflowstatusId      string   `json:"TaskflowstatusId" xml:"TaskflowstatusId"`
	Pos                   int      `json:"Pos" xml:"Pos"`
	RootCollectionId      string   `json:"RootCollectionId" xml:"RootCollectionId"`
	Content               string   `json:"Content" xml:"Content"`
	IsTemplate            bool     `json:"IsTemplate" xml:"IsTemplate"`
	ProjectId             string   `json:"ProjectId" xml:"ProjectId"`
	IsDone                bool     `json:"IsDone" xml:"IsDone"`
	Id                    string   `json:"Id" xml:"Id"`
	OrganizationId        string   `json:"OrganizationId" xml:"OrganizationId"`
	EndDate               string   `json:"EndDate" xml:"EndDate"`
	Accomplished          string   `json:"Accomplished" xml:"Accomplished"`
	Email                 string   `json:"Email" xml:"Email"`
	AliyunPk              string   `json:"AliyunPk" xml:"AliyunPk"`
	UniqueIdPrefix        string   `json:"UniqueIdPrefix" xml:"UniqueIdPrefix"`
	Note                  string   `json:"Note" xml:"Note"`
	StartDate             string   `json:"StartDate" xml:"StartDate"`
	Customfields          string   `json:"Customfields" xml:"Customfields"`
	Visible               string   `json:"Visible" xml:"Visible"`
	IsArchived            bool     `json:"IsArchived" xml:"IsArchived"`
	IsDeleted             bool     `json:"IsDeleted" xml:"IsDeleted"`
	IsSuspended           bool     `json:"IsSuspended" xml:"IsSuspended"`
	UniqueId              int      `json:"UniqueId" xml:"UniqueId"`
	Status                string   `json:"Status" xml:"Status"`
	TaskType              string   `json:"TaskType" xml:"TaskType"`
	Py                    string   `json:"Py" xml:"Py"`
	Pinyin                string   `json:"Pinyin" xml:"Pinyin"`
	ModifierId            string   `json:"ModifierId" xml:"ModifierId"`
	Visibility            string   `json:"Visibility" xml:"Visibility"`
	Name                  string   `json:"Name" xml:"Name"`
	Category              string   `json:"Category" xml:"Category"`
	Created               string   `json:"Created" xml:"Created"`
	ExecutorId            string   `json:"ExecutorId" xml:"ExecutorId"`
	DefaultCollectionId   string   `json:"DefaultCollectionId" xml:"DefaultCollectionId"`
	NormalType            string   `json:"NormalType" xml:"NormalType"`
	Executor              string   `json:"Executor" xml:"Executor"`
	SourceType            string   `json:"SourceType" xml:"SourceType"`
	SortMethod            string   `json:"SortMethod" xml:"SortMethod"`
	CreatorId             string   `json:"CreatorId" xml:"CreatorId"`
	DefaultRoleId         string   `json:"DefaultRoleId" xml:"DefaultRoleId"`
	Description           string   `json:"Description" xml:"Description"`
	Phone                 string   `json:"Phone" xml:"Phone"`
	SprintId              string   `json:"SprintId" xml:"SprintId"`
	Priority              int      `json:"Priority" xml:"Priority"`
	Updated               string   `json:"Updated" xml:"Updated"`
	TasklistId            string   `json:"TasklistId" xml:"TasklistId"`
	AncestorIds           string   `json:"AncestorIds" xml:"AncestorIds"`
	SourceId              string   `json:"SourceId" xml:"SourceId"`
	NextTaskUniqueId      int      `json:"NextTaskUniqueId" xml:"NextTaskUniqueId"`
	ScenarioFieldConfigId string   `json:"ScenarioFieldConfigId" xml:"ScenarioFieldConfigId"`
	DueDate               string   `json:"DueDate" xml:"DueDate"`
	Rating                int      `json:"Rating" xml:"Rating"`
	PlanToDo              PlanToDo `json:"PlanToDo" xml:"PlanToDo"`
}
