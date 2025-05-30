/*
Copyright 2016 The Kubernetes Authors.

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

/*
Adapted from k8s.io/apimachinery@v0.31.2/pkg/selection/operator.go

We're adding partial-match operators ~ and !~
*/

package selection

// Operator represents a key/field's relationship to value(s).
// See labels.Requirement and fields.Requirement for more details.
type Operator string

const (
	DoesNotExist     Operator = "!"
	Equals           Operator = "="
	DoubleEquals     Operator = "=="
	PartialEquals    Operator = "~"
	In               Operator = "in"
	NotEquals        Operator = "!="
	NotPartialEquals Operator = "!~"
	NotIn            Operator = "notin"
	Exists           Operator = "exists"
	GreaterThan      Operator = "gt"
	LessThan         Operator = "lt"
)
