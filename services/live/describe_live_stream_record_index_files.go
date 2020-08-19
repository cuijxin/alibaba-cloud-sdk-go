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

// DescribeLiveStreamRecordIndexFiles invokes the live.DescribeLiveStreamRecordIndexFiles API synchronously
// api document: https://help.aliyun.com/api/live/describelivestreamrecordindexfiles.html
func (client *Client) DescribeLiveStreamRecordIndexFiles(request *DescribeLiveStreamRecordIndexFilesRequest) (response *DescribeLiveStreamRecordIndexFilesResponse, err error) {
	response = CreateDescribeLiveStreamRecordIndexFilesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamRecordIndexFilesWithChan invokes the live.DescribeLiveStreamRecordIndexFiles API asynchronously
// api document: https://help.aliyun.com/api/live/describelivestreamrecordindexfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamRecordIndexFilesWithChan(request *DescribeLiveStreamRecordIndexFilesRequest) (<-chan *DescribeLiveStreamRecordIndexFilesResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRecordIndexFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRecordIndexFiles(request)
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

// DescribeLiveStreamRecordIndexFilesWithCallback invokes the live.DescribeLiveStreamRecordIndexFiles API asynchronously
// api document: https://help.aliyun.com/api/live/describelivestreamrecordindexfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamRecordIndexFilesWithCallback(request *DescribeLiveStreamRecordIndexFilesRequest, callback func(response *DescribeLiveStreamRecordIndexFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRecordIndexFilesResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRecordIndexFiles(request)
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

// DescribeLiveStreamRecordIndexFilesRequest is the request struct for api DescribeLiveStreamRecordIndexFiles
type DescribeLiveStreamRecordIndexFilesRequest struct {
	*requests.RpcRequest
	StartTime     string           `position:"Query" name:"StartTime"`
	PageNum       requests.Integer `position:"Query" name:"PageNum"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	StreamName    string           `position:"Query" name:"StreamName"`
	Order         string           `position:"Query" name:"Order"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamRecordIndexFilesResponse is the response struct for api DescribeLiveStreamRecordIndexFiles
type DescribeLiveStreamRecordIndexFilesResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	PageNum             int                 `json:"PageNum" xml:"PageNum"`
	PageSize            int                 `json:"PageSize" xml:"PageSize"`
	Order               string              `json:"Order" xml:"Order"`
	TotalNum            int                 `json:"TotalNum" xml:"TotalNum"`
	TotalPage           int                 `json:"TotalPage" xml:"TotalPage"`
	RecordIndexInfoList RecordIndexInfoList `json:"RecordIndexInfoList" xml:"RecordIndexInfoList"`
}

// CreateDescribeLiveStreamRecordIndexFilesRequest creates a request to invoke DescribeLiveStreamRecordIndexFiles API
func CreateDescribeLiveStreamRecordIndexFilesRequest() (request *DescribeLiveStreamRecordIndexFilesRequest) {
	request = &DescribeLiveStreamRecordIndexFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamRecordIndexFiles", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveStreamRecordIndexFilesResponse creates a response to parse from DescribeLiveStreamRecordIndexFiles response
func CreateDescribeLiveStreamRecordIndexFilesResponse() (response *DescribeLiveStreamRecordIndexFilesResponse) {
	response = &DescribeLiveStreamRecordIndexFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
