/* eslint-disable react/jsx-no-bind */
import React, { useEffect, useState, useReducer } from 'react';
import PropTypes from 'prop-types';
import ReactRouterPropTypes from 'react-router-prop-types';
import * as Icon from 'react-feather';
import { generatePath } from 'react-router-dom';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';
import get from 'lodash/get';

import CloseButton from 'Components/CloseButton';
import Dialog from 'Components/Dialog';
import PageHeader from 'Components/PageHeader';
import Panel from 'Components/Panel';
import PanelButton from 'Components/PanelButton';
import SearchInput from 'Components/SearchInput';
import StatusField from 'Components/StatusField';
import ToggleSwitch from 'Components/ToggleSwitch';
import CheckboxTable from 'Components/CheckboxTable';
import {
    defaultHeaderClassName,
    defaultColumnClassName,
    wrapClassName,
    rtTrActionsClassName,
} from 'Components/Table';
import TableHeader from 'Components/TableHeader';
import RowActionButton from 'Components/RowActionButton';
import useInterval from 'hooks/useInterval';
import { actions as clustersActions } from 'reducers/clusters';
import { selectors } from 'reducers';
import {
    fetchClustersAsArray,
    getAutoUpgradeConfig,
    deleteClusters,
    upgradeCluster,
    upgradeClusters,
    saveAutoUpgradeConfig,
} from 'services/ClustersService';
import { toggleRow, toggleSelectAll } from 'utils/checkboxUtils';
import { clustersPath } from 'routePaths';
import { sortVersion } from 'sorters/sorters';

import ClustersSidePanel from './ClustersSidePanel';

// @TODO, refactor these helper utilities to this folder,
//        when retiring clusters in Integrations section
import {
    clusterTablePollingInterval,
    formatClusterType,
    formatCollectionMethod,
    formatConfiguredField,
    formatLastCheckIn,
    formatSensorVersion,
    parseUpgradeStatus,
    getUpgradeableClusters,
    getCredentialExpirationProps,
} from './cluster.helpers';
import CredentialExpiration from './CredentialExpiration';

const ClustersPage = ({
    history,
    location: { search },
    match: {
        params: { clusterId },
    },
    searchOptions,
    searchModifiers,
    searchSuggestions,
    setSearchModifiers,
    setSearchOptions,
    setSearchSuggestions,
}) => {
    const [currentClusters, setCurrentClusters] = useState([]);
    const [showDialog, setShowDialog] = useState(false);
    const [autoUpgradeConfig, setAutoUpgradeConfig] = useState({});
    const [checkedClusterIds, setCheckedClusters] = useState([]);
    const [upgradableClusters, setUpgradableClusters] = useState([]);
    const [tableRef, setTableRef] = useState(null);
    const [selectedClusterId, setSelectedClusterId] = useState(clusterId);
    const [pollingCount, setPollingCount] = useState(0);

    function notificationsReducer(state, action) {
        switch (action.type) {
            case 'ADD_NOTIFICATION': {
                return [...state, action.payload];
            }
            case 'REMOVE_NOTIFICATION': {
                return state.filter((note) => note !== action.payload);
            }
            default: {
                return state;
            }
        }
    }
    const [notifications, dispatch] = useReducer(notificationsReducer, []);

    // @TODO, implement actual delete logic into this stub function
    const onDeleteHandler = (cluster) => (e) => {
        e.stopPropagation();
        setCheckedClusters([cluster.id]);
        setShowDialog(true);
    };

    function renderRowActionButtons(cluster) {
        return (
            <div className="border-2 border-r-2 border-base-400 bg-base-100">
                <RowActionButton
                    text="Delete cluster"
                    icon={<Icon.Trash2 className="my-1 h-4 w-4" />}
                    className="hover:bg-alert-200 text-alert-600 hover:text-alert-700"
                    onClick={onDeleteHandler(cluster)}
                />
            </div>
        );
    }

    function calculateUpgradeableClusters(selection) {
        const currentlySelectedClusters = currentClusters.filter((cluster) =>
            selection.includes(cluster.id)
        );

        const upgradeableList = getUpgradeableClusters(currentlySelectedClusters);

        setUpgradableClusters(upgradeableList);
    }

    function toggleCluster(id) {
        const selection = toggleRow(id, checkedClusterIds);
        setCheckedClusters(selection);

        calculateUpgradeableClusters(selection);
    }

    function toggleAllClusters() {
        const rowsLength = checkedClusterIds.length;
        const ref = tableRef.reactTable;
        const selection = toggleSelectAll(rowsLength, checkedClusterIds, ref);
        setCheckedClusters(selection);

        calculateUpgradeableClusters(selection);
    }

    // @TODO: Change table component to use href for accessibility and better UX here, instead of an onclick
    function handleRowClick(cluster) {
        const newClusterId = (cluster && cluster.id) || '';
        setSelectedClusterId(newClusterId);
    }

    function refreshClusterList() {
        return fetchClustersAsArray(searchOptions).then((clusters) => {
            setCurrentClusters(clusters);
        });
    }

    useEffect(
        () => {
            refreshClusterList();
        },
        // eslint-disable-next-line react-hooks/exhaustive-deps
        [searchOptions, pollingCount]
    );

    // use a custom hook to set up polling, thanks Dan Abramov and Rob Stark
    useInterval(() => {
        setPollingCount(pollingCount + 1);
    }, clusterTablePollingInterval);

    function fetchConfig() {
        getAutoUpgradeConfig().then((config) => {
            setAutoUpgradeConfig(config);
        });
    }

    function onAddCluster() {
        setSelectedClusterId('new');
    }

    function upgradeSelectedClusters() {
        upgradeClusters(checkedClusterIds).then(() => {
            setCheckedClusters([]);

            refreshClusterList();
        });
    }

    function upgradeSingleCluster(id) {
        upgradeCluster(id)
            .then(() => {
                refreshClusterList();
            })
            .catch((error) => {
                const serverError = get(
                    error,
                    'response.data.message',
                    'An unknown error has occurred.'
                );
                const givenCluster = currentClusters.find((cluster) => cluster.id === id);
                const clusterName = givenCluster ? givenCluster.name : '-';
                const payload = `Failed to trigger upgrade for cluster ${clusterName}. Error: ${serverError}`;

                dispatch({ type: 'ADD_NOTIFICATION', payload });
            });
    }

    function deleteSelectedClusters() {
        setShowDialog(true);
    }

    function hideDialog() {
        setShowDialog(false);
    }

    function makeDeleteRequest() {
        deleteClusters(checkedClusterIds)
            .then(() => {
                setCheckedClusters([]);

                fetchClustersAsArray().then((clusters) => {
                    setCurrentClusters(clusters);
                });
            })
            .finally(() => {
                setShowDialog(false);
            });
    }

    useEffect(() => {
        fetchConfig();
    }, []);

    // When the selected cluster changes, update the URL.
    useEffect(() => {
        const newPath = selectedClusterId
            ? generatePath(clustersPath, { clusterId: selectedClusterId })
            : clustersPath.replace('/:clusterId?', '');
        history.push({
            pathname: newPath,
            search,
        });
    }, [history, search, selectedClusterId]);

    function toggleAutoUpgrade() {
        // @TODO, wrap this settings change in a confirmation prompt of some sort
        const previousValue = autoUpgradeConfig.enableAutoUpgrade;
        const newConfig = {
            ...autoUpgradeConfig,
            enableAutoUpgrade: !previousValue,
        };

        setAutoUpgradeConfig(newConfig); // optimistically set value before API call

        saveAutoUpgradeConfig(newConfig).catch(() => {
            // reverse the optimistic update of the control in the UI
            const rollbackConfig = {
                ...autoUpgradeConfig,
                enableAutoUpgrade: previousValue,
            };
            setAutoUpgradeConfig(rollbackConfig);

            // also, re-fetch the data from the server, just in case it did update but we didn't get the network response
            fetchConfig();
        });
    }

    function getUpgradeStatusField(original) {
        const status = parseUpgradeStatus(get(original, 'status.upgradeStatus', null));
        if (!status) {
            return '-';
        }

        if (status.action) {
            status.action.actionHandler = (e) => {
                e.stopPropagation();
                upgradeSingleCluster(original.id);
            };
        }

        return (
            <StatusField
                displayValue={status.displayValue}
                type={status.type}
                action={status.action}
            />
        );
    }

    const headerActions = (
        <>
            <PanelButton
                icon={<Icon.DownloadCloud className="h-4 w-4 ml-1" />}
                tooltip={`Upgrade (${upgradableClusters.length})`}
                className="btn btn-tertiary ml-2"
                onClick={upgradeSelectedClusters}
                disabled={upgradableClusters.length === 0 || !!selectedClusterId}
            >
                {`Upgrade (${upgradableClusters.length})`}
            </PanelButton>
            <PanelButton
                icon={<Icon.Trash2 className="h-4 w-4 ml-1" />}
                tooltip={`Delete (${checkedClusterIds.length})`}
                className="btn btn-alert ml-2"
                onClick={deleteSelectedClusters}
                disabled={checkedClusterIds.length === 0 || !!selectedClusterId}
            >
                {`Delete (${checkedClusterIds.length})`}
            </PanelButton>
            <PanelButton
                icon={<Icon.Plus className="h-4 w-4 ml-1" />}
                tooltip="New Cluster"
                className="btn btn-base ml-2"
                onClick={onAddCluster}
                disabled={!!selectedClusterId}
            >
                New Cluster
            </PanelButton>
        </>
    );

    const headerComponent = (
        <TableHeader length={currentClusters.length} type="Cluster" isViewFiltered={false} />
    );

    // Because of fixed checkbox width, total of column ratios must be less than 100%
    // 1 * 1/9 + 7 * 1/8 = 98.6%
    const clusterColumns = [
        {
            accessor: 'name',
            Header: 'Name',
            headerClassName: `w-1/8 ${defaultHeaderClassName}`,
            className: `w-1/8 ${wrapClassName} ${defaultColumnClassName}`,
        },
        {
            Header: 'Orchestrator',
            Cell: ({ original }) => formatClusterType(original.type),
            headerClassName: `w-1/9 ${defaultHeaderClassName}`,
            className: `w-1/9 ${wrapClassName} ${defaultColumnClassName}`,
        },
        {
            Header: 'Runtime collection',
            Cell: ({ original }) => formatCollectionMethod(original.collectionMethod),
            headerClassName: `w-1/8 ${defaultHeaderClassName}`,
            className: `w-1/8 ${wrapClassName} ${defaultColumnClassName}`,
        },
        {
            Header: 'Admission Controller',
            Cell: ({ original }) => formatConfiguredField(original.admissionController),
            headerClassName: `w-1/8 ${defaultHeaderClassName}`,
            className: `w-1/8 ${wrapClassName} ${defaultColumnClassName}`,
        },
        {
            Header: 'Last check-in',
            Cell: ({ original }) => formatLastCheckIn(original.status),
            headerClassName: `w-1/8 ${defaultHeaderClassName}`,
            className: `w-1/8 ${wrapClassName} ${defaultColumnClassName}`,
        },
        {
            Header: 'Upgrade status',
            Cell: ({ original }) => getUpgradeStatusField(original),
            headerClassName: `w-1/8 ${defaultHeaderClassName}`,
            className: `w-1/8 ${wrapClassName} ${defaultColumnClassName}`,
        },
        {
            Header: 'Current Sensor version',
            Cell: ({ original }) => formatSensorVersion(original.status),
            headerClassName: `w-1/8 ${defaultHeaderClassName}`,
            className: `w-1/8 ${wrapClassName} ${defaultColumnClassName} word-break`,
            sortMethod: sortVersion,
        },
        {
            Header: 'Credential Expiration',
            Cell: ({ original }) => {
                const props = getCredentialExpirationProps(original.status?.certExpiryStatus);
                if (!props) {
                    return 'N/A';
                }
                return (
                    <CredentialExpiration
                        expiration={props.sensorCertExpiry}
                        showExpiringSoon={props.showExpiringSoon}
                        type={props.messageType}
                        diffInWords={props.diffInWords}
                    />
                );
            },
            headerClassName: `w-1/8 ${defaultHeaderClassName}`,
            className: `w-1/8 ${wrapClassName} ${defaultColumnClassName} word-break`,
            sortMethod: sortVersion,
        },
        {
            Header: '',
            accessor: '',
            headerClassName: 'hidden',
            className: rtTrActionsClassName,
            Cell: ({ original }) => renderRowActionButtons(original),
        },
    ];

    const headerText = 'Clusters';
    const subHeaderText = 'Resource list';
    const defaultOption = searchModifiers.find((x) => x.value === 'Cluster:');

    const pageHeader = (
        <PageHeader header={headerText} subHeader={subHeaderText}>
            <div className="flex flex-1 items-center justify-end">
                <SearchInput
                    className="w-full"
                    id="clusters-search"
                    searchOptions={searchOptions}
                    searchModifiers={searchModifiers}
                    searchSuggestions={searchSuggestions}
                    setSearchOptions={setSearchOptions}
                    setSearchModifiers={setSearchModifiers}
                    setSearchSuggestions={setSearchSuggestions}
                    defaultOption={defaultOption}
                    autoCompleteCategories={['CLUSTERS']}
                />
                <div className="flex items-center min-w-64 ml-4">
                    <ToggleSwitch
                        id="enableAutoUpgrade"
                        toggleHandler={toggleAutoUpgrade}
                        label="Automatically upgrade secured clusters"
                        enabled={autoUpgradeConfig.enableAutoUpgrade}
                    />
                </div>
            </div>
        </PageHeader>
    );

    const messages = notifications.map((note) => (
        <div
            key={note}
            className="flex flex-1 border-b border-base-400 items-center justify-end relative py-0 pl-3 w-full"
        >
            <span className="w-full">{note}</span>
            <CloseButton
                onClose={() => {
                    dispatch({ type: 'REMOVE_NOTIFICATION', payload: note });
                }}
                className="border-base-400 border-l"
            />
        </div>
    ));

    // getValueFromSelectedClusterAtPath returns the value at the given path
    // within the selected cluster, if there is a selected cluster.
    // It returns null if there is no selected cluster, or if the selected cluster
    // has no value at path.
    function getValueFromSelectedClusterAtPath(path) {
        if (!currentClusters || !currentClusters.length) {
            return null;
        }
        const selectedCluster = currentClusters.find((cluster) => cluster.id === selectedClusterId);
        if (!selectedCluster) {
            return null;
        }
        return get(selectedCluster, path, null);
    }

    function getCurrentCertExpiryStatusOrNull() {
        return getValueFromSelectedClusterAtPath('status.certExpiryStatus');
    }

    function getCurrentUpgradeStatusOrNull() {
        return getValueFromSelectedClusterAtPath('status.upgradeStatus');
    }

    return (
        <section className="flex flex-1 flex-col h-full">
            <div className="flex flex-1 flex-col">
                {pageHeader}
                <div className="flex flex-1 relative">
                    <div className="shadow border-primary-300 w-full overflow-hidden">
                        <Panel
                            headerTextComponent={headerComponent}
                            headerComponents={headerActions}
                        >
                            {messages.length > 0 && (
                                <div className="flex flex-col w-full items-center bg-warning-200 text-warning-8000 justify-center font-700 text-center">
                                    {messages}
                                </div>
                            )}
                            <div className="w-full">
                                <CheckboxTable
                                    ref={(table) => {
                                        setTableRef(table);
                                    }}
                                    rows={currentClusters}
                                    columns={clusterColumns}
                                    onRowClick={handleRowClick}
                                    toggleRow={toggleCluster}
                                    toggleSelectAll={toggleAllClusters}
                                    selection={checkedClusterIds}
                                    selectedRowId={selectedClusterId}
                                    noDataText="No clusters to show."
                                    minRows={20}
                                />
                            </div>
                        </Panel>
                    </div>
                    <ClustersSidePanel
                        selectedClusterId={selectedClusterId}
                        setSelectedClusterId={setSelectedClusterId}
                        upgradeStatus={getCurrentUpgradeStatusOrNull()}
                        certExpiryStatus={getCurrentCertExpiryStatusOrNull()}
                    />
                </div>
            </div>
            <Dialog
                className="w-1/3"
                isOpen={showDialog}
                text={`Cluster deletion won't tear down StackRox services running on this cluster. You can remove them from the corresponding cluster by running the "delete-sensor.sh" script from the sensor installation bundle. Are you sure you want to delete ${checkedClusterIds.length} cluster(s)?`}
                onConfirm={makeDeleteRequest}
                confirmText="Delete"
                onCancel={hideDialog}
                isDestructive
            />
        </section>
    );
};

ClustersPage.propTypes = {
    history: ReactRouterPropTypes.history.isRequired,
    location: ReactRouterPropTypes.location.isRequired,
    match: ReactRouterPropTypes.match.isRequired,

    // Search specific input.
    searchOptions: PropTypes.arrayOf(PropTypes.object).isRequired,
    searchModifiers: PropTypes.arrayOf(PropTypes.object).isRequired,
    searchSuggestions: PropTypes.arrayOf(PropTypes.object).isRequired,
    setSearchOptions: PropTypes.func.isRequired,
    setSearchModifiers: PropTypes.func.isRequired,
    setSearchSuggestions: PropTypes.func.isRequired,
};

const mapStateToProps = createStructuredSelector({
    searchOptions: selectors.getClustersSearchOptions,
    searchModifiers: selectors.getClustersSearchModifiers,
    searchSuggestions: selectors.getClustersSearchSuggestions,
});

const mapDispatchToProps = {
    setSearchOptions: clustersActions.setClustersSearchOptions,
    setSearchModifiers: clustersActions.setClustersSearchModifiers,
    setSearchSuggestions: clustersActions.setClustersSearchSuggestions,
};

export default connect(mapStateToProps, mapDispatchToProps)(ClustersPage);
