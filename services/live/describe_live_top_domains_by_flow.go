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

// DescribeLiveTopDomainsByFlow invokes the live.DescribeLiveTopDomainsByFlow API synchronously
// api document: https://help.aliyun.com/api/live/describelivetopdomainsbyflow.html
func (client *Client) DescribeLiveTopDomainsByFlow(request *DescribeLiveTopDomainsByFlowRequest) (response *DescribeLiveTopDomainsByFlowResponse, err error) {
	response = CreateDescribeLiveTopDomainsByFlowResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveTopDomainsByFlowWithChan invokes the live.DescribeLiveTopDomainsByFlow API asynchronously
// api document: https://help.aliyun.com/api/live/describelivetopdomainsbyflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveTopDomainsByFlowWithChan(request *DescribeLiveTopDomainsByFlowRequest) (<-chan *DescribeLiveTopDomainsByFlowResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveTopDomainsByFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveTopDomainsByFlow(request)
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

// DescribeLiveTopDomainsByFlowWithCallback invokes the live.DescribeLiveTopDomainsByFlow API asynchronously
// api document: https://help.aliyun.com/api/live/describelivetopdomainsbyflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveTopDomainsByFlowWithCallback(request *DescribeLiveTopDomainsByFlowRequest, callback func(response *DescribeLiveTopDomainsByFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveTopDomainsByFlowResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveTopDomainsByFlow(request)
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

// DescribeLiveTopDomainsByFlowRequest is the request struct for api DescribeLiveTopDomainsByFlow
type DescribeLiveTopDomainsByFlowRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Limit     requests.Integer `position:"Query" name:"Limit"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveTopDomainsByFlowResponse is the response struct for api DescribeLiveTopDomainsByFlow
type DescribeLiveTopDomainsByFlowResponse struct {
	*responses.BaseResponse
	RequestId         string     `json:"RequestId" xml:"RequestId"`
	StartTime         string     `json:"StartTime" xml:"StartTime"`
	EndTime           string     `json:"EndTime" xml:"EndTime"`
	DomainCount       int64      `json:"DomainCount" xml:"DomainCount"`
	DomainOnlineCount int64      `json:"DomainOnlineCount" xml:"DomainOnlineCount"`
	TopDomains        TopDomains `json:"TopDomains" xml:"TopDomains"`
}

// CreateDescribeLiveTopDomainsByFlowRequest creates a request to invoke DescribeLiveTopDomainsByFlow API
func CreateDescribeLiveTopDomainsByFlowRequest() (request *DescribeLiveTopDomainsByFlowRequest) {
	request = &DescribeLiveTopDomainsByFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveTopDomainsByFlow", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveTopDomainsByFlowResponse creates a response to parse from DescribeLiveTopDomainsByFlow response
func CreateDescribeLiveTopDomainsByFlowResponse() (response *DescribeLiveTopDomainsByFlowResponse) {
	response = &DescribeLiveTopDomainsByFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
