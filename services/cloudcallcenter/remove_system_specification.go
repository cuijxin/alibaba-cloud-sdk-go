package cloudcallcenter

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

// RemoveSystemSpecification invokes the cloudcallcenter.RemoveSystemSpecification API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/removesystemspecification.html
func (client *Client) RemoveSystemSpecification(request *RemoveSystemSpecificationRequest) (response *RemoveSystemSpecificationResponse, err error) {
	response = CreateRemoveSystemSpecificationResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveSystemSpecificationWithChan invokes the cloudcallcenter.RemoveSystemSpecification API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/removesystemspecification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveSystemSpecificationWithChan(request *RemoveSystemSpecificationRequest) (<-chan *RemoveSystemSpecificationResponse, <-chan error) {
	responseChan := make(chan *RemoveSystemSpecificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveSystemSpecification(request)
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

// RemoveSystemSpecificationWithCallback invokes the cloudcallcenter.RemoveSystemSpecification API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/removesystemspecification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveSystemSpecificationWithCallback(request *RemoveSystemSpecificationRequest, callback func(response *RemoveSystemSpecificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveSystemSpecificationResponse
		var err error
		defer close(result)
		response, err = client.RemoveSystemSpecification(request)
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

// RemoveSystemSpecificationRequest is the request struct for api RemoveSystemSpecification
type RemoveSystemSpecificationRequest struct {
	*requests.RpcRequest
}

// RemoveSystemSpecificationResponse is the response struct for api RemoveSystemSpecification
type RemoveSystemSpecificationResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateRemoveSystemSpecificationRequest creates a request to invoke RemoveSystemSpecification API
func CreateRemoveSystemSpecificationRequest() (request *RemoveSystemSpecificationRequest) {
	request = &RemoveSystemSpecificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "RemoveSystemSpecification", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveSystemSpecificationResponse creates a response to parse from RemoveSystemSpecification response
func CreateRemoveSystemSpecificationResponse() (response *RemoveSystemSpecificationResponse) {
	response = &RemoveSystemSpecificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}