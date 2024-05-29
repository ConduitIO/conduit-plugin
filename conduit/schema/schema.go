// Copyright © 2024 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"context"
	"fmt"

	"github.com/conduitio/conduit-commons/schema"
	v1 "github.com/conduitio/conduit-connector-protocol/conduit/schema/v1"
	"google.golang.org/grpc"
)

type schemaServiceKey struct{}

type Service interface {
	Create(context.Context, CreateRequest) (schema.Instance, error)
	Get(context.Context, GetRequest) (schema.Instance, error)
}

func WithSchemaService(ctx context.Context, s Service) context.Context {
	return context.WithValue(ctx, schemaServiceKey{}, s)
}

func New(ctx context.Context) (Service, error) {
	service := ctx.Value(schemaServiceKey{})
	if service != nil {
		return service.(Service), nil
	}

	conn, err := grpc.NewClient("localhost:8085")
	if err != nil {
		return nil, fmt.Errorf("failed creating gRPC client: %w", err)
	}

	return v1.NewClient(conn)
}

type CreateRequest struct {
	Name  string
	Bytes []byte
}

type GetRequest struct {
	ID string
}
