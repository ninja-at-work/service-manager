/*
 * Copyright 2018 The Service Manager Authors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
package util

import (
	"bytes"
	"text/template"
)

// tsprintf stands for "template Sprintf" and fills the specified templated string with the provided data
func Tsprintf(tmpl string, data map[string]interface{}) (string, error) {
	t, err := template.New("sql").Parse(tmpl)
	if err != nil {
		return "", err
	}
	buff := &bytes.Buffer{}
	if err := t.Execute(buff, data); err != nil {
		return "", nil
	}
	return buff.String(), nil
}
