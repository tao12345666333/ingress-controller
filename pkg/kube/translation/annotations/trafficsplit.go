// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package annotations

//	apisixv1 "github.com/apache/apisix-ingress-controller/pkg/types/apisix/v1"

const (
	_enableTrafficPlugin = "k8s.apisix.apache.org/canary"
	_trafficPluginWeight = "k8s.apisix.apache.org/canary-weight"
)

type trafficSplit struct{}

// NewCorsHandler creates a handler to convert annotations about
// cors to APISIX cors plugin.
func NewTrafficSplitHandler() Handler {
	return &trafficSplit{}
}

func (ts *trafficSplit) PluginName() string {
	return "traffic-split"
}

func (ts *trafficSplit) Handle(e Extractor) (interface{}, error) {
	if !e.GetBoolAnnotation(_enableTrafficPlugin) {
		return nil, nil
	}

	return nil, nil

}
