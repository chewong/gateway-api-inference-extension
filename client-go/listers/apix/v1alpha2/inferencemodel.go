/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	apixv1alpha2 "sigs.k8s.io/gateway-api-inference-extension/apix/v1alpha2"
)

// InferenceModelLister helps list InferenceModels.
// All objects returned here must be treated as read-only.
type InferenceModelLister interface {
	// List lists all InferenceModels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apixv1alpha2.InferenceModel, err error)
	// InferenceModels returns an object that can list and get InferenceModels.
	InferenceModels(namespace string) InferenceModelNamespaceLister
	InferenceModelListerExpansion
}

// inferenceModelLister implements the InferenceModelLister interface.
type inferenceModelLister struct {
	listers.ResourceIndexer[*apixv1alpha2.InferenceModel]
}

// NewInferenceModelLister returns a new InferenceModelLister.
func NewInferenceModelLister(indexer cache.Indexer) InferenceModelLister {
	return &inferenceModelLister{listers.New[*apixv1alpha2.InferenceModel](indexer, apixv1alpha2.Resource("inferencemodel"))}
}

// InferenceModels returns an object that can list and get InferenceModels.
func (s *inferenceModelLister) InferenceModels(namespace string) InferenceModelNamespaceLister {
	return inferenceModelNamespaceLister{listers.NewNamespaced[*apixv1alpha2.InferenceModel](s.ResourceIndexer, namespace)}
}

// InferenceModelNamespaceLister helps list and get InferenceModels.
// All objects returned here must be treated as read-only.
type InferenceModelNamespaceLister interface {
	// List lists all InferenceModels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apixv1alpha2.InferenceModel, err error)
	// Get retrieves the InferenceModel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*apixv1alpha2.InferenceModel, error)
	InferenceModelNamespaceListerExpansion
}

// inferenceModelNamespaceLister implements the InferenceModelNamespaceLister
// interface.
type inferenceModelNamespaceLister struct {
	listers.ResourceIndexer[*apixv1alpha2.InferenceModel]
}
