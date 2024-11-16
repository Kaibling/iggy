import React from "react";
import '../utils/list.css';
import { Run } from '../utils/types'
import { formatDate } from '../utils/api'

const RunCard: React.FC<{ run: Run }> = ({ run}) => {
  return (
    <>
      <tr>
        <td>{formatDate(run.start_time)}</td>
        <td>{formatDate(run.finish_time)}</td>
        <td>{run.run_time}</td>
        <td>{formatDate(run.meta.created_at)}</td>
        <td>{run.meta.created_by}</td>
        <td>{formatDate(run.meta.modified_at)}</td>
        <td>{run.meta.modified_by}</td>
      </tr>
    </>
  );
};


// Main component to render the workflow list
const run_list = (runs: Run[] | null ) => {

  if (runs == null || runs.length == 0) {
    return (
      <div className="site-container">
        <table className="item-table">
          <thead>
            <tr>
              <th>Start Time</th>
              <th>Finish Time</th>
              <th>Duration</th>
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
          <th>Start Time</th>
              <th>Finish Time</th>
              <th>Duration</th>
              <th>Created At</th>
              <th>Created By</th>
              <th>Modified At</th>
              <th>Modified By</th>
          </tr>
        </thead>
        <tbody>
          {runs.map((run: Run) => (
            <RunCard key={run.id} run={run} />
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default run_list;