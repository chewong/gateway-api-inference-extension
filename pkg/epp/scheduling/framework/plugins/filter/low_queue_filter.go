/*
Copyright 2025 The Kubernetes Authors.

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

package filter

import (
	"context"
	"encoding/json"
	"fmt"

	"sigs.k8s.io/gateway-api-inference-extension/pkg/epp/plugins"
	"sigs.k8s.io/gateway-api-inference-extension/pkg/epp/scheduling/framework"
	"sigs.k8s.io/gateway-api-inference-extension/pkg/epp/scheduling/types"
)

const (
	LowQueueFilterType           = "low-queue-filter"
	DefaultQueueingThresholdLoRA = 128
)

type lowQueueFilterParameters struct {
	Threshold int `json:"threshold"`
}

// compile-time type validation
var _ framework.Filter = &LowQueueFilter{}

// LowQueueFilterFactory defines the factory function for LowQueueFilter.
func LowQueueFilterFactory(name string, rawParameters json.RawMessage, _ plugins.Handle) (plugins.Plugin, error) {
	parameters := lowQueueFilterParameters{Threshold: DefaultQueueingThresholdLoRA}
	if rawParameters != nil {
		if err := json.Unmarshal(rawParameters, &parameters); err != nil {
			return nil, fmt.Errorf("failed to parse the parameters of the '%s' filter - %w", LowQueueFilterType, err)
		}
	}

	return NewLowQueueFilter(parameters.Threshold).WithName(name), nil
}

// NewLowQueueFilter initializes a new LowQueueFilter and returns its pointer.
func NewLowQueueFilter(threshold int) *LowQueueFilter {
	return &LowQueueFilter{
		typedName:             plugins.TypedName{Type: LowQueueFilterType, Name: LowQueueFilterType},
		queueingThresholdLoRA: threshold,
	}
}

// LowQueueFilter returns pods that their waiting queue size is less than a configured threshold
type LowQueueFilter struct {
	typedName             plugins.TypedName
	queueingThresholdLoRA int
}

// TypedName returns the type and name tuple of this plugin instance.
func (f *LowQueueFilter) TypedName() plugins.TypedName {
	return f.typedName
}

// WithName sets the name of the filter.
func (f *LowQueueFilter) WithName(name string) *LowQueueFilter {
	f.typedName.Name = name
	return f
}

// Filter filters out pods that doesn't meet the filter criteria.
func (f *LowQueueFilter) Filter(_ context.Context, _ *types.CycleState, _ *types.LLMRequest, pods []types.Pod) []types.Pod {
	filteredPods := []types.Pod{}

	for _, pod := range pods {
		if pod.GetMetrics().WaitingQueueSize <= f.queueingThresholdLoRA {
			filteredPods = append(filteredPods, pod)
		}
	}

	return filteredPods
}
