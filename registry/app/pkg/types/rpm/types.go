//  Copyright 2023 Harness, Inc.
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

package rpm

import (
	"github.com/harness/gitness/registry/app/metadata/rpm"
	"github.com/harness/gitness/registry/app/pkg"
)

type ArtifactInfo struct {
	pkg.ArtifactInfo
	Version     string
	Arch        string
	FileName    string
	PackagePath string
	Metadata    rpm.Metadata
}

func (a ArtifactInfo) GetVersion() string {
	return a.Version
}

// BaseArtifactInfo implements pkg.PackageArtifactInfo interface.
func (a ArtifactInfo) BaseArtifactInfo() pkg.ArtifactInfo {
	return a.ArtifactInfo
}

func (a ArtifactInfo) GetImageVersion() (exists bool, imageVersion string) {
	if a.Image != "" && a.Version != "" {
		return true, pkg.JoinWithSeparator(":", a.Image, a.Version)
	}
	return false, ""
}

func (a ArtifactInfo) GetFileName() string {
	return a.FileName
}

type File struct {
	FileURL string
	Name    string
}

type PackageMetadata struct {
	Name  string
	Files []File
}
