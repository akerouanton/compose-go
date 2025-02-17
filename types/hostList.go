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

package types

import (
	"encoding/json"
	"fmt"
	"sort"
)

// HostsList is a list of colon-separated host-ip mappings
type HostsList map[string]string

// AsList return host-ip mappings as a list of colon-separated strings
func (h HostsList) AsList() []string {
	l := make([]string, 0, len(h))
	for k, v := range h {
		l = append(l, fmt.Sprintf("%s:%s", k, v))
	}
	return l
}

func (h HostsList) MarshalYAML() (interface{}, error) {
	list := h.AsList()
	sort.Strings(list)
	return list, nil
}

func (h HostsList) MarshalJSON() ([]byte, error) {
	list := h.AsList()
	sort.Strings(list)
	return json.Marshal(list)
}

func (h *HostsList) DecodeMapstructure(value interface{}) error {
	switch v := value.(type) {
	case map[string]interface{}:
		list := make(HostsList, len(v))
		for i, e := range v {
			if e == nil {
				e = ""
			}
			list[i] = fmt.Sprint(e)
		}
		*h = list
	case []interface{}:
		*h = decodeMapping(v, ":")
	default:
		return fmt.Errorf("unexpected value type %T for mapping", value)
	}
	return nil
}
