package mapping

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/custom"
	_ "github.com/blevesearch/bleve/analysis/analyzer/keyword"  // Import the keyword analyzer so that it can be referred to from proto files
	_ "github.com/blevesearch/bleve/analysis/analyzer/standard" // Import the standard analyzer so that it can be referred to from proto files
	"github.com/blevesearch/bleve/analysis/token/lowercase"
	"github.com/blevesearch/bleve/analysis/tokenizer/whitespace"
	"github.com/blevesearch/bleve/mapping"
	activeComponentMappings "github.com/stackrox/rox/central/activecomponent/datastore/index/mappings"
	alertMapping "github.com/stackrox/rox/central/alert/mappings"
	clusterMapping "github.com/stackrox/rox/central/cluster/index/mappings"
	"github.com/stackrox/rox/central/compliance/standards/index"
	componentVulnEdgeMapping "github.com/stackrox/rox/central/componentcveedge/mappings"
	cveMapping "github.com/stackrox/rox/central/cve/mappings"
	imageComponentMapping "github.com/stackrox/rox/central/imagecomponent/mappings"
	imageComponentEdgeMapping "github.com/stackrox/rox/central/imagecomponentedge/mappings"
	imageCVEEdgeMapping "github.com/stackrox/rox/central/imagecveedge/mappings"
	imageIntegrationMapping "github.com/stackrox/rox/central/imageintegration/index/mappings"
	namespaceMapping "github.com/stackrox/rox/central/namespace/index/mappings"
	nodeMapping "github.com/stackrox/rox/central/node/index/mappings"
	nodeComponentEdgeMapping "github.com/stackrox/rox/central/nodecomponentedge/mappings"
	podMapping "github.com/stackrox/rox/central/pod/mappings"
	policyMapping "github.com/stackrox/rox/central/policy/index/mappings"
	policyCategoryMapping "github.com/stackrox/rox/central/policycategory/index/mappings"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/schema"

	processBaselineMapping "github.com/stackrox/rox/central/processbaseline/index/mappings"
	roleOptions "github.com/stackrox/rox/central/rbac/k8srole/mappings"
	roleBindingOptions "github.com/stackrox/rox/central/rbac/k8srolebinding/mappings"
	subjectMapping "github.com/stackrox/rox/central/rbac/service/mapping"
	reportConfigurationsMapping "github.com/stackrox/rox/central/reportconfigurations/mappings"
	riskMappings "github.com/stackrox/rox/central/risk/mappings"
	secretOptions "github.com/stackrox/rox/central/secret/mappings"
	serviceAccountOptions "github.com/stackrox/rox/central/serviceaccount/mappings"
	vulnReqMapping "github.com/stackrox/rox/central/vulnerabilityrequest/mappings"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/blevesearch"
	"github.com/stackrox/rox/pkg/search/options/deployments"
	imageMapping "github.com/stackrox/rox/pkg/search/options/images"
	"github.com/stackrox/rox/pkg/search/options/processindicators"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	indexMappingOnce sync.Once
	indexMapping     *mapping.IndexMappingImpl
)

// GetIndexMapping returns the current index mapping
func GetIndexMapping() mapping.IndexMapping {
	indexMappingOnce.Do(func() {
		indexMapping = bleve.NewIndexMapping()

		utils.Must(indexMapping.AddCustomAnalyzer("single_term", singleTermAnalyzer()))
		indexMapping.DefaultAnalyzer = "single_term" // Default to our analyzer

		indexMapping.IndexDynamic = false
		indexMapping.StoreDynamic = false
		indexMapping.TypeField = "Type"
		indexMapping.DefaultMapping = getDefaultDocMapping()

		for category, optMap := range GetEntityOptionsMap() {
			indexMapping.AddDocumentMapping(category.String(), blevesearch.DocumentMappingFromOptionsMap(optMap.Original()))
		}

		disabledSection := bleve.NewDocumentDisabledMapping()
		indexMapping.AddDocumentMapping("_all", disabledSection)
	})
	return indexMapping
}

func getDefaultDocMapping() *mapping.DocumentMapping {
	return &mapping.DocumentMapping{
		Enabled: false,
		Dynamic: false,
	}
}

// This is the custom analyzer definition
func singleTermAnalyzer() map[string]interface{} {
	return map[string]interface{}{
		"type":         custom.Name,
		"char_filters": []string{},
		"tokenizer":    whitespace.Name,
		// Ignore case sensitivity
		"token_filters": []string{
			lowercase.Name,
		},
	}
}

// GetEntityOptionsMap is a mapping from search categories to the options
func GetEntityOptionsMap() map[v1.SearchCategory]search.OptionsMap {
	var fullImageOptionsMap search.OptionsMap
	var fullNodeOptionsMap search.OptionsMap
	var fullClusterCVEOptionsMap search.OptionsMap

	if features.PostgresDatastore.Enabled() {
		fullNodeOptionsMap = search.CombineOptionsMaps(
			schema.NodesSchema.OptionsMap,
			schema.NodeComponentsSchema.OptionsMap,
			schema.NodeComponentEdgesSchema.OptionsMap,
			schema.NodeComponentsCvesEdgesSchema.OptionsMap,
			schema.NodeCvesSchema.OptionsMap,
		)

		fullImageOptionsMap = search.CombineOptionsMaps(
			schema.ImagesSchema.OptionsMap,
			imageMapping.ImageDeploymentOptions,
			schema.ImageComponentEdgesSchema.OptionsMap,
			schema.ImageComponentsSchema.OptionsMap,
			schema.ImageComponentCveEdgesSchema.OptionsMap,
			schema.ImageCvesSchema.OptionsMap,
			schema.ImageCveEdgesSchema.OptionsMap,
		)

		fullClusterCVEOptionsMap = search.CombineOptionsMaps(
			schema.ClustersSchema.OptionsMap,
			schema.ClusterCveEdgesSchema.OptionsMap,
			schema.ClusterCvesSchema.OptionsMap,
		)
	} else {
		fullNodeOptionsMap = search.CombineOptionsMaps(
			nodeMapping.OptionsMap,
			nodeComponentEdgeMapping.OptionsMap,
			imageComponentMapping.OptionsMap,
			componentVulnEdgeMapping.OptionsMap,
			cveMapping.OptionsMap,
		)

		fullImageOptionsMap = search.CombineOptionsMaps(
			imageMapping.OptionsMap,
			imageMapping.ImageDeploymentOptions,
			imageComponentEdgeMapping.OptionsMap,
			imageComponentMapping.OptionsMap,
			componentVulnEdgeMapping.OptionsMap,
			imageCVEEdgeMapping.OptionsMap,
			cveMapping.OptionsMap,
		)
	}

	// EntityOptionsMap is a mapping from search categories to the options map for that category.
	// search document maps are also built off this map
	entityOptionsMap := map[v1.SearchCategory]search.OptionsMap{
		v1.SearchCategory_ACTIVE_COMPONENT:      activeComponentMappings.OptionsMap,
		v1.SearchCategory_ALERTS:                alertMapping.OptionsMap,
		v1.SearchCategory_DEPLOYMENTS:           deployments.OptionsMap,
		v1.SearchCategory_PODS:                  podMapping.OptionsMap,
		v1.SearchCategory_IMAGES:                fullImageOptionsMap,
		v1.SearchCategory_POLICIES:              policyMapping.OptionsMap,
		v1.SearchCategory_POLICY_CATEGORIES:     policyCategoryMapping.OptionsMap,
		v1.SearchCategory_SECRETS:               secretOptions.OptionsMap,
		v1.SearchCategory_PROCESS_INDICATORS:    processindicators.OptionsMap,
		v1.SearchCategory_COMPLIANCE_STANDARD:   index.StandardOptions,
		v1.SearchCategory_COMPLIANCE_CONTROL:    index.ControlOptions,
		v1.SearchCategory_CLUSTERS:              clusterMapping.OptionsMap,
		v1.SearchCategory_NAMESPACES:            namespaceMapping.OptionsMap,
		v1.SearchCategory_NODES:                 fullNodeOptionsMap,
		v1.SearchCategory_PROCESS_BASELINES:     processBaselineMapping.OptionsMap,
		v1.SearchCategory_REPORT_CONFIGURATIONS: reportConfigurationsMapping.OptionsMap,
		v1.SearchCategory_RISKS:                 riskMappings.OptionsMap,
		v1.SearchCategory_ROLES:                 roleOptions.OptionsMap,
		v1.SearchCategory_ROLEBINDINGS:          roleBindingOptions.OptionsMap,
		v1.SearchCategory_SERVICE_ACCOUNTS:      serviceAccountOptions.OptionsMap,
		v1.SearchCategory_SUBJECTS:              subjectMapping.OptionsMap,
		v1.SearchCategory_IMAGE_COMPONENTS:      fullImageOptionsMap,
		v1.SearchCategory_VULN_REQUEST:          vulnReqMapping.OptionsMap,
		v1.SearchCategory_IMAGE_INTEGRATIONS:    imageIntegrationMapping.OptionsMap,
	}
	if features.PostgresDatastore.Enabled() {
		entityOptionsMap[v1.SearchCategory_NODE_COMPONENTS] = fullImageOptionsMap

		entityOptionsMap[v1.SearchCategory_NODE_VULNERABILITIES] = fullImageOptionsMap
		entityOptionsMap[v1.SearchCategory_IMAGE_VULNERABILITIES] = fullImageOptionsMap
		entityOptionsMap[v1.SearchCategory_CLUSTER_VULNERABILITIES] = fullClusterCVEOptionsMap
	} else {
		entityOptionsMap[v1.SearchCategory_VULNERABILITIES] = search.CombineOptionsMaps(fullImageOptionsMap, fullNodeOptionsMap)
	}

	return entityOptionsMap
}
