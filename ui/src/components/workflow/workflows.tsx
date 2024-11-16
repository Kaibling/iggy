import '../list.css';
import { getApiData } from '../utils/api.tsx'
import { useState, useEffect } from 'react';
import workflow_list from '../utils/list.tsx';
import { Workflow } from '../utils/types.tsx'
import { showError,NotifyContainer } from '../utils/notify.tsx';
import { useNavigate } from 'react-router-dom';
import {sendApi} from '../utils/api.tsx'

const WorkflowsList = () => {
  const [workflowList, setWorkflowList] = useState<Workflow[] | null>(null);


  const fetchData = async () => {
    const dataResponse = await getApiData("workflows","?depth=10");
    if (dataResponse.error) {
      showError(String(dataResponse.error));
      return;
    }
    setWorkflowList(dataResponse.response);
    document.title = "Workflows - iggy";
  };
  useEffect(() => {
    fetchData()
  }, []);

  const navigate = useNavigate();

  const handleButtonClick = () => {
    navigate('/workflows/new');
  };

  const handleChildRemoveClick = async (child_id : string,parent_id: string) => {
    console.log(`delete child clicked for child ${child_id} and parent ${parent_id}`);
    const postData = { "items": [child_id] };
    await sendApi(`workflows/${parent_id}/remove-children`,"POST",postData);
    fetchData();
  };

const handleDeleteClick = async (id: String) => {
    console.log(`delete clicked for ${id}`);
    await sendApi(`workflows/${id}`,"DELETE","");
    fetchData();
  };

  return (
    <div>
      <NotifyContainer/>
      <div className='top-banner' >
        <div><h2>Workflows</h2></div>
        
        <div className='list-actions'>
          <button id='add-button'
            onClick={handleButtonClick}
            style={{ height: '50px' }}
          >New</button>
        </div>
      </div>
      {workflow_list(workflowList,handleDeleteClick,handleChildRemoveClick)}
    </div>
  );
};

export default WorkflowsList;

