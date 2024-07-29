package object

// Taken and modified from the Kubernetes plugin of CoreDNS, under Apache 2.0.

import (
	"fmt"
	"net/netip"
	"regexp"
	"slices"
	"strings"

	"github.com/miekg/dns"
	api "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Service is a stripped down api.Service with only the items we need.
type Service struct {
	Version   string
	Name      string
	Namespace string
	// ClusterIPs contains IP addresses in binary format.
	ClusterIPs   []string
	ExternalName string
	Ports        []Port
	Headless     bool

	*Empty
}

var domainNameRegexp = regexp.MustCompile(`^([-a-z0-9]{1,63}\.)+$`)

const ExternalNameInvalid = "."

// ToService converts an api.Service to a *Service.
func ToService(obj meta.Object) (meta.Object, error) {
	svc, ok := obj.(*api.Service)
	if !ok {
		return nil, fmt.Errorf("unexpected object %v", obj)
	}

	s := &Service{
		Version:   svc.GetResourceVersion(),
		Name:      svc.GetName(),
		Namespace: svc.GetNamespace(),
	}

	if svc.Spec.Type == api.ServiceTypeExternalName {
		// Make the name fully qualified.
		externalName := dns.Fqdn(svc.Spec.ExternalName)
		// Check if the name is valid. Even names that pass Kubernetes validation
		// can fail this check, because Kubernetes does not validate that labels
		// must be at most 63 characters.
		if !domainNameRegexp.MatchString(externalName) || len(externalName) > 254 {
			externalName = ExternalNameInvalid
		}
		s.ExternalName = externalName
	} else {
		if svc.Spec.ClusterIP == api.ClusterIPNone {
			s.Headless = true
		} else {
			s.ClusterIPs = make([]string, 0, len(svc.Spec.ClusterIPs))
			for _, rawIP := range svc.Spec.ClusterIPs {
				parsedIP, err := netip.ParseAddr(rawIP)
				if err != nil || parsedIP.Zone() != "" {
					continue
				}
				parsedIP = parsedIP.Unmap()
				s.ClusterIPs = append(s.ClusterIPs, string(parsedIP.AsSlice()))
			}

			s.Ports = make([]Port, 0, len(svc.Spec.Ports))
			for _, p := range svc.Spec.Ports {
				if p.Port >= 1 && p.Port <= 0xffff && p.Name != "" {
					ep := Port{
						Port:     uint16(p.Port),
						Name:     strings.ToLower(p.Name),
						Protocol: strings.ToLower(string(p.Protocol)),
					}
					s.Ports = append(s.Ports, ep)
				}
			}
		}
	}

	*svc = api.Service{}

	return s, nil
}

var _ runtime.Object = &Service{}

// DeepCopyObject implements the ObjectKind interface.
func (s *Service) DeepCopyObject() runtime.Object {
	s1 := &Service{
		Version:      s.Version,
		Name:         s.Name,
		Namespace:    s.Namespace,
		ClusterIPs:   make([]string, len(s.ClusterIPs)),
		ExternalName: s.ExternalName,
		Ports:        make([]Port, len(s.Ports)),
		Headless:     s.Headless,
	}
	copy(s1.ClusterIPs, s.ClusterIPs)
	copy(s1.Ports, s.Ports)
	return s1
}

// GetNamespace implements the metav1.Object interface.
func (s *Service) GetNamespace() string { return s.Namespace }

// SetNamespace implements the metav1.Object interface.
func (s *Service) SetNamespace(namespace string) {}

// GetName implements the metav1.Object interface.
func (s *Service) GetName() string { return s.Name }

// SetName implements the metav1.Object interface.
func (s *Service) SetName(name string) {}

// GetResourceVersion implements the metav1.Object interface.
func (s *Service) GetResourceVersion() string { return s.Version }

// SetResourceVersion implements the metav1.Object interface.
func (s *Service) SetResourceVersion(version string) {}

// ServiceModified checks if the update to a service is something
// that matters to us or if they are effectively equivalent.
func ServiceModified(oldSvc, newSvc *Service) bool {
	if oldSvc.ExternalName != newSvc.ExternalName {
		return true
	}
	if oldSvc.Headless != newSvc.Headless {
		return true
	}
	if !slices.Equal(oldSvc.ClusterIPs, newSvc.ClusterIPs) {
		return true
	}
	if !slices.Equal(oldSvc.Ports, newSvc.Ports) {
		return true
	}
	return false
}
