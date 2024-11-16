import React, { useState, useEffect } from 'react';
import SearchDropdown from '../utils/SearchDropdown';
import { getApiData, sendApiData } from '../utils/api'
import { Workflow } from '../utils/types';
import { NotifyContainer, showError } from '../utils/notify';
import { Option } from '../utils/types';


interface FormWithDropdownProps {
  onClose: () => void;
  workflow_id: string | undefined;
}

const FormAddWorkflow: React.FC<FormWithDropdownProps> = ({ onClose, workflow_id }) => {
  const [selectedWorkflowOption, setSelectedWorkflowOption] = useState<Option | null>(null);
  const [workflows, setWorkflows] = useState<Option[] | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      const apiResponse = await getApiData("workflows");
        if (apiResponse.response) {
            const workflowData: Workflow[] = apiResponse.response;
            const option_array : Option[]= workflowData.map(wf => {
                return {
                    name:wf.name,
                    id:wf.id,
                } as Option;
            });
            setWorkflows(option_array);
        } else {
            console.log(apiResponse.error)
        }
    }
    fetchData()
  }, []);

  const handleWorkflowSelect = (option: Option) => {
    setSelectedWorkflowOption(option);
  };


  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    const postData = { "items": [selectedWorkflowOption?.id] };
    const error = await sendApiData("workflows/"+workflow_id+"/add-children", postData);
    if (error !== false) {
      showError(error.error);
    } else {
      onClose();
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <NotifyContainer />
        <h3>Append a Workflow</h3>
        <label>
          Name
        </label>
        <SearchDropdown options={workflows} onSelect={handleWorkflowSelect} placeholder='Search Workflow' />
        {selectedWorkflowOption?.name && <p>Selected Option: {selectedWorkflowOption.name}</p>}
      </div>

      <button type="submit">Add</button>
    </form>
  );
};

export default FormAddWorkflow;