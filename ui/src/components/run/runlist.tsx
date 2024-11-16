import '../utils/list.css';
import { getApiData } from '../utils/api.tsx'
import { useState, useEffect } from 'react';
import { Run } from '../utils/types.tsx'
// import { showError } from '../utils/notify.tsx';
import run_list from './run_list.tsx';


const RunsList = () => {
  const [runList, setRunList] = useState<Run[] | null>(null);


  const fetchData = async () => {
    const dataResponse = await getApiData("runs");

    if (dataResponse.error) {
      //showError(String(dataResponse.error));
      return;
    }
    setRunList(dataResponse.response);
    document.title = "Runs - iggy";
  };

  useEffect(() => {
    fetchData()
  }, []);


  return (
    <div>
      <div className='top-banner' >
        <div><h2>Runs</h2></div>
      </div>
      {run_list(runList)}
    </div>
  );
};

export default RunsList;

