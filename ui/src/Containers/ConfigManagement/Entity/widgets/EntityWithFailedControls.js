import React from 'react';
import PropTypes from 'prop-types';
import entityTypes from 'constants/entityTypes';
import { entityAcrossControlsColumns } from 'constants/listColumns';

import TableWidget from './TableWidget';

export const getRelatedEntities = (data, entityType) => {
    const { results } = data;
    if (!results.length) return [];
    const relatedEntities = {};
    let entityKey = 0;
    results[0].aggregationKeys.forEach(({ scope }, idx) => {
        if (scope === entityTypes[entityType]) entityKey = idx;
    });
    results.forEach(({ keys, numFailing }) => {
        const { id, name, clusterName } = keys[entityKey];
        if (!relatedEntities[id]) {
            relatedEntities[id] = {
                id,
                name,
                clusterName
            };
        } else if (numFailing) relatedEntities[id].passing = false;
    });
    return Object.values(relatedEntities);
};

const EntityWithFailedControls = ({ entityType, entities }) => {
    const relatedEntities = getRelatedEntities(entities, entityType);
    const failingRelatedEntities = relatedEntities.filter(relatedEntity => !relatedEntity.passing);
    const tableHeader = `${failingRelatedEntities.length} nodes have failed across this control`;
    return (
        <TableWidget
            entityType={entityType}
            header={tableHeader}
            rows={failingRelatedEntities}
            noDataText="No Nodes"
            className="bg-base-100 w-full"
            columns={entityAcrossControlsColumns[entityType]}
            idAttribute="id"
        />
    );
};

EntityWithFailedControls.propTypes = {
    entityType: PropTypes.string.isRequired,
    entities: PropTypes.shape({
        results: PropTypes.arrayOf(PropTypes.shape())
    })
};

EntityWithFailedControls.defaultProps = {
    entities: {
        results: []
    }
};

export default EntityWithFailedControls;
