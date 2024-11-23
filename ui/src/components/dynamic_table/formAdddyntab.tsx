import React, { useState, ChangeEvent } from 'react';
import { sendApiData } from '../utils/api'
import { NewDynamicSchema } from '../utils/types';
import { NotifyContainer, showError } from '../utils/notify';

interface Form {
  onClose: () => void;
  dyn_tab_id: string | undefined;
}

const FormAddDynamicTable: React.FC<Form> = ({ onClose, dyn_tab_id }) => {
  const [formData, setFormData] = useState<NewDynamicSchema>({
    name: "",
    variable_type: "",
    dynamic_table_id: "",
  });

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    formData.dynamic_table_id = dyn_tab_id!;
    const error = await sendApiData("dynamic-tables/" + dyn_tab_id + "/fields", [formData]);
    if (error !== false) {
      showError(error.error);
    } else {
      onClose();
    }
  };

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <NotifyContainer />
        <h3>add field</h3>

        <div className="detail-row">
          <span className="detail-label">Name</span>
          <span className="detail-value">   <input
            type="input"
            id="name"
            name="name"
            value={formData.name}
            onChange={handleChange}
          /></span>
        </div>
        <div className="detail-row">
          <span className="detail-label">type</span>
          <span className="detail-value">   <input
            type="input"
            id="variable_type"
            name="variable_type"
            value={formData.variable_type}
            onChange={handleChange}
          /></span>
        </div>



      </div>

      <button type="submit">Add</button>
    </form>
  );
};

export default FormAddDynamicTable;