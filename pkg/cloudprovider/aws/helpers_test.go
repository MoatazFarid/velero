/*
Copyright 2018 the Velero contributors.

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
package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS3URL(t *testing.T) {
	assert.True(t, IsValidS3URLScheme("http://foo"))
	assert.True(t, IsValidS3URLScheme("https://foo"))
	assert.False(t, IsValidS3URLScheme("httpd://foo"))
	assert.False(t, IsValidS3URLScheme(""))
}
