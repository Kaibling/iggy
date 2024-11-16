import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { getApiData } from '../utils/api.tsx';
import { formatDate } from '../utils/api'
import { Run, RunLog } from '../utils/types.tsx'
import '../detail.css'
import { NotifyContainer, showError } from '../utils/notify.tsx';




const RunLogCard: React.FC<{ runLog: RunLog }> = ({ runLog }) => {
    return (
        <>
            <tr>
                {/* <td>{runLog.id}</td> */}
                <td>{formatDate(runLog.timestamp)}</td>
                <td>{runLog.message}</td>
                
            </tr>
        </>
    );
};

const RunDetail: React.FC = () => {
    const { id } = useParams<{ id: string | undefined }>();
    const [run, setRun] = useState<Run | null>(null);
    const [logs, setLogs] = useState<RunLog[] | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);


    const fetchData = async () => {
        try {
            setLoading(true);
            if (!id) {
                throw new Error('No ID provided');
            }
            const runResponse = await getApiData("runs/" + id);
            if (runResponse.error) {
                showError(String(runResponse.error));
                return;
            }
            setRun(runResponse.response);

            const logsResponse = await getApiData("runs/" + id + "/logs");
            if (logsResponse.error) {
                showError(String(logsResponse.error));
                return;
            }
            setLogs(logsResponse.response);

            document.title = runResponse.response?.id + " - iggy" || "Workflow - iggy";
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

    if (!run) {
        return <div>No run found</div>;
    }

    return (
        // <div className="detail-container">
        <div>
            <NotifyContainer />
            <div className='top-banner' >
                <h2 >{run.workflow.name} Run</h2>
            </div>

            <div className="detail-container-row">

                <div className="detail-card">
                    <h2 >Details</h2>

                    <div className="detail-row">
                        <span className="detail-label">ID:</span>
                        <span className="detail-value">{run.id}</span>
                    </div>

                    <div className="detail-row">
                        <span className="detail-label">Workflow:</span>
                        <span className="detail-value">{run.workflow.name}</span>
                    </div>

                    <div className="detail-row">
                        <span className="detail-label">error:</span>
                        <span className="detail-value">{run.error}</span>
                    </div>

                    <div className="detail-row">
                        <span className="detail-label">result:</span>
                        <span className="detail-value"> {run.result}</span>
                    </div>

                    <div className="detail-row">
                        <span className="detail-label">Run time:</span>
                        <span className="detail-value">{run.run_time}</span>
                    </div>
                </div>

            </div>
            <div className="detail-container-row">
                <div className="detail-card-max-width">
                    <h2 >Logs</h2>
                    <table>
                        <tbody>
                            {logs?.map((log: RunLog) => (
                                <RunLogCard key={log.id} runLog={log} />
                            ))}
                        </tbody>

                    </table>
                </div>
            </div>
        </div>
    );
};

export default RunDetail;