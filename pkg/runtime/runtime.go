// Copyright 2020 Google LLC
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

// Package runtime is used to perform general runtime actions.
package runtime

import (
	"fmt"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/buildpacks/pkg/env"
	gcp "github.com/GoogleCloudPlatform/buildpacks/pkg/gcpbuildpack"
)

// CheckOverride returns a Detect result or nil based on the GOOGLE_RUNTIME environment variable value.
//
// o If the GOOGLE_RUNTIME environment variable is not set or set to an empty string this returns nil.
//   Indicates an open source build with no user provided env var value set.
// o If the GOOGLE_RUNTIME environment variable is set to a value that matches the wantRuntime value
//   this returns an OptIn result. Indicates an open source build with a specified value set
//   by the user.
// o If the GOOGLE_RUNTIME environment variable is set to a value that starts with the wantRuntime
//   value this returns an OptIn reult. Indicates a gae or gcf build where the environment value
//   is set to a Google runtime name such ase 'python37' which is supported by the buildpack
//   performing detection.
// o If the GOOGLE_RUNTIME environment variable is set to another value this returns an OptOut result..
//   Indicates a gae or gcf build and the runtime needed for the build is not supported by the
//   buildpack performing detection.
func CheckOverride(wantRuntime string) gcp.DetectResult {
	envRuntime := strings.ToLower(strings.TrimSpace(os.Getenv(env.Runtime)))
	if envRuntime == "" {
		return nil
	}

	if strings.HasPrefix(envRuntime, wantRuntime) {
		return gcp.OptIn(fmt.Sprintf("%s  matches %q", env.Runtime, wantRuntime))
	}
	return gcp.OptOut(fmt.Sprintf("%s does not match to %q", env.Runtime, wantRuntime))
}
