// Copyright Istio Authors
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

package fixtures

import (
	"istio.io/istio/pkg/config"
	"istio.io/istio/pkg/config/analysis"
	"istio.io/istio/pkg/config/analysis/diag"
	"istio.io/istio/pkg/config/resource"
)

// Context is a test fixture of analysis.Context
type Context struct {
	Resources []*resource.Instance
	Reports   []diag.Message
}

var _ analysis.Context = &Context{}

// Report implements analysis.Context
func (ctx *Context) Report(_ config.GroupVersionKind, t diag.Message) {
	ctx.Reports = append(ctx.Reports, t)
}

// Find implements analysis.Context
func (ctx *Context) Find(config.GroupVersionKind, resource.FullName) *resource.Instance { return nil }

// Exists implements analysis.Context
func (ctx *Context) Exists(config.GroupVersionKind, resource.FullName) bool { return false }

// ForEach implements analysis.Context
func (ctx *Context) ForEach(_ config.GroupVersionKind, fn analysis.IteratorFn) {
	for _, r := range ctx.Resources {
		fn(r)
	}
}

// Canceled implements analysis.Context
func (ctx *Context) Canceled() bool { return false }

func (ctx *Context) SetAnalyzer(analyzerName string) {}
