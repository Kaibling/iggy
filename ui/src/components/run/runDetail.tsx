import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { getApiData } from '../utils/api.tsx';

import { Run } from '../utils/types.tsx'
import '../detail.css'
import { NotifyContainer, showError } from '../utils/notify.tsx';


const RunDetail: React.FC = () => {
    const { id } = useParams<{ id: string | undefined }>();
    const [run, setRun] = useState<Run | null>(null);
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
                <h1 >{run.id}</h1>
            </div>

            <div className="detail-container-row">

                    <div className="detail-card">
                        <h2 >Details</h2>

                        <div className="detail-row">
                            <span className="detail-label">ID:</span>
                            <span className="detail-value">{run.id}</span>
                        </div>

                        <div className="detail-row">
                            <span className="detail-label">name:</span>
                            <span className="detail-value">{run.id}</span>
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

                        <div className="detail-row">
                            <span className="detail-label">Created at:</span>
                            <span className="detail-value">{run.meta.created_at}</span>
                        </div>
                        <div className="detail-row">
                            <span className="detail-label">Created by:</span>
                            <span className="detail-value">{run.meta.created_by}</span>
                        </div>
                        <div className="detail-row">
                            <span className="detail-label">Updated_at:</span>
                            <span className="detail-value">{run.meta.modified_at}</span>
                        </div>
                        <div className="detail-row">
                            <span className="detail-label">Updated_by:</span>
                            <span className="detail-value">{run.meta.modified_by}</span>
                        </div>
                    </div>
            </div>
        </div>
    );
};

export default RunDetail;