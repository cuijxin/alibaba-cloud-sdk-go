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

// EffectCasterUrgent invokes the live.EffectCasterUrgent API synchronously
// api document: https://help.aliyun.com/api/live/effectcasterurgent.html
func (client *Client) EffectCasterUrgent(request *EffectCasterUrgentRequest) (response *EffectCasterUrgentResponse, err error) {
	response = CreateEffectCasterUrgentResponse()
	err = client.DoAction(request, response)
	return
}

// EffectCasterUrgentWithChan invokes the live.EffectCasterUrgent API asynchronously
// api document: https://help.aliyun.com/api/live/effectcasterurgent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EffectCasterUrgentWithChan(request *EffectCasterUrgentRequest) (<-chan *EffectCasterUrgentResponse, <-chan error) {
	responseChan := make(chan *EffectCasterUrgentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EffectCasterUrgent(request)
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

// EffectCasterUrgentWithCallback invokes the live.EffectCasterUrgent API asynchronously
// api document: https://help.aliyun.com/api/live/effectcasterurgent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EffectCasterUrgentWithCallback(request *EffectCasterUrgentRequest, callback func(response *EffectCasterUrgentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EffectCasterUrgentResponse
		var err error
		defer close(result)
		response, err = client.EffectCasterUrgent(request)
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

// EffectCasterUrgentRequest is the request struct for api EffectCasterUrgent
type EffectCasterUrgentRequest struct {
	*requests.RpcRequest
	CasterId string           `position:"Query" name:"CasterId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	SceneId  string           `position:"Query" name:"SceneId"`
}

// EffectCasterUrgentResponse is the response struct for api EffectCasterUrgent
type EffectCasterUrgentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEffectCasterUrgentRequest creates a request to invoke EffectCasterUrgent API
func CreateEffectCasterUrgentRequest() (request *EffectCasterUrgentRequest) {
	request = &EffectCasterUrgentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "EffectCasterUrgent", "", "")
	request.Method = requests.POST
	return
}

// CreateEffectCasterUrgentResponse creates a response to parse from EffectCasterUrgent response
func CreateEffectCasterUrgentResponse() (response *EffectCasterUrgentResponse) {
	response = &EffectCasterUrgentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
