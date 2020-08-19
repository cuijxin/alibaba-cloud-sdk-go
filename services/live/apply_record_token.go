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

// ApplyRecordToken invokes the live.ApplyRecordToken API synchronously
// api document: https://help.aliyun.com/api/live/applyrecordtoken.html
func (client *Client) ApplyRecordToken(request *ApplyRecordTokenRequest) (response *ApplyRecordTokenResponse, err error) {
	response = CreateApplyRecordTokenResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyRecordTokenWithChan invokes the live.ApplyRecordToken API asynchronously
// api document: https://help.aliyun.com/api/live/applyrecordtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyRecordTokenWithChan(request *ApplyRecordTokenRequest) (<-chan *ApplyRecordTokenResponse, <-chan error) {
	responseChan := make(chan *ApplyRecordTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyRecordToken(request)
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

// ApplyRecordTokenWithCallback invokes the live.ApplyRecordToken API asynchronously
// api document: https://help.aliyun.com/api/live/applyrecordtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyRecordTokenWithCallback(request *ApplyRecordTokenRequest, callback func(response *ApplyRecordTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyRecordTokenResponse
		var err error
		defer close(result)
		response, err = client.ApplyRecordToken(request)
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

// ApplyRecordTokenRequest is the request struct for api ApplyRecordToken
type ApplyRecordTokenRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// ApplyRecordTokenResponse is the response struct for api ApplyRecordToken
type ApplyRecordTokenResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	SecurityToken   string `json:"SecurityToken" xml:"SecurityToken"`
	AccessKeySecret string `json:"AccessKeySecret" xml:"AccessKeySecret"`
	AccessKeyId     string `json:"AccessKeyId" xml:"AccessKeyId"`
	Expiration      string `json:"Expiration" xml:"Expiration"`
}

// CreateApplyRecordTokenRequest creates a request to invoke ApplyRecordToken API
func CreateApplyRecordTokenRequest() (request *ApplyRecordTokenRequest) {
	request = &ApplyRecordTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ApplyRecordToken", "", "")
	request.Method = requests.POST
	return
}

// CreateApplyRecordTokenResponse creates a response to parse from ApplyRecordToken response
func CreateApplyRecordTokenResponse() (response *ApplyRecordTokenResponse) {
	response = &ApplyRecordTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
