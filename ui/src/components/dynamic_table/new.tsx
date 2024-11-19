import { useState, ChangeEvent, FormEvent } from 'react';
import { NewDynamicTable } from '../utils/types'
import { useNavigate } from 'react-router-dom';
import { sendApiData } from '../utils/api'
import { NotifyContainer, showError } from '../utils/notify';

const DynamicTableNew = () => {

  const [formData, setFormData] = useState<NewDynamicTable>({
    name: "",
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
    console.log(name);
    console.log(value);
    console.log(formData);
  };

  const navigate = useNavigate();
  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log('Form submitted:', formData);
    const error = await sendApiData("dynamic-tables", [formData])
    if (error !== false) {
      showError(error.error)
    } else {
      navigate('/dynamic-tables');
    }
  };

  return (
    <div className="detail-container">
      <NotifyContainer />
      <div className='detail-card'>
    <h2>Create DynamicTable</h2>

        <form onSubmit={handleSubmit}>
          <div className="detail-row">
            <label className="detail-label" htmlFor="name">name </label>
            <input
              type="text"
              id="name"
              name="name"
              value={formData.name}
              onChange={handleChange}
              required
            />
          </div>

          <button type="submit">Add</button>
        </form>

      </div>
    </div>
  );
}

export default DynamicTableNew;