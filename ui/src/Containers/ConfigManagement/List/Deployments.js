import React, { useContext } from 'react';
import pluralize from 'pluralize';

import {
    defaultHeaderClassName,
    defaultColumnClassName,
    nonSortableHeaderClassName,
} from 'Components/Table';
import StatusChip from 'Components/StatusChip';
import searchContext from 'Containers/searchContext';
import entityTypes from 'constants/entityTypes';
import { deploymentSortFields } from 'constants/sortFields';
import { entityListPropTypes, entityListDefaultprops } from 'constants/entityPageProps';
import { CLIENT_SIDE_SEARCH_OPTIONS as SEARCH_OPTIONS } from 'constants/searchOptions';
import { DEPLOYMENTS_QUERY } from 'queries/deployment';
import queryService from 'utils/queryService';
import URLService from 'utils/URLService';
import List from './List';
import TableCellLink from './Link';
import filterByPolicyStatus from './utilities/filterByPolicyStatus';

export const defaultDeploymentSort = [
    {
        id: deploymentSortFields.DEPLOYMENT,
        desc: false,
    },
];

const buildTableColumns = (match, location, entityContext) => {
    const tableColumns = [
        {
            Header: 'Id',
            headerClassName: 'hidden',
            className: 'hidden',
            accessor: 'id',
        },
        {
            Header: `Deployment`,
            headerClassName: `w-1/8 ${defaultHeaderClassName}`,
            className: `w-1/8 ${defaultColumnClassName}`,
            accessor: 'name',
            id: deploymentSortFields.DEPLOYMENT,
            sortField: deploymentSortFields.DEPLOYMENT,
        },
        entityContext && entityContext[entityTypes.CLUSTER]
            ? null
            : {
                  Header: `Cluster`,
                  headerClassName: `w-1/8 ${defaultHeaderClassName}`,
                  className: `w-1/8 ${defaultColumnClassName}`,
                  accessor: 'clusterName',
                  // eslint-disable-next-line
                  Cell: ({ original, pdf }) => {
                      const { clusterName, clusterId, id } = original;
                      const url = URLService.getURL(match, location)
                          .push(id)
                          .push(entityTypes.CLUSTER, clusterId)
                          .url();
                      return <TableCellLink pdf={pdf} url={url} text={clusterName} />;
                  },
                  id: deploymentSortFields.CLUSTER,
                  sortField: deploymentSortFields.CLUSTER,
              },
        entityContext && entityContext[entityTypes.NAMESPACE]
            ? null
            : {
                  Header: `Namespace`,
                  headerClassName: `w-1/8 ${defaultHeaderClassName}`,
                  className: `w-1/8 ${defaultColumnClassName}`,
                  accessor: 'namespace',
                  // eslint-disable-next-line
                  Cell: ({ original, pdf }) => {
                      const { namespace, namespaceId, id } = original;
                      const url = URLService.getURL(match, location)
                          .push(id)
                          .push(entityTypes.NAMESPACE, namespaceId)
                          .url();
                      return <TableCellLink pdf={pdf} url={url} text={namespace} />;
                  },
                  id: deploymentSortFields.NAMESPACE,
                  sortField: deploymentSortFields.NAMESPACE,
              },
        {
            Header: `Policy Status`,
            headerClassName: `w-1/8 ${nonSortableHeaderClassName}`,
            className: `w-1/8 ${defaultColumnClassName}`,
            // eslint-disable-next-line
            Cell: ({ original, pdf }) => {
                const { policyStatus } = original;
                return <StatusChip status={policyStatus} asString={pdf} />;
            },
            id: 'policyStatus',
            accessor: 'policyStatus',
            sortable: false,
        },
        {
            Header: `Images`,
            headerClassName: `w-1/8 ${nonSortableHeaderClassName}`,
            className: `w-1/8 ${defaultColumnClassName}`,
            // eslint-disable-next-line
            Cell: ({ original, pdf }) => {
                const { imageCount, id } = original;
                if (imageCount === 0) return 'No images';
                const url = URLService.getURL(match, location)
                    .push(id)
                    .push(entityTypes.IMAGE)
                    .url();
                return (
                    <TableCellLink
                        pdf={pdf}
                        url={url}
                        text={`${imageCount} ${pluralize('image', imageCount)}`}
                    />
                );
            },
            accessor: 'imageCount',
            sortable: false,
        },
        {
            Header: `Secrets`,
            headerClassName: `w-1/8 ${nonSortableHeaderClassName}`,
            className: `w-1/8 ${defaultColumnClassName}`,
            // eslint-disable-next-line
            Cell: ({ original, pdf }) => {
                const { secretCount, id } = original;
                if (secretCount === 0) return 'No secrets';
                const url = URLService.getURL(match, location)
                    .push(id)
                    .push(entityTypes.SECRET)
                    .url();
                return (
                    <TableCellLink
                        pdf={pdf}
                        url={url}
                        text={`${secretCount} ${pluralize('secret', secretCount)}`}
                    />
                );
            },
            accessor: 'secretCount',
            sortable: false,
        },
        entityContext && entityContext[entityTypes.SERVICE_ACCOUNT]
            ? null
            : {
                  Header: `Service Account`,
                  headerClassName: `w-1/8 ${defaultHeaderClassName}`,
                  className: `w-1/8 ${defaultColumnClassName}`,
                  accessor: 'serviceAccount',
                  // eslint-disable-next-line
                  Cell: ({ original, pdf }) => {
                      const { serviceAccount, serviceAccountID, id } = original;
                      const url = URLService.getURL(match, location)
                          .push(id)
                          .push(entityTypes.SERVICE_ACCOUNT, serviceAccountID)
                          .url();
                      return <TableCellLink pdf={pdf} url={url} text={serviceAccount} />;
                  },
                  id: deploymentSortFields.SERVICE_ACCOUNT,
                  sortField: deploymentSortFields.SERVICE_ACCOUNT,
              },
    ];
    return tableColumns.filter((col) => col);
};

const createTableRows = (data) => data.results;

const Deployments = ({
    match,
    location,
    className,
    selectedRowId,
    onRowClick,
    query,
    data,
    totalResults,
    entityContext,
}) => {
    const searchParam = useContext(searchContext);

    const autoFocusSearchInput = !selectedRowId;

    const tableColumns = buildTableColumns(match, location, entityContext);
    const {
        [SEARCH_OPTIONS.POLICY_STATUS.CATEGORY]: policyStatus,
        ...restQuery
    } = queryService.getQueryBasedOnSearchContext(query, searchParam);
    const queryText = queryService.objectToWhereClause({ ...restQuery });
    const variables = queryText ? { query: queryText } : null;

    function createTableRowsFilteredByPolicyStatus(items) {
        const tableRows = createTableRows(items);
        const filteredTableRows = filterByPolicyStatus(tableRows, policyStatus);
        return filteredTableRows;
    }

    return (
        <List
            className={className}
            query={DEPLOYMENTS_QUERY}
            variables={variables}
            entityType={entityTypes.DEPLOYMENT}
            tableColumns={tableColumns}
            createTableRows={createTableRowsFilteredByPolicyStatus}
            onRowClick={onRowClick}
            selectedRowId={selectedRowId}
            idAttribute="id"
            defaultSorted={defaultDeploymentSort}
            defaultSearchOptions={[SEARCH_OPTIONS.POLICY_STATUS.CATEGORY]}
            data={filterByPolicyStatus(data, policyStatus)}
            totalResults={totalResults}
            autoFocusSearchInput={autoFocusSearchInput}
        />
    );
};
Deployments.propTypes = entityListPropTypes;
Deployments.defaultProps = entityListDefaultprops;

export default Deployments;
