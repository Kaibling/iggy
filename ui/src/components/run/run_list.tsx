import React from "react";
import '../utils/list.css';
import { Run } from '../utils/types'
import { formatDate } from '../utils/api'

const RunCard: React.FC<{ run: Run }> = ({ run}) => {
  return (
    <>
      <tr>
        <td>{run.workflow.name}</td>
        <td><a href={`/runs/${run.id}`}>{formatDate(run.start_time)}</a></td>
        <td>{formatDate(run.finish_time)}</td>
        <td>{run.run_time}</td>
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
              <th>Workflow</th>
              <th>Start Time</th>
              <th>Finish Time</th>
              <th>Duration</th>
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
              <th>Workflow</th>
              <th>Start Time</th>
              <th>Finish Time</th>
              <th>Duration</th>
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