/*
   LabAssistant
   Copyright (C) 2020  aarontillekeratne

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
  Author: aarontillekeratne
  Contact: github.com/codeBehindMe
*/

package tracker_server

import (
	"com.github.com/codeBehindMe/LabAssistant/utils"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type TrackerServer struct {
	addr string
}

func (t *TrackerServer) GetRegisteredLabs(empty *Empty, server TrackerService_GetRegisteredLabsServer) error {
	err := server.Send(&Lab{})
	if err != nil {
		return err
	}
	return nil
}

func New(addr string) *TrackerServer {
	return &TrackerServer{addr: addr}
}

func (t *TrackerServer) GetServiceVersion(ctx context.Context, e *Empty) (*ServiceVersionInfo, error) {
	log.Printf("Received request for version info from client")
	version, err := utils.GetDistributionVersion()
	if err != nil {
		log.Fatalf("Error while getting version info: %v", err)
	}
	return &ServiceVersionInfo{
		Version: version,
	}, nil
}

func (t *TrackerServer) RegisterLabCreation(ctx context.Context, l *Lab) (*SuccessfulOperation, error) {
	log.Printf("Received request for registration of lab creation from client")
	return &SuccessfulOperation{
		Code: 0,
	}, nil
}

func (t TrackerServer) Serve() {
	lis, err := net.Listen("tcp", t.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterTrackerServiceServer(grpcServer, &t)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server for TrackerServer: %v", err)
	}
}
