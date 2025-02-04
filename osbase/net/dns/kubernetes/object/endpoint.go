// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package object

// Taken and modified from the Kubernetes plugin of CoreDNS, under Apache 2.0.

import (
	"fmt"
	"net/netip"
	"regexp"
	"slices"
	"strings"
	"time"

	api "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Endpoints is a stripped down discovery.EndpointSlice
// with only the items we need.
type Endpoints struct {
	Version               string
	Name                  string
	Namespace             string
	LastChangeTriggerTime time.Time
	Index                 string
	Addresses             []EndpointAddress
	Ports                 []Port

	*Empty
}

// EndpointAddress is a tuple that describes single IP address.
type EndpointAddress struct {
	// IP contains the IP address in binary format.
	IP       string
	Hostname string
}

// Port is a tuple that describes a single port.
type Port struct {
	Port     uint16
	Name     string
	Protocol string
}

// EndpointsKey returns a string using for the index.
func EndpointsKey(name, namespace string) string { return name + "." + namespace }

var hostnameRegexp = regexp.MustCompile(`^[-a-z0-9]{1,63}$`)

// EndpointSliceToEndpoints converts a *discovery.EndpointSlice to a *Endpoints.
func EndpointSliceToEndpoints(obj meta.Object) (meta.Object, error) {
	ends, ok := obj.(*discovery.EndpointSlice)
	if !ok {
		return nil, fmt.Errorf("unexpected object %v", obj)
	}
	e := &Endpoints{
		Version:   ends.GetResourceVersion(),
		Name:      ends.GetName(),
		Namespace: ends.GetNamespace(),
		Index:     EndpointsKey(ends.Labels[discovery.LabelServiceName], ends.GetNamespace()),
	}

	// In case of parse error, the value is time.Zero.
	e.LastChangeTriggerTime, _ = time.Parse(time.RFC3339Nano, ends.Annotations[api.EndpointsLastChangeTriggerTime])

	e.Ports = make([]Port, 0, len(ends.Ports))
	for _, p := range ends.Ports {
		if p.Port != nil && *p.Port >= 1 && *p.Port <= 0xffff &&
			p.Name != nil && *p.Name != "" && p.Protocol != nil {
			ep := Port{
				Port:     uint16(*p.Port),
				Name:     strings.ToLower(*p.Name),
				Protocol: strings.ToLower(string(*p.Protocol)),
			}
			e.Ports = append(e.Ports, ep)
		}
	}

	for _, end := range ends.Endpoints {
		if !endpointsliceReady(end.Conditions.Ready) {
			continue
		}

		var endHostname string
		if end.Hostname != nil {
			endHostname = *end.Hostname
		}
		if endHostname != "" && !hostnameRegexp.MatchString(endHostname) {
			endHostname = ""
		}

		for _, rawIP := range end.Addresses {
			parsedIP, err := netip.ParseAddr(rawIP)
			if err != nil || parsedIP.Zone() != "" {
				continue
			}
			parsedIP = parsedIP.Unmap()
			// The IP address is converted to a binary string, not human readable.
			// That way we don't need to parse it again later.
			ea := EndpointAddress{IP: string(parsedIP.AsSlice())}
			if endHostname != "" {
				ea.Hostname = endHostname
			} else {
				ea.Hostname = strings.ReplaceAll(strings.ReplaceAll(parsedIP.String(), ".", "-"), ":", "-")
			}
			e.Addresses = append(e.Addresses, ea)
		}
	}

	*ends = discovery.EndpointSlice{}

	return e, nil
}

func endpointsliceReady(ready *bool) bool {
	// Per API docs: a nil value indicates an unknown state. In most cases
	// consumers should interpret this unknown state as ready.
	if ready == nil {
		return true
	}
	return *ready
}

var _ runtime.Object = &Endpoints{}

// DeepCopyObject implements the ObjectKind interface.
func (e *Endpoints) DeepCopyObject() runtime.Object {
	e1 := &Endpoints{
		Version:   e.Version,
		Name:      e.Name,
		Namespace: e.Namespace,
		Index:     e.Index,
		Addresses: make([]EndpointAddress, len(e.Addresses)),
		Ports:     make([]Port, len(e.Ports)),
	}
	copy(e1.Addresses, e.Addresses)
	copy(e1.Ports, e.Ports)
	return e1
}

// GetNamespace implements the metav1.Object interface.
func (e *Endpoints) GetNamespace() string { return e.Namespace }

// SetNamespace implements the metav1.Object interface.
func (e *Endpoints) SetNamespace(namespace string) {}

// GetName implements the metav1.Object interface.
func (e *Endpoints) GetName() string { return e.Name }

// SetName implements the metav1.Object interface.
func (e *Endpoints) SetName(name string) {}

// GetResourceVersion implements the metav1.Object interface.
func (e *Endpoints) GetResourceVersion() string { return e.Version }

// SetResourceVersion implements the metav1.Object interface.
func (e *Endpoints) SetResourceVersion(version string) {}

// EndpointsModified checks if the update to an endpoint is something
// that matters to us or if they are effectively equivalent.
func EndpointsModified(a, b *Endpoints) bool {
	if a.Index != b.Index {
		return true
	}
	if !slices.Equal(a.Addresses, b.Addresses) {
		return true
	}
	if !slices.Equal(a.Ports, b.Ports) {
		return true
	}
	return false
}
