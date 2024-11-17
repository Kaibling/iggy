import '../utils/list.css';
import { getApiData } from '../utils/api.tsx'
import { useState, useEffect } from 'react';
import { DynamicTable } from '../utils/types.tsx'
import list from './list_detail.tsx';


const DynamicTableList = () => {
  const [dynamicTables, setDynamicTables] = useState<DynamicTable[] | null>(null);


  const fetchData = async () => {
    const dataResponse = await getApiData("dynamic-tables");

    if (dataResponse.error) {
      //showError(String(dataResponse.error));
      return;
    }
    setDynamicTables(dataResponse.response);
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
      {list(dynamicTables)}
    </div>
  );
};

export default DynamicTableList;

