import React, { useState } from "react";
import './list.css';
import arrow_img from '../../assets/right-arrow.svg';
import { Workflow } from './types'
import false_img from '../../assets/cancel.svg'
import true_img from '../../assets/success.svg'
import folder_img from '../../assets/folder.svg'
import js_img from '../../assets/js.svg'
import os_img from '../../assets/bash.svg'
import { formatDate } from './api'
import deleteImage from '../../assets/trash.svg'; 
import deleteChildImage from '../../assets/trash-child.svg'; 
import playImage from '../../assets/play.svg'; 
import { sendApi } from "./api";

const WorkflowCard: React.FC<{ workflow: Workflow ,parentId?: string,handleDeleteClick: (id: string) => void,handlechildRemoveClick: (child_id: string,parent_id: string) => void}> = ({ workflow,parentId,handleDeleteClick, handlechildRemoveClick}) => {
  const [showChildren, setShowChildren] = useState<boolean>(false);

  const toggleChildren = () => {
    setShowChildren(!showChildren);
  };

  return (
    <>
      <tr>
        <td>
          {workflow.children && workflow.children.length > 0 && (
            <button onClick={toggleChildren} className="button-toggle">
              <img
                src={arrow_img}
                alt="Arrow"
                className={`arrow-img ${showChildren ? 'rotated' : ''}`}
                style={{ width: '20px', height: '20px' }}
              />
            </button>
          )}
        </td>
        <td><a href={`/workflows/${workflow.id}`}> {workflow.name || "N/A"}</a></td>
        <td>{workflow.object_type == "javascript" ? (
          <img src={js_img} alt="javascript" style={{ width: "30px", height: "30px" }} />
        ) : workflow.object_type == "external" ? (
          <img src={os_img} alt="external" style={{ width: "30px", height: "30px" }} />
        ) : workflow.object_type == "folder" ? (
          <img src={folder_img} alt="folder" style={{ width: "35px", height: "35px" }} />
        ) : (
          workflow.object_type
        )}</td>
        <td>{workflow.fail_on_error ? (
          <img src={true_img} alt="True" style={{ width: "30px", height: "30px" }} />
        ) : (
          <img src={false_img} alt="False" style={{ width: "30px", height: "30px" }} />
        )
        }</td>
        <td>{formatDate(workflow.meta.created_at)}</td>
        <td>{workflow.meta.created_by}</td>
        <td>{formatDate(workflow.meta.modified_at)}</td>
        <td>{workflow.meta.modified_by}</td>
        <td>
          <button onClick={() => handleRunClick(workflow.id)} style={{ border: 'none', background: 'none' }}>
            <img
              src={playImage}
              alt="Run"
              title={'run workflow'}
              style={{ width: '25px', height: '25px' }} />
          </button>
          <button onClick={() => {
            if (parentId) {
              handlechildRemoveClick(workflow.id, parentId); 
            } else {
              handleDeleteClick(workflow.id);
            }
          }} style={{ border: 'none', background: 'none' }}>
                      <img
              src={parentId ? deleteChildImage : deleteImage} // Use deleteChildImage for children
              alt="Delete"
              title={parentId ? 'Delete child workflow' : 'Delete workflow'}
              style={{ width: '25px', height: '25px' }} />
          </button>
            </td>
      </tr>

      {showChildren && workflow.children && workflow.children.map((child: Workflow) => (
        <WorkflowCard key={child.id} workflow={child} parentId={workflow.id} handleDeleteClick={handleDeleteClick} handlechildRemoveClick={handlechildRemoveClick} />
      ))}
    </>
  );
};


// Main component to render the workflow list
const workflow_list = (workflows: Workflow[] | null,handleDeleteClick: (id: string) => void,handlechildRemoveClick: (child_id: string,parent_id: string) => void) => {

  if (workflows == null || workflows.length == 0) {
    return (
      <div className="site-container">
        <table className="workflow-table">
          <thead>
            <tr>
              <th></th>
              <th>Name</th>
              <th>Type</th>
              <th>Fail on error</th>
              <th>Created At</th>
              <th>Created By</th>
              <th>Modified At</th>
              <th>Modified By</th>
              <th>Actions</th>
            </tr>
          </thead>
        </table>
        <div id="no-data">No Data found</div>
      </div>
    )
  }
  return (
    <div className="site-container">
      <table className="workflow-table">
        <thead>
          <tr>
            <th></th>
            <th>Name</th>
            <th>Type</th>
            <th>Fail on error</th>
            <th>Created At</th>
            <th>Created By</th>
            <th>Modified At</th>
            <th>Modified By</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {workflows.map((workflow: Workflow) => (
            <WorkflowCard key={workflow.id} workflow={workflow} handleDeleteClick={handleDeleteClick} handlechildRemoveClick={handlechildRemoveClick} />
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default workflow_list;

  const handleRunClick = (id: string) => {
    console.log(`run clicked for ${id}`);
    sendApi(`workflows/${id}/run`,"POST","");
  };
