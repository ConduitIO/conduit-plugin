// Copyright © 2022 Meroxa, Inc.
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

package toproto

import (
	"github.com/conduitio/conduit-connector-protocol/cpluginv1"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

func DestinationConfigureResponse(_ cpluginv1.DestinationConfigureResponse) (*connectorv1.Destination_Configure_Response, error) {
	return &connectorv1.Destination_Configure_Response{}, nil
}

func DestinationStartResponse(_ cpluginv1.DestinationStartResponse) (*connectorv1.Destination_Start_Response, error) {
	return &connectorv1.Destination_Start_Response{}, nil
}

func DestinationRunResponse(in cpluginv1.DestinationRunResponse) (*connectorv1.Destination_Run_Response, error) {
	out := connectorv1.Destination_Run_Response{
		AckPosition: in.AckPosition,
		Error:       in.Error,
	}
	return &out, nil
}

func DestinationStopResponse(_ cpluginv1.DestinationStopResponse) (*connectorv1.Destination_Stop_Response, error) {
	return &connectorv1.Destination_Stop_Response{}, nil
}

func DestinationTeardownResponse(_ cpluginv1.DestinationTeardownResponse) (*connectorv1.Destination_Teardown_Response, error) {
	return &connectorv1.Destination_Teardown_Response{}, nil
}

func DestinationLifecycleOnCreatedResponse(_ cpluginv1.DestinationLifecycleOnCreatedResponse) (*connectorv1.Destination_Lifecycle_OnCreated_Response, error) {
	return &connectorv1.Destination_Lifecycle_OnCreated_Response{}, nil
}
func DestinationLifecycleOnUpdatedResponse(_ cpluginv1.DestinationLifecycleOnUpdatedResponse) (*connectorv1.Destination_Lifecycle_OnUpdated_Response, error) {
	return &connectorv1.Destination_Lifecycle_OnUpdated_Response{}, nil
}
func DestinationLifecycleOnDeletedResponse(_ cpluginv1.DestinationLifecycleOnDeletedResponse) (*connectorv1.Destination_Lifecycle_OnDeleted_Response, error) {
	return &connectorv1.Destination_Lifecycle_OnDeleted_Response{}, nil
}
