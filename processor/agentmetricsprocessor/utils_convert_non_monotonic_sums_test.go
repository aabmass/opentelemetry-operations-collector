// Copyright 2020, Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package agentmetricsprocessor

import (
	"go.opentelemetry.io/collector/consumer/pdata"
)

func generateNonMonotonicSumsInput() pdata.Metrics {
	input := pdata.NewMetrics()

	rmb := newResourceMetricsBuilder()
	b := rmb.addResourceMetrics(nil)

	mb1 := b.addMetric("m1", pdata.MetricDataTypeIntSum, false)
	mb1.addIntDataPoint(1, map[string]string{"label1": "value1"})
	mb1.addIntDataPoint(2, map[string]string{"label1": "value2"})

	mb2 := b.addMetric("m2", pdata.MetricDataTypeDoubleSum, false)
	mb2.addDoubleDataPoint(3, map[string]string{"label1": "value1"})
	mb2.addDoubleDataPoint(4, map[string]string{"label1": "value2"})

	rmb.Build().CopyTo(input.ResourceMetrics())
	return input
}

func generateNonMonotonicSumsExpected() pdata.Metrics {
	expected := pdata.NewMetrics()

	rmb := newResourceMetricsBuilder()
	b := rmb.addResourceMetrics(nil)

	mb1 := b.addMetric("m1", pdata.MetricDataTypeIntGauge, false)
	mb1.addIntDataPoint(1, map[string]string{"label1": "value1"})
	mb1.addIntDataPoint(2, map[string]string{"label1": "value2"})

	mb2 := b.addMetric("m2", pdata.MetricDataTypeDoubleGauge, false)
	mb2.addDoubleDataPoint(3, map[string]string{"label1": "value1"})
	mb2.addDoubleDataPoint(4, map[string]string{"label1": "value2"})

	rmb.Build().CopyTo(expected.ResourceMetrics())
	return expected
}
