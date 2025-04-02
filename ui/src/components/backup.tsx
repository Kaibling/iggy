import "./backup.css";
import { getApiData,sendApiData } from './utils/api.tsx';
import { UIConfig } from './utils/types.tsx';
import { NotifyContainer,showSuccess, showError } from './utils/notify';

import { useState, useEffect } from 'react';

const Backup = () => {
  const [uiconfig, setUIConfig] = useState<UIConfig | null>(null);
  const [task, setTask] = useState<string>("import");
  // const [path, setPath] = useState<string>("");

  // Fetch data from the API
  const fetchData = async () => {
    const dataResponse = await getApiData("ui-config");

    if (dataResponse.error) {
      // showError(String(dataResponse.error));
      return;
    }

    setUIConfig(dataResponse.response);
    // setPath(dataResponse.response.ImportLocalPath);
  };

  useEffect(() => {
    fetchData();
  }, []);

  // Handle task selection change
  const handleTaskChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedTask = event.target.value;
    setTask(selectedTask);

    // if (uiconfig) {
    //   setPath(selectedTask === "import" ? uiconfig.ImportLocalPath : uiconfig.ExportLocalPath);
    // }
  };

  // Allow editing of the path input
  // const handlePathChange = (event: React.ChangeEvent<HTMLInputElement>) => {
  //   setPath(event.target.value);
  // };

  // Execute function on form submission
  const execute = async (formData: { task: string; path: string }) => {
  const error = await sendApiData(`backup/${task}`,null);
    if (error !== false) {
      showError(error.error);
    } else {
      if (task == "import") {
        showSuccess("Data imported");

      } else {
        showSuccess("Data exported");
      }
    }
  };

  // Handle form submission
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault(); // Prevent page refresh
    execute({ task, path:"" });
  };

  return (
    <div className="container">
      <NotifyContainer/>
      <h2>Import/Export</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="task">Task</label>
          <select
            id="task"
            className="input-field"
            value={task}
            onChange={handleTaskChange}
          >
            <option value="import">Import</option>
            <option value="export">Export</option>
          </select>
        </div>

        {/* <div className="form-group">
          <label htmlFor="path">Path</label>
          <input
            id="path"
            type="text"
            className="input-field"
            value={path}
            onChange={handlePathChange}
          />
        </div> */}

        <button type="submit" className="button">
          Execute
        </button>
      </form>
    </div>
  );
};

export default Backup;
