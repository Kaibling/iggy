import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { getApiData } from '../utils/api.tsx';
import { DynamicTableVariable, DynamicTable } from '../utils/types.tsx'
import '../detail.css'
import { NotifyContainer, showError } from '../utils/notify.tsx';
import Modal from '../utils/modal.tsx';
import FormAddDynamicTable from './formAdddyntab.tsx';
import { formatDate } from '../utils/api'




const VariableCard: React.FC<{ variable: DynamicTableVariable }> = ({ variable }) => {
    return (
        <>
            <tr>
                <td>{variable.name}</td>
                <td>{variable.variable_type}</td>
                    <td>{formatDate(variable.meta.created_at)}</td>
                    <td>{variable.meta.created_by}</td>
                    <td>{formatDate(variable.meta.modified_at)}</td>
                    <td>{variable.meta.modified_by}</td>

            </tr>
        </>
    );
};

const RunDetail: React.FC = () => {
    const { id } = useParams<{ id: string | undefined }>();
    const [vars, setVars] = useState<DynamicTableVariable[] | null>(null);
    const [dynamicTable, setDynamicTable] = useState<DynamicTable | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [isModalOpen, setIsModalOpen] = useState(false);

    const fetchData = async () => {
        try {
            setLoading(true);
            if (!id) {
                throw new Error('No ID provided');
            }
            const dynamicTableResponse = await getApiData("dynamic-tables/" + id);
            if (dynamicTableResponse.error) {
                showError(String(dynamicTableResponse.error));
                return;
            }
            setDynamicTable(dynamicTableResponse.response);

            const varsResponse = await getApiData("dynamic-tables/" + id + "/variables");
            if (varsResponse.error) {
                showError(String(varsResponse.error));
                return;
            }
            setVars(varsResponse.response);

            document.title = dynamicTableResponse.response?.id + " - iggy" || "Dynamic Tables - iggy";
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message);
            } else {
                setError('An unexpected error occurred');
            }
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchData();
    }, [id]);

    const handleOpenModal = () => {
        setIsModalOpen(true);
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
        fetchData();
    };

    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    if (!dynamicTable) {
        return <div>No run found</div>;
    }

    return (
        // <div className="detail-container">
        <div>
            <NotifyContainer />

            <Modal show={isModalOpen} onClose={handleCloseModal}>
                <FormAddDynamicTable onClose={handleCloseModal} dyn_tab_id={id} />
            </Modal>

            <div className='top-banner' >
                <h1 >{dynamicTable.name}</h1>
                <div className='list-actions'>

                    <button form='wf_form' type="submit">Update</button>
                </div>
            </div>

            <div className="detail-container-row">

                <div className="detail-card">
                    <h2 >Details</h2>

                    <div className="detail-row">
                        <span className="detail-label">ID:</span>
                        <span className="detail-value">{dynamicTable.id}</span>
                    </div>

                    <div className="detail-row">
                        <span className="detail-label">Name:</span>
                        <span className="detail-value">{dynamicTable.name}</span>
                    </div>
                    <div className="detail-row">
                        <span className="detail-label">Created at:</span>
                        <span className="detail-value">{formatDate(dynamicTable.meta.created_at)}</span>
                    </div>
                    <div className="detail-row">
                        <span className="detail-label">Created by:</span>
                        <span className="detail-value">{dynamicTable.meta.created_by}</span>
                    </div>
                    <div className="detail-row">
                        <span className="detail-label">Updated_at:</span>
                        <span className="detail-value">{formatDate(dynamicTable.meta.modified_at)}</span>
                    </div>
                    <div className="detail-row">
                        <span className="detail-label">Updated_by:</span>
                        <span className="detail-value">{dynamicTable.meta.modified_by}</span>
                    </div>
                </div>

            </div>
            <div className="detail-container-row">
                <div className="detail-card-max-width">
                    <h2 >Schema</h2>
                   
                    <table>
                        <thead>
                            <th>Name</th>
                            <th>Type</th>
                            <th>Created At</th>
                            <th>Created By</th>
                            <th>Modified At</th>
                            <th>Modified By</th>
                        </thead>
                        <tbody>
                            {vars?.map((variable: DynamicTableVariable) => (
                                <VariableCard key={variable.id} variable={variable} />
                            ))}
                        </tbody>

                    </table>
                    <button onClick={handleOpenModal}>Add Field</button>
                </div>
            </div>
        </div>
    );
};

export default RunDetail;