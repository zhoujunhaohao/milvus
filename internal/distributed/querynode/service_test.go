// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcquerynode

import (
	"context"
	"errors"
	"testing"

	"github.com/milvus-io/milvus/internal/types"

	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/proto/milvuspb"
	"github.com/milvus-io/milvus/internal/proto/querypb"
	"github.com/stretchr/testify/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type MockQueryNode struct {
	states     *internalpb.ComponentStates
	status     *commonpb.Status
	err        error
	initErr    error
	startErr   error
	regErr     error
	stopErr    error
	strResp    *milvuspb.StringResponse
	infoResp   *querypb.GetSegmentInfoResponse
	metricResp *milvuspb.GetMetricsResponse
}

func (m *MockQueryNode) Init() error {
	return m.initErr
}

func (m *MockQueryNode) Start() error {
	return m.startErr
}

func (m *MockQueryNode) Stop() error {
	return m.stopErr
}

func (m *MockQueryNode) Register() error {
	return m.regErr
}

func (m *MockQueryNode) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	return m.states, m.err
}

func (m *MockQueryNode) GetStatisticsChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return m.strResp, m.err
}

func (m *MockQueryNode) GetTimeTickChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return m.strResp, m.err
}

func (m *MockQueryNode) AddQueryChannel(ctx context.Context, req *querypb.AddQueryChannelRequest) (*commonpb.Status, error) {
	return m.status, m.err
}

func (m *MockQueryNode) RemoveQueryChannel(ctx context.Context, req *querypb.RemoveQueryChannelRequest) (*commonpb.Status, error) {
	return m.status, m.err
}

func (m *MockQueryNode) WatchDmChannels(ctx context.Context, req *querypb.WatchDmChannelsRequest) (*commonpb.Status, error) {
	return m.status, m.err
}

func (m *MockQueryNode) WatchDeltaChannels(ctx context.Context, req *querypb.WatchDeltaChannelsRequest) (*commonpb.Status, error) {
	return m.status, m.err
}

func (m *MockQueryNode) LoadSegments(ctx context.Context, req *querypb.LoadSegmentsRequest) (*commonpb.Status, error) {
	return m.status, m.err
}

func (m *MockQueryNode) ReleaseCollection(ctx context.Context, req *querypb.ReleaseCollectionRequest) (*commonpb.Status, error) {
	return m.status, m.err
}

func (m *MockQueryNode) ReleasePartitions(ctx context.Context, req *querypb.ReleasePartitionsRequest) (*commonpb.Status, error) {
	return m.status, m.err
}

func (m *MockQueryNode) ReleaseSegments(ctx context.Context, req *querypb.ReleaseSegmentsRequest) (*commonpb.Status, error) {
	return m.status, m.err
}

func (m *MockQueryNode) GetSegmentInfo(ctx context.Context, req *querypb.GetSegmentInfoRequest) (*querypb.GetSegmentInfoResponse, error) {
	return m.infoResp, m.err
}

func (m *MockQueryNode) GetMetrics(ctx context.Context, req *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error) {
	return m.metricResp, m.err
}

func (m *MockQueryNode) SetEtcdClient(client *clientv3.Client) {
}

func (m *MockQueryNode) UpdateStateCode(code internalpb.StateCode) {
}

func (m *MockQueryNode) SetRootCoord(rc types.RootCoord) error {
	return m.err
}

func (m *MockQueryNode) SetIndexCoord(index types.IndexCoord) error {
	return m.err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type MockRootCoord struct {
	types.RootCoord
	initErr  error
	startErr error
	regErr   error
	stopErr  error
	stateErr commonpb.ErrorCode
}

func (m *MockRootCoord) Init() error {
	return m.initErr
}

func (m *MockRootCoord) Start() error {
	return m.startErr
}

func (m *MockRootCoord) Stop() error {
	return m.stopErr
}

func (m *MockRootCoord) Register() error {
	return m.regErr
}

func (m *MockRootCoord) SetEtcdClient(client *clientv3.Client) {
}

func (m *MockRootCoord) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	return &internalpb.ComponentStates{
		State:  &internalpb.ComponentInfo{StateCode: internalpb.StateCode_Healthy},
		Status: &commonpb.Status{ErrorCode: m.stateErr},
	}, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type MockIndexCoord struct {
	types.IndexCoord
	initErr  error
	startErr error
	regErr   error
	stopErr  error
	stateErr commonpb.ErrorCode
}

func (m *MockIndexCoord) Init() error {
	return m.initErr
}

func (m *MockIndexCoord) Start() error {
	return m.startErr
}

func (m *MockIndexCoord) Stop() error {
	return m.stopErr
}

func (m *MockIndexCoord) Register() error {
	return m.regErr
}

func (m *MockIndexCoord) SetEtcdClient(client *clientv3.Client) {
}

func (m *MockIndexCoord) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	return &internalpb.ComponentStates{
		State:  &internalpb.ComponentInfo{StateCode: internalpb.StateCode_Healthy},
		Status: &commonpb.Status{ErrorCode: m.stateErr},
	}, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func Test_NewServer(t *testing.T) {
	ctx := context.Background()
	server, err := NewServer(ctx, nil)
	assert.Nil(t, err)
	assert.NotNil(t, server)

	mqn := &MockQueryNode{
		states:     &internalpb.ComponentStates{State: &internalpb.ComponentInfo{StateCode: internalpb.StateCode_Healthy}},
		status:     &commonpb.Status{ErrorCode: commonpb.ErrorCode_Success},
		err:        nil,
		strResp:    &milvuspb.StringResponse{Status: &commonpb.Status{ErrorCode: commonpb.ErrorCode_Success}},
		infoResp:   &querypb.GetSegmentInfoResponse{Status: &commonpb.Status{ErrorCode: commonpb.ErrorCode_Success}},
		metricResp: &milvuspb.GetMetricsResponse{Status: &commonpb.Status{ErrorCode: commonpb.ErrorCode_Success}},
	}
	server.querynode = mqn

	t.Run("Run", func(t *testing.T) {
		server.rootCoord = &MockRootCoord{}
		server.indexCoord = &MockIndexCoord{}

		err = server.Run()
		assert.Nil(t, err)
	})

	t.Run("GetComponentStates", func(t *testing.T) {
		req := &internalpb.GetComponentStatesRequest{}
		states, err := server.GetComponentStates(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, internalpb.StateCode_Healthy, states.State.StateCode)
	})

	t.Run("GetStatisticsChannel", func(t *testing.T) {
		req := &internalpb.GetStatisticsChannelRequest{}
		resp, err := server.GetStatisticsChannel(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.Status.ErrorCode)
	})

	t.Run("GetTimeTickChannel", func(t *testing.T) {
		req := &internalpb.GetTimeTickChannelRequest{}
		resp, err := server.GetTimeTickChannel(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.Status.ErrorCode)
	})

	t.Run("AddQueryChannel", func(t *testing.T) {
		req := &querypb.AddQueryChannelRequest{}
		resp, err := server.AddQueryChannel(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.ErrorCode)
	})

	t.Run("RemoveQueryChannel", func(t *testing.T) {
		req := &querypb.RemoveQueryChannelRequest{}
		resp, err := server.RemoveQueryChannel(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.ErrorCode)
	})

	t.Run("WatchDmChannels", func(t *testing.T) {
		req := &querypb.WatchDmChannelsRequest{}
		resp, err := server.WatchDmChannels(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.ErrorCode)
	})

	t.Run("LoadSegments", func(t *testing.T) {
		req := &querypb.LoadSegmentsRequest{}
		resp, err := server.LoadSegments(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.ErrorCode)
	})

	t.Run("ReleaseCollection", func(t *testing.T) {
		req := &querypb.ReleaseCollectionRequest{}
		resp, err := server.ReleaseCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.ErrorCode)
	})

	t.Run("ReleasePartitions", func(t *testing.T) {
		req := &querypb.ReleasePartitionsRequest{}
		resp, err := server.ReleasePartitions(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.ErrorCode)
	})

	t.Run("ReleaseSegments", func(t *testing.T) {
		req := &querypb.ReleaseSegmentsRequest{}
		resp, err := server.ReleaseSegments(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.ErrorCode)
	})

	t.Run("GetSegmentInfo", func(t *testing.T) {
		req := &querypb.GetSegmentInfoRequest{}
		resp, err := server.GetSegmentInfo(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.Status.ErrorCode)
	})

	t.Run("GetMetrics", func(t *testing.T) {
		req := &milvuspb.GetMetricsRequest{
			Request: "",
		}
		resp, err := server.GetMetrics(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.Status.ErrorCode)
	})

	err = server.Stop()
	assert.Nil(t, err)
}

func Test_Run(t *testing.T) {
	ctx := context.Background()
	server, err := NewServer(ctx, nil)
	assert.Nil(t, err)
	assert.NotNil(t, server)

	server.querynode = &MockQueryNode{}
	server.indexCoord = &MockIndexCoord{}
	server.rootCoord = &MockRootCoord{initErr: errors.New("failed")}
	assert.Panics(t, func() { err = server.Run() })

	server.rootCoord = &MockRootCoord{startErr: errors.New("Failed")}
	assert.Panics(t, func() { err = server.Run() })

	server.querynode = &MockQueryNode{}
	server.rootCoord = &MockRootCoord{}
	server.indexCoord = &MockIndexCoord{initErr: errors.New("Failed")}
	assert.Panics(t, func() { err = server.Run() })

	server.indexCoord = &MockIndexCoord{startErr: errors.New("Failed")}
	assert.Panics(t, func() { err = server.Run() })

	server.indexCoord = &MockIndexCoord{}
	server.rootCoord = &MockRootCoord{}
	server.querynode = &MockQueryNode{initErr: errors.New("Failed")}
	err = server.Run()
	assert.Error(t, err)

	server.querynode = &MockQueryNode{startErr: errors.New("Failed")}
	err = server.Run()
	assert.Error(t, err)

	server.querynode = &MockQueryNode{regErr: errors.New("Failed")}
	err = server.Run()
	assert.Error(t, err)

	server.querynode = &MockQueryNode{stopErr: errors.New("Failed")}
	err = server.Stop()
	assert.Error(t, err)
}
