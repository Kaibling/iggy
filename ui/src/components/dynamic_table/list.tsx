import '../utils/list.css';
import { getApiData } from '../utils/api.tsx'
import { useState, useEffect } from 'react';
import { DynamicTable } from '../utils/types.tsx'
import list from './list_detail.tsx';
import { useNavigate } from 'react-router-dom';


const DynamicTableList = () => {
  const [dynamicTables, setDynamicTables] = useState<DynamicTable[] | null>(null);


  const fetchData = async () => {
    const dataResponse = await getApiData("dynamic-tables");
    console.log(dataResponse)
    if (dataResponse.error) {
      //showError(String(dataResponse.error));
      return;
    }
    setDynamicTables(dataResponse.response);
    document.title = "Dynamic Tables - iggy";
  };

  useEffect(() => {
    fetchData()
  }, []);

  const navigate = useNavigate();

  const handleButtonClick = () => {
    navigate('/dynamic-tables/new');
  };


  return (
    <div>
      <div className='top-banner' >
        <div><h2>Dynamic Tables</h2></div>
        <div className='list-actions'>
        <button id='add-button'
          onClick={handleButtonClick}
          style={{ height: '50px' }}
        >New</button>
      </div>
    </div>
      {list(dynamicTables)}
    </div>
  );
};

export default DynamicTableList;