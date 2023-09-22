/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package raftapi

import (
	"github.com/codeallergy/glue"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/serf/cmd/serf/command/agent"
	"github.com/hashicorp/serf/serf"
	"github.com/sprintframework/raftpb"
	"github.com/sprintframework/sprint"
	"google.golang.org/grpc"
	"reflect"
)

var RaftGrpcServerClass = reflect.TypeOf((*RaftGrpcServer)(nil)).Elem()

type RaftGrpcServer interface {
	glue.InitializingBean
	sprint.Component
}

var RaftClientPoolClass = reflect.TypeOf((*RaftClientPool)(nil)).Elem()

type RaftClientPool interface {
	glue.InitializingBean
	glue.DisposableBean

	GetAPIEndpoint(raftAddress string) (string, error)

	GetAPIConn(raftAddress raft.ServerAddress) (*grpc.ClientConn, error)

	Close() error

}

/**
Finite State Machine Response
 */
type FSMResponse struct {
	Status   *raftpb.Status
	Err      error
}

var RaftServiceClass = reflect.TypeOf((*RaftService)(nil)).Elem()

type RaftService interface {
	glue.InitializingBean
	raft.FSM

}

var RaftServerClass = reflect.TypeOf((*RaftServer)(nil)).Elem()

type RaftServer interface {
	sprint.Server
	sprint.Component

	Transport() (raft.Transport, bool)

	Raft() (*raft.Raft, bool)

	IsLeader() bool

}

var SerfServerClass = reflect.TypeOf((*SerfServer)(nil)).Elem()

type SerfServer interface {
	sprint.Server
	sprint.Component

	Config() (*serf.Config, bool)

	Serf() (*serf.Serf, bool)

	Agent() (*agent.Agent, bool)

}
