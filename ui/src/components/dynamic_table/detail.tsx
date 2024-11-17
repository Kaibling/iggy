import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { getApiData } from '../utils/api.tsx';
import { formatDate } from '../utils/api.tsx'
import { Run, DynamicTable } from '../utils/types.tsx'
import '../detail.css'
import { NotifyContainer, showError } from '../utils/notify.tsx';




// const DynamicTableCard: React.FC<{ runLog: DynamicTable }> = ({ runLog }) => {
//     return (
//         <>
//             <tr>
//                 {/* <td>{runLog.id}</td> */}
//                 <td>{formatDate(runLog.timestamp)}</td>
//                 <td>{runLog.message}</td>

//             </tr>
//         </>
//     );
// };

const RunDetail: React.FC = () => {
    const { id } = useParams<{ id: string | undefined }>();
    // const [run, setRun] = useState<Run | null>(null);
    const [dynamicTable, setDynamicTable] = useState<DynamicTable | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);


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

            // const logsResponse = await getApiData("runs/" + id + "/logs");
            // if (logsResponse.error) {
            //     showError(String(logsResponse.error));
            //     return;
            // }
            // setLogs(logsResponse.response);

            document.title = dynamicTableResponse.response?.id + " - iggy" || "Workflow - iggy";
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
            <div className='top-banner' >
                <h2 >{dynamicTable.table_name}</h2>
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
                        <span className="detail-value">{dynamicTable.table_name}</span>
                    </div>
                    <div className="detail-row">
                        <span className="detail-label">Created at:</span>
                        <span className="detail-value">{dynamicTable.meta.created_at}</span>
                    </div>
                    <div className="detail-row">
                        <span className="detail-label">Created by:</span>
                        <span className="detail-value">{dynamicTable.meta.created_by}</span>
                    </div>
                    <div className="detail-row">
                        <span className="detail-label">Updated_at:</span>
                        <span className="detail-value">{dynamicTable.meta.modified_at}</span>
                    </div>
                    <div className="detail-row">
                        <span className="detail-label">Updated_by:</span>
                        <span className="detail-value">{dynamicTable.meta.modified_by}</span>
                    </div>
                </div>

            </div>
            {/* <div className="detail-container-row">
                <div className="detail-card-max-width">
                    <h2 >Logs</h2>
                    <table>
                        <tbody>
                            {dynamicTable?.map((log: DynamicTable) => (
                                <DynamicTableCard key={log.id} runLog={log} />
                            ))}
                        </tbody>

                    </table>
                </div>
            </div> */}
        </div>
    );
};

export default RunDetail;