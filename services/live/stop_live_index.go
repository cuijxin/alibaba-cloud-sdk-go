package live

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

// StopLiveIndex invokes the live.StopLiveIndex API synchronously
// api document: https://help.aliyun.com/api/live/stopliveindex.html
func (client *Client) StopLiveIndex(request *StopLiveIndexRequest) (response *StopLiveIndexResponse, err error) {
	response = CreateStopLiveIndexResponse()
	err = client.DoAction(request, response)
	return
}

// StopLiveIndexWithChan invokes the live.StopLiveIndex API asynchronously
// api document: https://help.aliyun.com/api/live/stopliveindex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopLiveIndexWithChan(request *StopLiveIndexRequest) (<-chan *StopLiveIndexResponse, <-chan error) {
	responseChan := make(chan *StopLiveIndexResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopLiveIndex(request)
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

// StopLiveIndexWithCallback invokes the live.StopLiveIndex API asynchronously
// api document: https://help.aliyun.com/api/live/stopliveindex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopLiveIndexWithCallback(request *StopLiveIndexRequest, callback func(response *StopLiveIndexResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopLiveIndexResponse
		var err error
		defer close(result)
		response, err = client.StopLiveIndex(request)
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

// StopLiveIndexRequest is the request struct for api StopLiveIndex
type StopLiveIndexRequest struct {
	*requests.RpcRequest
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	TaskId     string           `position:"Query" name:"TaskId"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// StopLiveIndexResponse is the response struct for api StopLiveIndex
type StopLiveIndexResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopLiveIndexRequest creates a request to invoke StopLiveIndex API
func CreateStopLiveIndexRequest() (request *StopLiveIndexRequest) {
	request = &StopLiveIndexRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "StopLiveIndex", "", "")
	request.Method = requests.POST
	return
}

// CreateStopLiveIndexResponse creates a response to parse from StopLiveIndex response
func CreateStopLiveIndexResponse() (response *StopLiveIndexResponse) {
	response = &StopLiveIndexResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
