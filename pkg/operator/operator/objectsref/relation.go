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
	"bytes"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/logger"
)

const (
	relationNodeSystem           = "node_with_system_relation"
	relationNodePod              = "node_with_pod_relation"
	relationJobPod               = "job_with_pod_relation"
	relationPodReplicaset        = "pod_with_replicaset_relation"
	relationPodStatefulset       = "pod_with_statefulset_relation"
	relationDaemonsetPod         = "daemonset_with_pod_relation"
	relationDeploymentReplicaset = "deployment_with_replicaset_relation"
	relationAddressService       = "address_with_service_relation"
	relationAddressPod           = "address_with_pod_relation"
)

type RelationMetric struct {
	Name   string
	Labels []RelationLabel
}

type RelationLabel struct {
	Name  string
	Value string
}

func (oc *ObjectsController) GetNodeRelations() []RelationMetric {
	var metrics []RelationMetric
	for node, ip := range oc.nodeObjs.Addrs() {
		metrics = append(metrics, RelationMetric{
			Name: relationNodeSystem,
			Labels: []RelationLabel{
				{Name: "node", Value: node},
				{Name: "bk_target_ip", Value: ip},
			},
		})
	}
	return metrics
}

func (oc *ObjectsController) GetEpSvcRelations() []RelationMetric {
	var metrics []RelationMetric
	appendAddrMetrics := func(namespace, svc string, addresses []addressEntity) {
		for _, addr := range addresses {
			for _, port := range addr.ports {
				metrics = append(metrics, RelationMetric{
					Name: relationAddressService,
					Labels: []RelationLabel{
						{Name: "namespace", Value: namespace},
						{Name: "service", Value: svc},
						{Name: "ip", Value: addr.ip},
						{Name: "port", Value: port.port},
						{Name: "protocol", Value: port.protocol},
					},
				})
			}
		}
	}

	oc.epsSvcObjs.rangeServices(func(namespace string, services serviceEntities) {
		for _, svc := range services {
			eps := oc.epsSvcObjs.endpoints[namespace]
			for _, ep := range eps {
				matched := matchLabels(svc.labels, ep.labels)
				if !matched {
					continue
				}
				appendAddrMetrics(namespace, svc.name, ep.addresses)
			}
		}
	})

	return metrics
}

func (oc *ObjectsController) GetAddressPodRelations() []RelationMetric {
	type svcPods struct {
		namespace string
		name      string
		pods      map[string]string //map[podname]podip
	}

	var lst []svcPods
	oc.epsSvcObjs.rangeServices(func(namespace string, services serviceEntities) {
		pods := oc.podObjs.GetByNamespace(namespace)
		for _, svc := range services {
			item := svcPods{namespace: namespace, name: svc.name, pods: map[string]string{}}
			for _, pod := range pods {
				if !matchLabels(svc.labels, pod.Labels) {
					continue
				}
				item.pods[pod.ID.Name] = pod.PodIP
			}
			lst = append(lst, item)
		}
	})

	var metrics []RelationMetric
	for _, item := range lst {
		logger.Errorf("mando:item: %+v", item)
		// endpoints <-> service 需要 namespace/name 保持对齐
		ep, ok := oc.epsSvcObjs.getEndpoints(item.namespace, item.name)
		if !ok {
			continue
		}
		logger.Errorf("mando:ep: %+v", ep)

		for _, addr := range ep.addresses {
			for pod, ip := range item.pods {
				if addr.ip != ip {
					continue
				}

				for _, port := range addr.ports {
					metrics = append(metrics, RelationMetric{
						Name: relationAddressPod,
						Labels: []RelationLabel{
							{Name: "namespace", Value: item.namespace},
							{Name: "pod", Value: pod},
							{Name: "ip", Value: addr.ip},
							{Name: "port", Value: port.port},
							{Name: "protocol", Value: port.protocol},
						},
					})
				}
			}
		}
	}

	return metrics
}

func (oc *ObjectsController) GetReplicasetRelations() []RelationMetric {
	var metrics []RelationMetric
	for _, rs := range oc.replicaSetObjs.GetAll() {
		ownerRef := LookupOnce(rs.ID, oc.replicaSetObjs, oc.objsMap())
		if ownerRef == nil {
			continue
		}

		labels := []RelationLabel{
			{Name: "namespace", Value: rs.ID.Namespace},
			{Name: "replicaset", Value: rs.ID.Name},
		}

		switch ownerRef.Kind {
		case kindDeployment:
			labels = append(labels, RelationLabel{
				Name:  "deployment",
				Value: ownerRef.Name,
			})
			metrics = append(metrics, RelationMetric{
				Name:   relationDeploymentReplicaset,
				Labels: labels,
			})
		}
	}
	return metrics
}

func (oc *ObjectsController) GetPodRelations() []RelationMetric {
	var metrics []RelationMetric
	for _, pod := range oc.podObjs.GetAll() {
		ownerRef := LookupOnce(pod.ID, oc.podObjs, oc.objsMap())
		if ownerRef == nil {
			continue
		}

		metrics = append(metrics, RelationMetric{
			Name: relationNodePod,
			Labels: []RelationLabel{
				{Name: "namespace", Value: pod.ID.Namespace},
				{Name: "pod", Value: pod.ID.Name},
				{Name: "node", Value: pod.NodeName},
			},
		})

		labels := []RelationLabel{
			{Name: "namespace", Value: pod.ID.Namespace},
			{Name: "pod", Value: pod.ID.Name},
		}
		switch ownerRef.Kind {
		case kindJob:
			labels = append(labels, RelationLabel{
				Name:  "job",
				Value: ownerRef.Name,
			})
			metrics = append(metrics, RelationMetric{
				Name:   relationJobPod,
				Labels: labels,
			})

		case kindReplicaSet:
			labels = append(labels, RelationLabel{
				Name:  "replicaset",
				Value: ownerRef.Name,
			})
			metrics = append(metrics, RelationMetric{
				Name:   relationPodReplicaset,
				Labels: labels,
			})

		case kindGameStatefulSet:
			labels = append(labels, RelationLabel{
				Name:  "statefulset",
				Value: ownerRef.Name,
			})
			metrics = append(metrics, RelationMetric{
				Name:   relationPodStatefulset,
				Labels: labels,
			})

		case kindDaemonSet:
			labels = append(labels, RelationLabel{
				Name:  "daemonset",
				Value: ownerRef.Name,
			})
			metrics = append(metrics, RelationMetric{
				Name:   relationDaemonsetPod,
				Labels: labels,
			})
		}
	}
	return metrics
}

func RelationToPromFormat(metrics []RelationMetric) []byte {
	var lines []byte
	for _, metric := range metrics {
		var buf bytes.Buffer
		buf.WriteString(metric.Name)
		buf.WriteString(`{`)

		var n int
		for _, label := range metric.Labels {
			if n > 0 {
				buf.WriteString(`,`)
			}
			n++
			buf.WriteString(label.Name)
			buf.WriteString(`="`)
			buf.WriteString(label.Value)
			buf.WriteString(`"`)
		}

		buf.WriteString("} 1")
		buf.WriteString("\n")
		lines = append(lines, buf.Bytes()...)
	}
	return lines
}
