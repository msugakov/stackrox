package getters

import (
	"context"

	v1 "github.com/stackrox/stackrox/generated/api/v1"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/search"
)

// MockAlertsGetter is a mock AlertsGetter.
type MockAlertsGetter struct {
	Alerts []*storage.ListAlert
}

// ListAlerts supports a limited set of request parameters.
// It only needs to be as specific as the production code.
func (m MockAlertsGetter) ListAlerts(ctx context.Context, req *v1.ListAlertsRequest) (alerts []*storage.ListAlert, err error) {
	q, err := search.ParseQuery(req.GetQuery())
	if err != nil {
		return nil, err
	}

	state := storage.ViolationState_ACTIVE.String()
	search.ApplyFnToAllBaseQueries(q, func(bq *v1.BaseQuery) {
		mfQ, ok := bq.GetQuery().(*v1.BaseQuery_MatchFieldQuery)
		if ok && mfQ.MatchFieldQuery.GetField() == search.ViolationState.String() {
			state = mfQ.MatchFieldQuery.GetValue()
		}
	})

	for _, a := range m.Alerts {
		if a.GetState().String() == state {
			alerts = append(alerts, a)
		}
	}
	return
}
