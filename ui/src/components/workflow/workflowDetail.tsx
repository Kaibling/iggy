import { useParams } from 'react-router-dom';
import { useState, useEffect, ChangeEvent, FormEvent } from 'react';
import { getApiData, sendApi } from '../utils/api.tsx';

import { Workflow, NewWorkflow } from '../utils/types.tsx'
import GenTable from '../utils/table.tsx'
import { WorkflowFieldsSmall } from '../utils/constant.tsx';
import '../detail.css'
import Editor from '@monaco-editor/react';
import { NotifyContainer, showError } from '../utils/notify.tsx';
import FormAddWorkflow from './formAddWorkflow.tsx'
import Modal from '../utils/modal.tsx';
import { ActionButtonProps } from '../utils/actionButton.tsx'


const WorkflowDetail: React.FC = () => {
    const { id } = useParams<{ id: string | undefined }>();
    const [workflow, setWorkflow] = useState<Workflow | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [code, setCode] = useState<string | null>(null);
    const [isModalOpen, setIsModalOpen] = useState(false);


    const [formData, setFormData] = useState<NewWorkflow>({
        name: "",
        code: "",
        object_type: "",
        fail_on_error: false
    });

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value, type, checked } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: type === 'checkbox' ? checked : value,
        }));
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        console.log('Code in Editor:', code);
        e.preventDefault();
        if (code != null) {
            formData.code = code;
        }

        console.log('Form submitted:', formData);
        const error = await sendApi("workflows/" + id, "PATCH", formData)
        if (error !== false) {
            showError(error.error)
        } else {
            //reload
            fetchData();
        }
    };

    const handleSelectChange = (e: ChangeEvent<HTMLSelectElement>) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
          ...prevData,
          [name]: value,
        }));
        console.log(formData);
      };

    function handleEditorChange(value: string, event: any) {
        setCode(value);
    }

    const handleOpenModal = () => {
        setIsModalOpen(true);
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
        fetchData();
    };


    const fetchData = async () => {

        try {
            setLoading(true);
            if (!id) {
                throw new Error('No ID provided');
            }
            const workflowResponse = await getApiData("workflows/" + id);
            if (workflowResponse.error) {
                showError(String(workflowResponse.error));
                return;
            }
            setWorkflow(workflowResponse.response);
            document.title = workflowResponse.response?.name + " - iggy" || "Workflow - iggy";
            setFormData({
                name: workflowResponse.response?.name,
                code: workflowResponse.response?.code,
                object_type: workflowResponse.response?.object_type,
                fail_on_error: workflowResponse.response?.fail_on_error
            });
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



    const actions: ActionButtonProps[] = [
        { id: '2', actionType: 'remove-child', collection: 'workflows', postFunction: fetchData }
    ];


    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    if (!workflow) {
        return <div>No workflow found</div>;
    }

    return (
        // <div className="detail-container">
        <div>
            <Modal show={isModalOpen} onClose={handleCloseModal}>
                <FormAddWorkflow onClose={handleCloseModal} workflow_id={id} />
            </Modal>
            <NotifyContainer />
            <div className='top-banner' >
                <h1 >{workflow.name}</h1>
                <div className='list-actions'>
                    <button onClick={handleOpenModal}>Add Workflow</button>
                    <button form='wf_form' type="submit">Update</button>
                </div>
            </div>

            <div className="detail-container-row">



                <div className="detail-card">
                    <form id='wf_form' onSubmit={handleSubmit}>
                        <h2 >Details</h2>

                        <div className="detail-row">
                            <span className="detail-label">ID:</span>
                            <span className="detail-value">{workflow.id}</span>
                        </div>

                        <div className="detail-row">
                            <span className="detail-label">name:</span>
                            <span className="detail-value">
                                <input
                                    type="text"
                                    id="name"
                                    name="name"
                                    value={formData.name}
                                    onChange={handleChange}
                                    required
                                /></span>
                        </div>

                        <div className="detail-row">
                            <span className="detail-label">object_type:</span>
                            <span className="detail-value">
                                <select id="object_type" name="object_type"
                                    value={formData.object_type}
                                    onChange={handleSelectChange}
                                    required
                                >
                                    <option value="javascript">Javascript</option>
                                    <option value="external">External</option>
                                    <option value="folder">Folder</option>
                                </select>
                            </span>
                        </div>


                        <div className="detail-row">
                            <span className="detail-label">fail_on_error:</span>
                            <span className="detail-value">   <input
                                type="checkbox"
                                id="fail_on_error"
                                name="fail_on_error"
                                checked={formData.fail_on_error}
                                onChange={handleChange}
                            /></span>
                        </div>

                        <div className="detail-row">
                            <span className="detail-label">Created at:</span>
                            <span className="detail-value">{workflow.created_at}</span>
                        </div>
                        <div className="detail-row">
                            <span className="detail-label">Created by:</span>
                            <span className="detail-value">{workflow.created_by}</span>
                        </div>
                        <div className="detail-row">
                            <span className="detail-label">Updated_at:</span>
                            <span className="detail-value">{workflow.modified_at}</span>
                        </div>
                        <div className="detail-row">
                            <span className="detail-label">Updated_by:</span>
                            <span className="detail-value">{workflow.modified_by}</span>
                        </div>
                    </form>
                </div>



                <div className="detail-card-max-width">
                    <h2>Sub workflows</h2>
                    {GenTable(workflow.children, WorkflowFieldsSmall, actions)}
                </div>
            </div>

            <div className="detail-card-max-width">
                <h2>Code</h2>
                <div className="editor-container">
                    <Editor
                        height="60vh"
                        width="100%"
                        defaultLanguage="javascript"
                        theme="vs-light"
                        defaultValue={workflow.code}
                        onChange={handleEditorChange}
                    />

                </div>
            </div>



        </div>
    );
};

export default WorkflowDetail;