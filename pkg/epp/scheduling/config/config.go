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

package config

import (
	"sigs.k8s.io/controller-runtime/pkg/log"

	commonconfig "sigs.k8s.io/gateway-api-inference-extension/pkg/epp/common/config"
	envutil "sigs.k8s.io/gateway-api-inference-extension/pkg/epp/util/env"
	logutil "sigs.k8s.io/gateway-api-inference-extension/pkg/epp/util/logging"
)

// Config holds all the configuration values for the scheduler
type Config struct {
	KVCacheThreshold       float64
	QueueThresholdCritical int
	QueueingThresholdLoRA  int
	LoraAffinityThreshold  float64
}

const (
	// Default values for LoRA specific thresholds
	DefaultQueueingThresholdLoRA = 128
	DefaultLoraAffinityThreshold = 0.999
)

// LoadConfig loads configuration from environment variables
func LoadConfig() Config {
	// Use a default logger for initial configuration loading
	baseLogger := log.Log.WithName("scheduling-config")

	config := Config{
		KVCacheThreshold:       envutil.GetEnvFloat("KV_CACHE_THRESHOLD", commonconfig.DefaultKVCacheThreshold, baseLogger),
		QueueThresholdCritical: envutil.GetEnvInt("QUEUE_THRESHOLD_CRITICAL", commonconfig.DefaultQueueThresholdCritical, baseLogger),
		QueueingThresholdLoRA:  envutil.GetEnvInt("QUEUING_THRESHOLD_LORA", DefaultQueueingThresholdLoRA, baseLogger),
		LoraAffinityThreshold:  envutil.GetEnvFloat("LORA_AFFINITY_THRESHOLD", DefaultLoraAffinityThreshold, baseLogger),
	}

	baseLogger.V(logutil.DEFAULT).Info("Scheduler configuration loaded", "config", config)

	return config
}

var Conf = LoadConfig()
