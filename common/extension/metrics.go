/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package extension

import (
	"github.com/apache/dubbo-go/metrics"
)

var (
	// can not declare the map as map[string]*MetricManager which causes cycle dependency
	metricManagerMap = make(map[string]metrics.MetricManager, 4)
)

// the manager should be a pointer of MetricManager.
func SetMetricManager(name string, manager metrics.MetricManager) {
	metricManagerMap[name] = manager
}

func GetMetricManager(name string) metrics.MetricManager {
	manager, found := metricManagerMap[name]
	if !found {
		panic("Can not find the metric manager with name: " + name +
			", please check that the MetricManager had registered by invoking SetMetricManager. ")
	}
	return manager
}