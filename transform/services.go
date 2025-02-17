/*
   Copyright 2020 The Compose Specification Authors.

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

package transform

import (
	"github.com/compose-spec/compose-go/v2/tree"
)

func transformService(data any, p tree.Path) (any, error) {
	value := data.(map[string]any)
	if _, ok := value["scale"]; !ok {
		value["scale"] = 1 // TODO(ndeloof) we should make scale a *int
	}
	return transformMapping(value, p)
}

func transformServiceNetworks(data any, p tree.Path) (any, error) {
	if slice, ok := data.([]any); ok {
		networks := make(map[string]any, len(slice))
		for _, net := range slice {
			networks[net.(string)] = nil
		}
		return networks, nil
	}
	return data, nil
}
