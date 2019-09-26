import React, { useContext } from 'react';
import gql from 'graphql-tag';
import entityTypes from 'constants/entityTypes';
import queryService from 'modules/queryService';
import Query from 'Components/ThrowingQuery';
import Loader from 'Components/Loader';
import PageNotFound from 'Components/PageNotFound';
import CollapsibleSection from 'Components/CollapsibleSection';
import ControlDetails from 'Components/ControlDetails';
import RelatedEntityListCount from 'Containers/ConfigManagement/Entity/widgets/RelatedEntityListCount';
import isGQLLoading from 'utils/gqlLoading';
import Widget from 'Components/Widget';
import searchContext from 'Containers/searchContext';
import { entityComponentPropTypes, entityComponentDefaultProps } from 'constants/entityPageProps';
import NodesWithFailedControls from './widgets/NodesWithFailedControls';
import Nodes from '../List/Nodes';

const QUERY = gql`
    query controlById($id: ID!, $where: String) {
        results: complianceControl(id: $id) {
            interpretationText
            description
            id
            name
            standardId
        }

        entities: complianceControl(id: $id) {
            complianceControlNodes {
                name
                clusterName
                id
                clusterId
                osImage
                containerRuntimeVersion
                joinedAt
                nodeComplianceControlCount(query: $where) {
                    failingCount
                    passingCount
                    unknownCount
                }
            }
        }
    }
`;

const Control = ({ id, entityListType, query, match, location, entityContext }) => {
    const searchParam = useContext(searchContext);

    const variables = {
        id,
        where: queryService.objectToWhereClause({
            ...query[searchParam],
            'Control Id': id,
            cacheBuster: new Date().getUTCMilliseconds()
        })
    };

    return (
        <Query query={QUERY} variables={variables} fetchPolicy="no-cache">
            {({ loading, data }) => {
                if (isGQLLoading(loading, data)) return <Loader transparent />;

                if (!data || !data.results)
                    return <PageNotFound resourceType={entityTypes.CONTROL} />;
                const { results: entity, entities = {} } = data;
                const { complianceControlNodes } = entities;

                if (entityListType) {
                    return (
                        <Nodes
                            match={match}
                            location={location}
                            data={complianceControlNodes}
                            entityContext={{ ...entityContext, [entityTypes.CONTROL]: id }}
                        />
                    );
                }

                const {
                    standardId = '',
                    name = '',
                    description = '',
                    interpretationText = ''
                } = entity;
                const nodes = complianceControlNodes.map(node => {
                    const {
                        id: nodeId,
                        name: nodeName,
                        nodeComplianceControlCount,
                        clusterName
                    } = node;
                    return {
                        id: nodeId,
                        name: nodeName,
                        clusterName,
                        passing: !nodeComplianceControlCount.failingCount
                    };
                });

                return (
                    <div className="w-full" id="capture-dashboard-stretch">
                        <CollapsibleSection title="Control Details">
                            <div className="flex flex-wrap pdf-page">
                                <ControlDetails
                                    standardId={standardId}
                                    control={name}
                                    description={description}
                                    className="mx-4 min-w-48 h-48 mb-4"
                                />
                                {!!interpretationText.length && (
                                    <Widget
                                        className="mx-4 min-w-48 h-48 mb-4 w-1/3 overflow-auto"
                                        header="Control guidance"
                                    >
                                        <div className="p-4 leading-loose whitespace-pre-wrap overflow-auto">
                                            {interpretationText}
                                        </div>
                                    </Widget>
                                )}
                                <RelatedEntityListCount
                                    className="mx-4 min-w-48 h-48 mb-4"
                                    name="Nodes"
                                    value={nodes.length}
                                    entityType={entityTypes.NODE}
                                />
                            </div>
                        </CollapsibleSection>
                        {!(entityContext && entityContext[entityTypes.NODE]) && (
                            <CollapsibleSection title="Control Findings">
                                <div className="flex pdf-page pdf-stretch shadow rounded relative rounded bg-base-100 mb-4 ml-4 mr-4">
                                    <NodesWithFailedControls
                                        entityType={entityTypes.CONTROL}
                                        entityContext={{
                                            ...entityContext,
                                            [entityTypes.CONTROL]: id
                                        }}
                                    />
                                </div>
                            </CollapsibleSection>
                        )}
                    </div>
                );
            }}
        </Query>
    );
};

Control.propTypes = entityComponentPropTypes;
Control.defaultProps = entityComponentDefaultProps;

export default Control;
