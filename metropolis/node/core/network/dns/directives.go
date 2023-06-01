// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dns

import (
	"fmt"
	"net"
	"strings"
)

// Type ExtraDirective contains additional config directives for CoreDNS.
type ExtraDirective struct {
	// ID is the identifier of this directive. There can only be one directive
	// with a given ID active at once. The ID is also used to identify which
	// directive to purge.
	ID string
	// directive contains a full CoreDNS directive as a string. It can also use
	// the $FILE(<filename>) macro, which will be expanded to the path of a
	// file from the files field.
	directive string
	// files contains additional files used in the configuration. The map key
	// is used as the filename.
	files map[string][]byte
}

// NewUpstreamDirective creates a forward with no fallthrough that forwards all
// requests not yet matched to the given upstream DNS servers.
func NewUpstreamDirective(dnsServers []net.IP) *ExtraDirective {
	strb := strings.Builder{}
	if len(dnsServers) > 0 {
		strb.WriteString("forward .")
		for _, ip := range dnsServers {
			strb.WriteString(" ")
			strb.WriteString(ip.String())
		}
	}
	return &ExtraDirective{
		directive: strb.String(),
	}
}

var kubernetesDirective = `
kubernetes %v in-addr.arpa ip6.arpa {
	kubeconfig $FILE(kubeconfig) default
	pods insecure
	fallthrough in-addr.arpa ip6.arpa
	ttl 30
}
`

// NewKubernetesDirective creates a directive running a "Kubernetes DNS-Based
// Service Discovery" compliant service under clusterDomain. The given
// kubeconfig needs at least read access to services, endpoints and
// endpointslices.
//
//	[1] https://github.com/kubernetes/dns/blob/master/docs/specification.md
func NewKubernetesDirective(clusterDomain string, kubeconfig []byte) *ExtraDirective {
	var prefix string
	return &ExtraDirective{
		ID:        "k8s-clusterdns",
		directive: prefix + fmt.Sprintf(kubernetesDirective, clusterDomain),
		files: map[string][]byte{
			"kubeconfig": kubeconfig,
		},
	}
}
