import React, { useEffect } from 'react';
import { connect, useSelector } from 'react-redux';
import { createStructuredSelector } from 'reselect';

import { selectors } from 'reducers';
import useLocalStorage from 'hooks/useLocalStorage';
import { actions as dialogueActions } from 'reducers/network/dialogue';
import { actions as sidepanelActions } from 'reducers/network/sidepanel';
import { actions as pageActions } from 'reducers/network/page';
import dialogueStages from 'Containers/Network/Dialogue/dialogueStages';
import useNetworkPolicySimulation from 'Containers/Network/useNetworkPolicySimulation';
import useNetworkBaselineSimulation from 'Containers/Network/useNetworkBaselineSimulation';
import useFetchBaselineComparisons from 'Containers/Network/useFetchBaselineComparisons';
import Dialogue from 'Containers/Network/Dialogue';
import Graph from 'Containers/Network/Graph/Graph';
import SidePanel from 'Containers/Network/SidePanel/SidePanel';
import SimulationFrame from 'Components/SimulationFrame';
import { getErrorMessageFromServerResponse } from 'utils/networkGraphUtils';
import Header from './Header/Header';
import NoSelectedNamespace from './NoSelectedNamespace';
import GraphLoadErrorState from './GraphLoadErrorState';

function GraphFrame() {
    const [showNamespaceFlows, setShowNamespaceFlows] = useLocalStorage(
        'showNamespaceFlows',
        'show'
    );
    const { isNetworkSimulationOn, isNetworkSimulationError, stopNetworkSimulation } =
        useNetworkPolicySimulation();
    const { isBaselineSimulationOn, stopBaselineSimulation } = useNetworkBaselineSimulation();
    const { simulatedBaselines } = useFetchBaselineComparisons();

    const isSimulationOn = isNetworkSimulationOn || isBaselineSimulationOn;
    let onStop;
    if (isNetworkSimulationOn) {
        onStop = stopNetworkSimulation;
    }
    if (isBaselineSimulationOn) {
        onStop = stopBaselineSimulation;
    }
    const isError = isNetworkSimulationOn && isNetworkSimulationError;

    function handleNamespaceFlowsToggle(mode) {
        setShowNamespaceFlows(mode);
    }

    return isSimulationOn ? (
        <SimulationFrame isError={isError} onStop={onStop}>
            <div className="flex flex-1 relative">
                <Graph
                    isSimulationOn
                    showNamespaceFlows={showNamespaceFlows}
                    setShowNamespaceFlows={handleNamespaceFlowsToggle}
                    simulatedBaselines={simulatedBaselines}
                />
                <SidePanel />
            </div>
        </SimulationFrame>
    ) : (
        <div className="flex flex-1 relative">
            <Graph
                showNamespaceFlows={showNamespaceFlows}
                setShowNamespaceFlows={handleNamespaceFlowsToggle}
            />
            <SidePanel />
        </div>
    );
}

const networkPageContentSelector = createStructuredSelector({
    clusters: selectors.getClusters,
    selectedClusterId: selectors.getSelectedNetworkClusterId,
    selectedNamespaceFilters: selectors.getSelectedNamespaceFilters,
});

function NetworkPage({
    getNetworkFlowGraphState,
    serverErrorMessage,
    closeSidePanel,
    setDialogueStage,
    setNetworkModification,
}) {
    const { isNetworkSimulationOn } = useNetworkPolicySimulation();
    const { isBaselineSimulationOn } = useNetworkBaselineSimulation();
    const isSimulationOn = isNetworkSimulationOn || isBaselineSimulationOn;

    const { clusters, selectedClusterId, selectedNamespaceFilters } = useSelector(
        networkPageContentSelector
    );

    const clusterName = clusters.find((c) => c.id === selectedClusterId)?.name;

    // when this component unmounts, then close the side panel and exit network policy simulation
    useEffect(() => {
        return () => {
            closeSidePanel();
            setDialogueStage(dialogueStages.closed);
            setNetworkModification(null);
        };
    }, [closeSidePanel, setDialogueStage, setNetworkModification]);

    const hasNoSelectedNamespace = selectedNamespaceFilters.length === 0;
    const hasGraphLoadError = getNetworkFlowGraphState === 'ERROR';

    let content;
    if (hasNoSelectedNamespace) {
        content = <NoSelectedNamespace clusterName={clusterName} />;
    } else if (hasGraphLoadError) {
        const userMessage = getErrorMessageFromServerResponse(serverErrorMessage);
        content = <GraphLoadErrorState error={serverErrorMessage} userMessage={userMessage} />;
    } else {
        content = <GraphFrame />;
    }

    return (
        <>
            <Header
                isGraphDisabled={hasNoSelectedNamespace || hasGraphLoadError}
                isSimulationOn={isSimulationOn}
            />
            <section className="flex flex-1 h-full w-full">
                <div className="flex flex-1 flex-col w-full overflow-hidden">
                    <div className="flex flex-1 flex-col relative">{content}</div>
                </div>
                <Dialogue />
            </section>
        </>
    );
}

const mapStateToProps = createStructuredSelector({
    getNetworkFlowGraphState: selectors.getNetworkFlowGraphState,
    serverErrorMessage: selectors.getNetworkFlowErrorMessage,
});

const mapDispatchToProps = {
    closeSidePanel: pageActions.closeSidePanel,
    setNetworkModification: sidepanelActions.setNetworkPolicyModification,
    setDialogueStage: dialogueActions.setNetworkDialogueStage,
};

export default connect(mapStateToProps, mapDispatchToProps)(NetworkPage);
