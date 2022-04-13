package mappings

import (
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/search"
)

// OptionsMap defines the search options for Image.
var OptionsMap = search.Walk(v1.SearchCategory_IMAGES, "image", (*storage.Image)(nil))

// VulnerabilityOptionsMap defines the search options for Vulnerabilities stored in images.
var VulnerabilityOptionsMap = search.Walk(v1.SearchCategory_VULNERABILITIES, "image.scan.components.vulns", (*storage.EmbeddedVulnerability)(nil))

// ComponentOptionsMap defines the search options for image components stored in images.
var ComponentOptionsMap = search.Walk(v1.SearchCategory_IMAGE_COMPONENTS, "image.scan.components", (*storage.EmbeddedImageScanComponent)(nil))
