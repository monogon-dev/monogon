package kubernetes

// Taken and modified from the Kubernetes plugin of CoreDNS, under Apache 2.0.

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	api "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"source.monogon.dev/osbase/net/dns/kubernetes/object"
)

const (
	svcIPIndex           = "ServiceIP"
	epNameNamespaceIndex = "EndpointNameNamespace"
	epIPIndex            = "EndpointsIP"
)

// epLabelSelector selects EndpointSlices that belong to headless services.
// Endpoint DNS records are only served for headless services,
// and we can save resources by not fetching other EndpointSlices.
const epLabelSelector = api.IsHeadlessService

type dnsController interface {
	GetSvc(string) *object.Service
	SvcIndexReverse(string) []*object.Service
	EpIndex(string) []*object.Endpoints
	EpIndexReverse(string) []*object.Endpoints
	NamespaceExists(string) bool

	Start(<-chan struct{})
	HasSynced() bool

	// Modified returns the timestamp of the most recent changes to services.
	Modified() int64
}

type dnsControl struct {
	// modified tracks timestamp of the most recent changes
	modified  atomic.Int64
	hasSynced atomic.Bool

	svcController cache.Controller
	epController  cache.Controller
	nsController  cache.Controller

	svcLister cache.Indexer
	epLister  cache.Indexer
	nsLister  cache.Store
}

// newdnsController creates a controller.
func newdnsController(ctx context.Context, kubeClient kubernetes.Interface) *dnsControl {
	dns := dnsControl{}

	dns.svcLister, dns.svcController = object.NewIndexerInformer(
		&cache.ListWatch{
			ListFunc:  serviceListFunc(ctx, kubeClient, api.NamespaceAll),
			WatchFunc: serviceWatchFunc(ctx, kubeClient, api.NamespaceAll),
		},
		&api.Service{},
		&dns,
		cache.Indexers{svcIPIndex: svcIPIndexFunc},
		object.DefaultProcessor(object.ToService),
	)

	dns.epLister, dns.epController = object.NewIndexerInformer(
		&cache.ListWatch{
			ListFunc:  endpointSliceListFunc(ctx, kubeClient, api.NamespaceAll),
			WatchFunc: endpointSliceWatchFunc(ctx, kubeClient, api.NamespaceAll),
		},
		&discovery.EndpointSlice{},
		&dns,
		cache.Indexers{epNameNamespaceIndex: epNameNamespaceIndexFunc, epIPIndex: epIPIndexFunc},
		object.DefaultProcessor(object.EndpointSliceToEndpoints),
	)

	dns.nsLister, dns.nsController = object.NewIndexerInformer(
		&cache.ListWatch{
			ListFunc:  namespaceListFunc(ctx, kubeClient),
			WatchFunc: namespaceWatchFunc(ctx, kubeClient),
		},
		&api.Namespace{},
		&dns,
		cache.Indexers{},
		object.DefaultProcessor(object.ToNamespace),
	)

	return &dns
}

func svcIPIndexFunc(obj interface{}) ([]string, error) {
	svc, ok := obj.(*object.Service)
	if !ok {
		return nil, errObj
	}
	idx := make([]string, len(svc.ClusterIPs))
	copy(idx, svc.ClusterIPs)
	return idx, nil
}

func epNameNamespaceIndexFunc(obj interface{}) ([]string, error) {
	s, ok := obj.(*object.Endpoints)
	if !ok {
		return nil, errObj
	}
	return []string{s.Index}, nil
}

func epIPIndexFunc(obj interface{}) ([]string, error) {
	ep, ok := obj.(*object.Endpoints)
	if !ok {
		return nil, errObj
	}
	idx := make([]string, len(ep.Addresses))
	for i, addr := range ep.Addresses {
		idx[i] = addr.IP
	}
	return idx, nil
}

func serviceListFunc(ctx context.Context, c kubernetes.Interface, ns string) func(meta.ListOptions) (runtime.Object, error) {
	return func(opts meta.ListOptions) (runtime.Object, error) {
		return c.CoreV1().Services(ns).List(ctx, opts)
	}
}

func endpointSliceListFunc(ctx context.Context, c kubernetes.Interface, ns string) func(meta.ListOptions) (runtime.Object, error) {
	return func(opts meta.ListOptions) (runtime.Object, error) {
		opts.LabelSelector = epLabelSelector
		return c.DiscoveryV1().EndpointSlices(ns).List(ctx, opts)
	}
}

func namespaceListFunc(ctx context.Context, c kubernetes.Interface) func(meta.ListOptions) (runtime.Object, error) {
	return func(opts meta.ListOptions) (runtime.Object, error) {
		return c.CoreV1().Namespaces().List(ctx, opts)
	}
}

func serviceWatchFunc(ctx context.Context, c kubernetes.Interface, ns string) func(options meta.ListOptions) (watch.Interface, error) {
	return func(options meta.ListOptions) (watch.Interface, error) {
		return c.CoreV1().Services(ns).Watch(ctx, options)
	}
}

func endpointSliceWatchFunc(ctx context.Context, c kubernetes.Interface, ns string) func(options meta.ListOptions) (watch.Interface, error) {
	return func(options meta.ListOptions) (watch.Interface, error) {
		options.LabelSelector = epLabelSelector
		return c.DiscoveryV1().EndpointSlices(ns).Watch(ctx, options)
	}
}

func namespaceWatchFunc(ctx context.Context, c kubernetes.Interface) func(options meta.ListOptions) (watch.Interface, error) {
	return func(options meta.ListOptions) (watch.Interface, error) {
		return c.CoreV1().Namespaces().Watch(ctx, options)
	}
}

// Start starts the controller.
func (dns *dnsControl) Start(stopCh <-chan struct{}) {
	go dns.svcController.Run(stopCh)
	go dns.epController.Run(stopCh)
	go dns.nsController.Run(stopCh)
}

// HasSynced returns true if the initial data has been
// completely loaded into memory.
func (dns *dnsControl) HasSynced() bool {
	if dns.hasSynced.Load() {
		return true
	}
	a := dns.svcController.HasSynced()
	b := dns.epController.HasSynced()
	c := dns.nsController.HasSynced()
	if a && b && c {
		dns.hasSynced.Store(true)
		return true
	}
	return false
}

func (dns *dnsControl) GetSvc(key string) *object.Service {
	o, exists, err := dns.svcLister.GetByKey(key)
	if err != nil || !exists {
		return nil
	}
	s, ok := o.(*object.Service)
	if ok {
		return s
	}
	return nil
}

func (dns *dnsControl) SvcIndexReverse(ip string) (svcs []*object.Service) {
	os, err := dns.svcLister.ByIndex(svcIPIndex, ip)
	if err != nil {
		return nil
	}
	svcs = make([]*object.Service, 0, len(os))
	for _, o := range os {
		s, ok := o.(*object.Service)
		if ok {
			svcs = append(svcs, s)
		}
	}
	return svcs
}

func (dns *dnsControl) EpIndex(idx string) (ep []*object.Endpoints) {
	os, err := dns.epLister.ByIndex(epNameNamespaceIndex, idx)
	if err != nil {
		return nil
	}
	ep = make([]*object.Endpoints, 0, len(os))
	for _, o := range os {
		e, ok := o.(*object.Endpoints)
		if ok {
			ep = append(ep, e)
		}
	}
	return ep
}

func (dns *dnsControl) EpIndexReverse(ip string) (ep []*object.Endpoints) {
	os, err := dns.epLister.ByIndex(epIPIndex, ip)
	if err != nil {
		return nil
	}
	ep = make([]*object.Endpoints, 0, len(os))
	for _, o := range os {
		e, ok := o.(*object.Endpoints)
		if ok {
			ep = append(ep, e)
		}
	}
	return ep
}

// NamespaceExists returns true if a namespace with this name exists.
func (dns *dnsControl) NamespaceExists(name string) bool {
	_, exists, _ := dns.nsLister.GetByKey(name)
	return exists
}

func (dns *dnsControl) OnAdd(obj interface{}, isInInitialList bool) {
	dns.updateModified()
	switch obj := obj.(type) {
	case *object.Endpoints:
		// Don't record latency during initial sync, because measuring latency only
		// makes sense for changes that happen while the service is running.
		if !isInInitialList {
			recordDNSProgrammingLatency(obj.LastChangeTriggerTime)
		}
	}
}

func (dns *dnsControl) OnDelete(obj interface{}) {
	dns.updateModified()
	// Note: We cannot record programming latency on deletes, because the trigger
	// time annotation is not updated when the object is deleted.
}

func (dns *dnsControl) OnUpdate(oldObj, newObj interface{}) {
	// If both objects have the same resource version, they are identical.
	if oldObj.(meta.Object).GetResourceVersion() == newObj.(meta.Object).GetResourceVersion() {
		return
	}

	switch newObj := newObj.(type) {
	case *object.Service:
		if object.ServiceModified(oldObj.(*object.Service), newObj) {
			dns.updateModified()
		}
	case *object.Endpoints:
		oldObj := oldObj.(*object.Endpoints)
		if object.EndpointsModified(oldObj, newObj) {
			dns.updateModified()
			// If the trigger time has not changed, the process that last updated the
			// object did not update the trigger time, so we can't know the latency.
			if oldObj.LastChangeTriggerTime != newObj.LastChangeTriggerTime {
				recordDNSProgrammingLatency(newObj.LastChangeTriggerTime)
			}
		}
	}
}

func (dns *dnsControl) Modified() int64 {
	return dns.modified.Load()
}

// updateModified set dns.modified to the current time.
func (dns *dnsControl) updateModified() {
	unix := time.Now().Unix()
	dns.modified.Store(unix)
}

var errObj = errors.New("obj was not of the correct type")
