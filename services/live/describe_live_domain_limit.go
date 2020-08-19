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

// DescribeLiveDomainLimit invokes the live.DescribeLiveDomainLimit API synchronously
// api document: https://help.aliyun.com/api/live/describelivedomainlimit.html
func (client *Client) DescribeLiveDomainLimit(request *DescribeLiveDomainLimitRequest) (response *DescribeLiveDomainLimitResponse, err error) {
	response = CreateDescribeLiveDomainLimitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainLimitWithChan invokes the live.DescribeLiveDomainLimit API asynchronously
// api document: https://help.aliyun.com/api/live/describelivedomainlimit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveDomainLimitWithChan(request *DescribeLiveDomainLimitRequest) (<-chan *DescribeLiveDomainLimitResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainLimitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainLimit(request)
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

// DescribeLiveDomainLimitWithCallback invokes the live.DescribeLiveDomainLimit API asynchronously
// api document: https://help.aliyun.com/api/live/describelivedomainlimit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveDomainLimitWithCallback(request *DescribeLiveDomainLimitRequest, callback func(response *DescribeLiveDomainLimitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainLimitResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainLimit(request)
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

// DescribeLiveDomainLimitRequest is the request struct for api DescribeLiveDomainLimit
type DescribeLiveDomainLimitRequest struct {
	*requests.RpcRequest
	LiveapiRequestFrom string           `position:"Query" name:"LiveapiRequestFrom"`
	DomainName         string           `position:"Query" name:"DomainName"`
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainLimitResponse is the response struct for api DescribeLiveDomainLimit
type DescribeLiveDomainLimitResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	LiveDomainLimitList LiveDomainLimitList `json:"LiveDomainLimitList" xml:"LiveDomainLimitList"`
}

// CreateDescribeLiveDomainLimitRequest creates a request to invoke DescribeLiveDomainLimit API
func CreateDescribeLiveDomainLimitRequest() (request *DescribeLiveDomainLimitRequest) {
	request = &DescribeLiveDomainLimitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainLimit", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainLimitResponse creates a response to parse from DescribeLiveDomainLimit response
func CreateDescribeLiveDomainLimitResponse() (response *DescribeLiveDomainLimitResponse) {
	response = &DescribeLiveDomainLimitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
