/*
Copyright 2023 The Kubernetes Authors.

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

package localstorage

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/drainability"
	"k8s.io/autoscaler/cluster-autoscaler/utils/drain"
	pod_util "k8s.io/autoscaler/cluster-autoscaler/utils/pod"
)

// Rule is a drainability rule on how to handle local storage pods.
type Rule struct {
	enabled bool
}

// New creates a new Rule.
func New(enabled bool) *Rule {
	return &Rule{
		enabled: enabled,
	}
}

// Drainable decides what to do with local storage pods on node drain.
func (r *Rule) Drainable(drainCtx *drainability.DrainContext, pod *apiv1.Pod) drainability.Status {
	if r.enabled && !drain.IsPodLongTerminating(pod, drainCtx.Timestamp) && !pod_util.IsDaemonSetPod(pod) && !drain.HasSafeToEvictAnnotation(pod) && !drain.IsPodTerminal(pod) && drain.HasBlockingLocalStorage(pod) {
		return drainability.NewBlockedStatus(drain.LocalStorageRequested, fmt.Errorf("pod with local storage present: %s", pod.Name))
	}
	return drainability.NewUndefinedStatus()
}
