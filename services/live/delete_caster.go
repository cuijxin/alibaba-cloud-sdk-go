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

// DeleteCaster invokes the live.DeleteCaster API synchronously
// api document: https://help.aliyun.com/api/live/deletecaster.html
func (client *Client) DeleteCaster(request *DeleteCasterRequest) (response *DeleteCasterResponse, err error) {
	response = CreateDeleteCasterResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCasterWithChan invokes the live.DeleteCaster API asynchronously
// api document: https://help.aliyun.com/api/live/deletecaster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCasterWithChan(request *DeleteCasterRequest) (<-chan *DeleteCasterResponse, <-chan error) {
	responseChan := make(chan *DeleteCasterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCaster(request)
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

// DeleteCasterWithCallback invokes the live.DeleteCaster API asynchronously
// api document: https://help.aliyun.com/api/live/deletecaster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCasterWithCallback(request *DeleteCasterRequest, callback func(response *DeleteCasterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCasterResponse
		var err error
		defer close(result)
		response, err = client.DeleteCaster(request)
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

// DeleteCasterRequest is the request struct for api DeleteCaster
type DeleteCasterRequest struct {
	*requests.RpcRequest
	CasterId      string           `position:"Query" name:"CasterId"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DeleteCasterResponse is the response struct for api DeleteCaster
type DeleteCasterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CasterId  string `json:"CasterId" xml:"CasterId"`
}

// CreateDeleteCasterRequest creates a request to invoke DeleteCaster API
func CreateDeleteCasterRequest() (request *DeleteCasterRequest) {
	request = &DeleteCasterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteCaster", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCasterResponse creates a response to parse from DeleteCaster response
func CreateDeleteCasterResponse() (response *DeleteCasterResponse) {
	response = &DeleteCasterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
