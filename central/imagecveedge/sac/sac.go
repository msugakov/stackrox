package sac

import (
	"github.com/stackrox/stackrox/central/dackbox"
	"github.com/stackrox/stackrox/central/role/resources"
	"github.com/stackrox/stackrox/pkg/sac"
	"github.com/stackrox/stackrox/pkg/search/filtered"
)

var (
	imageSAC = sac.ForResource(resources.Image)

	imageCVEEdgeSACFilter = filtered.MustCreateNewSACFilter(
		filtered.WithResourceHelper(imageSAC),
		filtered.WithScopeTransform(dackbox.ImageCVEEdgeSACTransform),
		filtered.WithReadAccess())
)

// GetSACFilter returns the sac filter for image cve edge ids.
func GetSACFilter() filtered.Filter {
	return imageCVEEdgeSACFilter
}
