/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"testing"

	"github.com/nats-io/nats-server/v2/server"

	"github.com/numaproj/numaflow/pkg/shared/clients/nats"
)

// JetStreamClient is used to get a testing JetStream client instance
func JetStreamClient(t *testing.T, s *server.Server) nats.JetStreamClient {
	return nats.NewDefaultJetStreamClient(s.ClientURL())
}
