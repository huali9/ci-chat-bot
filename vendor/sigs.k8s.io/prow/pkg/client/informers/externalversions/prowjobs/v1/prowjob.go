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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	prowjobsv1 "sigs.k8s.io/prow/pkg/apis/prowjobs/v1"
	versioned "sigs.k8s.io/prow/pkg/client/clientset/versioned"
	internalinterfaces "sigs.k8s.io/prow/pkg/client/informers/externalversions/internalinterfaces"
	v1 "sigs.k8s.io/prow/pkg/client/listers/prowjobs/v1"
)

// ProwJobInformer provides access to a shared informer and lister for
// ProwJobs.
type ProwJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProwJobLister
}

type prowJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProwJobInformer constructs a new informer for ProwJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProwJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProwJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProwJobInformer constructs a new informer for ProwJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProwJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProwV1().ProwJobs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProwV1().ProwJobs(namespace).Watch(context.TODO(), options)
			},
		},
		&prowjobsv1.ProwJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *prowJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProwJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *prowJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&prowjobsv1.ProwJob{}, f.defaultInformer)
}

func (f *prowJobInformer) Lister() v1.ProwJobLister {
	return v1.NewProwJobLister(f.Informer().GetIndexer())
}
