import { useState, ChangeEvent, FormEvent } from 'react';
import { NewWorkflow } from '../utils/types'
import { useNavigate } from 'react-router-dom';
import { sendApiData } from '../utils/api'
import { NotifyContainer, showError } from '../utils/notify';

const WorkflowNew = () => {

  const [formData, setFormData] = useState<NewWorkflow>({
    name: "",
    code: "",
    object_type: "javascript",
    fail_on_error: false
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: type === 'checkbox' ? checked : value,
    }));
    console.log(formData);
  };

  const handleSelectChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const { name, value } = e.target;
    console.log(name);
    console.log(value);
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
    console.log(formData);
  };
  const navigate = useNavigate();
  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log('Form submitted:', formData);
    const error = await sendApiData("workflows", [formData])
    if (error !== false) {
      showError(error.error)
    } else {
      navigate('/workflows');
    }
  };

  return (
    <div className="detail-container">
      <NotifyContainer />
      <div className='detail-card'>
    <h2>Create Workflow</h2>

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
          <div className="detail-row">
            <label className="detail-label" htmlFor="code">code</label>
            <input
              type="text"
              id="code"
              name="code"
              value={formData.code}
              onChange={handleChange}
            />
          </div>
          <div className="detail-row">
            <label className="detail-label" htmlFor="object_type">Type</label>
            <select id="object_type" name="object_type"
             value={formData.object_type}
             onChange={handleSelectChange}
             required
            >
              <option value="javascript">Javascript</option>
              <option value="external">External</option>
              <option value="folder">Folder</option>
            </select>
          </div>
          <div className="detail-row">
            <label className="detail-label" htmlFor="fail_on_error">Fail on Error:</label>
            <input
              type="checkbox"
              id="fail_on_error"
              name="fail_on_error"
              checked={formData.fail_on_error}
              onChange={handleChange}
            />
          </div>

          <button type="submit">Add Workflow</button>
        </form>

      </div>
    </div>
  );
}

export default WorkflowNew;