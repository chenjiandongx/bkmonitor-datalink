// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package objectsref

import (
	"strconv"
	"sync"

	corev1 "k8s.io/api/core/v1"
)

type portEntity struct {
	port     string
	protocol string
}

type addressEntity struct {
	ip    string
	ports []portEntity
}

type endpointsEntity struct {
	name      string
	namespace string
	addresses []addressEntity
	labels    map[string]string
}

type endpointsEntities map[string]endpointsEntity

type serviceEntity struct {
	name   string
	labels map[string]string
}

type serviceEntities map[string]serviceEntity

type EpsSvcMap struct {
	mut       sync.Mutex
	endpoints map[string]endpointsEntities
	services  map[string]serviceEntities
}

func NewEpsSvcMap() *EpsSvcMap {
	return &EpsSvcMap{
		endpoints: map[string]endpointsEntities{},
		services:  map[string]serviceEntities{},
	}
}

func (m *EpsSvcMap) UpsertEndpoints(endpoints *corev1.Endpoints) {
	m.mut.Lock()
	defer m.mut.Unlock()

	if _, ok := m.endpoints[endpoints.Namespace]; !ok {
		m.endpoints[endpoints.Namespace] = make(endpointsEntities)
	}

	var addressEntities []addressEntity
	appendEntities := func(address []corev1.EndpointAddress, ports []corev1.EndpointPort) {
		for _, addr := range address {
			var portEntities []portEntity
			for _, port := range ports {
				portEntities = append(portEntities, portEntity{
					port:     strconv.Itoa(int(port.Port)),
					protocol: string(port.Protocol),
				})
			}
			addressEntities = append(addressEntities, addressEntity{
				ip:    addr.IP,
				ports: portEntities,
			})
		}
	}

	for _, subset := range endpoints.Subsets {
		appendEntities(subset.Addresses, subset.Ports)
		appendEntities(subset.NotReadyAddresses, subset.Ports)
	}

	m.endpoints[endpoints.Namespace][endpoints.Name] = endpointsEntity{
		name:      endpoints.Name,
		namespace: endpoints.Namespace,
		addresses: addressEntities,
		labels:    endpoints.Labels,
	}
}

func (m *EpsSvcMap) DeleteEndpoints(endpoints *corev1.Endpoints) {
	m.mut.Lock()
	defer m.mut.Unlock()

	if objs, ok := m.endpoints[endpoints.Namespace]; ok {
		delete(objs, endpoints.Name)
	}
}

func (m *EpsSvcMap) UpsertService(service *corev1.Service) {
	m.mut.Lock()
	defer m.mut.Unlock()

	if _, ok := m.services[service.Namespace]; !ok {
		m.services[service.Namespace] = make(serviceEntities)
	}

	m.services[service.Namespace][service.Name] = serviceEntity{
		name:   service.Name,
		labels: service.Labels,
	}
}

func (m *EpsSvcMap) DeleteService(service *corev1.Service) {
	m.mut.Lock()
	defer m.mut.Unlock()

	if objs, ok := m.services[service.Namespace]; ok {
		delete(objs, service.Name)
	}
}

func matchLabels(subset, set map[string]string) bool {
	for k, v := range subset {
		val, ok := set[k]
		if !ok || val != v {
			return false
		}
	}
	return true
}
