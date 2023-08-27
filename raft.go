/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package raftapi

import (
	"github.com/codeallergy/glue"
	"github.com/sprintframework/sprint"
	"github.com/openraft/raftpb"
	"github.com/hashicorp/raft"
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

}
