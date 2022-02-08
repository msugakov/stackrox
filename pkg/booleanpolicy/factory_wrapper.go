package booleanpolicy

import (
	"github.com/stackrox/rox/pkg/booleanpolicy/evaluator"
	"github.com/stackrox/rox/pkg/booleanpolicy/evaluator/pathutil"
	"github.com/stackrox/rox/pkg/booleanpolicy/query"
	"github.com/stackrox/rox/pkg/booleanpolicy/regocompile"
	"github.com/stackrox/rox/pkg/features"
)

type factoryWrapper struct {
	legacyFactory   evaluator.Factory
	opaBasedFactory regocompile.RegoCompilerForType
}

func (f *factoryWrapper) GenerateEvaluator(q *query.Query) (evaluator.Evaluator, error) {
	if features.OPABasedEvaluator.Enabled() {
		eval, err := f.opaBasedFactory.CompileRegoBasedEvaluator(q)
		if err == nil {
			return eval, nil
		}
		if err != regocompile.ErrRegoNotYetSupported {
			return nil, err
		}
	}
	return f.legacyFactory.GenerateEvaluator(q)
}

// MustCreateFactoryWrapper returns a factory wrapper.
// A factory wrapper routes between the OPA and the legacy factory
// depending on the value of the feature flag.
// This is temporary code until the OPA feature flag is removed.
func MustCreateFactoryWrapper(objMeta *pathutil.AugmentedObjMeta) evaluator.Factory {
	return &factoryWrapper{
		legacyFactory:   evaluator.MustCreateNewFactory(objMeta),
		opaBasedFactory: regocompile.MustCreateRegoCompiler(objMeta),
	}
}
