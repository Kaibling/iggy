import React from "react";
import '../utils/list.css';
import { DynamicTable } from '../utils/types'
import { formatDate } from '../utils/api'

const DynamicTableCard: React.FC<{ dynamicTable: DynamicTable }> = ({ dynamicTable}) => {
  return (
    <>
      <tr>
        <td>{dynamicTable.table_name}</td>
        <td>{formatDate(dynamicTable.meta.created_at)}</td>
        <td>{dynamicTable.meta.created_by}</td>
        <td>{formatDate(dynamicTable.meta.modified_at)}</td>
        <td>{dynamicTable.meta.modified_by}</td>
      </tr>
    </>
  );
};

// Main component to render the workflow list
const list = (runs: DynamicTable[] | null ) => {
  if (runs == null || runs.length == 0) {
    return (
      <div className="site-container">
        <table className="item-table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Created At</th>
              <th>Created By</th>
              <th>Modified At</th>
              <th>Modified By</th>
            </tr>
          </thead>
        </table>
        <div id="no-data">No Data found</div>
      </div>
    )
  }
  return (
    <div className="site-container">
      <table className="item-table">
        <thead>
          <tr>
              <th>Name</th>
              <th>Created At</th>
              <th>Created By</th>
              <th>Modified At</th>
              <th>Modified By</th>
          </tr>
        </thead>
        <tbody>
          {runs.map((dynamicTable: DynamicTable) => (
            <DynamicTableCard key={dynamicTable.id} dynamicTable={dynamicTable} />
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default list;