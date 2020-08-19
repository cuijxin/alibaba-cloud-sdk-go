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

// DescribeUpBpsPeakData invokes the live.DescribeUpBpsPeakData API synchronously
// api document: https://help.aliyun.com/api/live/describeupbpspeakdata.html
func (client *Client) DescribeUpBpsPeakData(request *DescribeUpBpsPeakDataRequest) (response *DescribeUpBpsPeakDataResponse, err error) {
	response = CreateDescribeUpBpsPeakDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUpBpsPeakDataWithChan invokes the live.DescribeUpBpsPeakData API asynchronously
// api document: https://help.aliyun.com/api/live/describeupbpspeakdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUpBpsPeakDataWithChan(request *DescribeUpBpsPeakDataRequest) (<-chan *DescribeUpBpsPeakDataResponse, <-chan error) {
	responseChan := make(chan *DescribeUpBpsPeakDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUpBpsPeakData(request)
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

// DescribeUpBpsPeakDataWithCallback invokes the live.DescribeUpBpsPeakData API asynchronously
// api document: https://help.aliyun.com/api/live/describeupbpspeakdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUpBpsPeakDataWithCallback(request *DescribeUpBpsPeakDataRequest, callback func(response *DescribeUpBpsPeakDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUpBpsPeakDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeUpBpsPeakData(request)
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

// DescribeUpBpsPeakDataRequest is the request struct for api DescribeUpBpsPeakData
type DescribeUpBpsPeakDataRequest struct {
	*requests.RpcRequest
	StartTime    string           `position:"Query" name:"StartTime"`
	DomainName   string           `position:"Query" name:"DomainName"`
	EndTime      string           `position:"Query" name:"EndTime"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	DomainSwitch string           `position:"Query" name:"DomainSwitch"`
}

// DescribeUpBpsPeakDataResponse is the response struct for api DescribeUpBpsPeakData
type DescribeUpBpsPeakDataResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	DescribeUpPeakTraffics DescribeUpPeakTraffics `json:"DescribeUpPeakTraffics" xml:"DescribeUpPeakTraffics"`
}

// CreateDescribeUpBpsPeakDataRequest creates a request to invoke DescribeUpBpsPeakData API
func CreateDescribeUpBpsPeakDataRequest() (request *DescribeUpBpsPeakDataRequest) {
	request = &DescribeUpBpsPeakDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeUpBpsPeakData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUpBpsPeakDataResponse creates a response to parse from DescribeUpBpsPeakData response
func CreateDescribeUpBpsPeakDataResponse() (response *DescribeUpBpsPeakDataResponse) {
	response = &DescribeUpBpsPeakDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
