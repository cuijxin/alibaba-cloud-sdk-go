package cms

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateTask invokes the cms.CreateTask API synchronously
// api document: https://help.aliyun.com/api/cms/createtask.html
func (client *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
	response = CreateCreateTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTaskWithChan invokes the cms.CreateTask API asynchronously
// api document: https://help.aliyun.com/api/cms/createtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTaskWithChan(request *CreateTaskRequest) (<-chan *CreateTaskResponse, <-chan error) {
	responseChan := make(chan *CreateTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTask(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateTaskWithCallback invokes the cms.CreateTask API asynchronously
// api document: https://help.aliyun.com/api/cms/createtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTaskWithCallback(request *CreateTaskRequest, callback func(response *CreateTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateTask(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateTaskRequest is the request struct for api CreateTask
type CreateTaskRequest struct {
	*requests.RpcRequest
	AgentGroup    string `position:"Query" name:"AgentGroup"`
	AlertName     string `position:"Query" name:"AlertName"`
	TaskId        string `position:"Query" name:"TaskId"`
	ClientIds     string `position:"Query" name:"ClientIds"`
	TaskType      string `position:"Query" name:"TaskType"`
	TaskName      string `position:"Query" name:"TaskName"`
	TaskState     string `position:"Query" name:"TaskState"`
	AgentType     string `position:"Query" name:"AgentType"`
	ReportProject string `position:"Query" name:"ReportProject"`
	Address       string `position:"Query" name:"Address"`
	Interval      string `position:"Query" name:"Interval"`
	EndTime       string `position:"Query" name:"EndTime"`
	IspCity       string `position:"Query" name:"IspCity"`
	Options       string `position:"Query" name:"Options"`
	AlertInfo     string `position:"Query" name:"AlertInfo"`
	AlertRule     string `position:"Query" name:"AlertRule"`
	Ip            string `position:"Query" name:"Ip"`
}

// CreateTaskResponse is the response struct for api CreateTask
type CreateTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTaskRequest creates a request to invoke CreateTask API
func CreateCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "CreateTask", "cms", "openAPI")
	return
}

// CreateCreateTaskResponse creates a response to parse from CreateTask response
func CreateCreateTaskResponse() (response *CreateTaskResponse) {
	response = &CreateTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}