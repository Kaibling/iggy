// import './list.css'
import './table.css'
import { Field } from './types.tsx'
import { ActionButtonProps, ActionButton } from './actionButton.tsx'
import false_img from '../../assets/cancel.svg'
import true_img from '../../assets/success.svg'
import folder_img from '../../assets/folder.svg'
import js_img from '../../assets/js.svg'
import os_img from '../../assets/bash.svg'
import { formatDate } from './api.tsx'
interface Date {
  [key: string]: any;
  meta: meta;
}
interface meta {
  created_at: string;
  modified_at: string;
}


const GenTable = (data: Date[] | null, fields: Field[], actions: ActionButtonProps[] = []) => {
  if (data == null || data.length == 0) {
    return (
      <div className="table-container">
        <table className="user-table">
          <thead>
            <tr>
              {fields.map((field, id) => (
                <th key={id}>{field.displayName}</th>
              ))}
            </tr>
          </thead>
        </table>
        <div id="no-data">No Data found</div>
      </div>
    )
  }
  return (
    <div className="table-container">
      <table className="user-table">
        <thead>
          <tr>
            {fields.map((field, id) => (
              <th key={id}>{field.displayName}</th>
            ))}
            {actions.length > 0 && (
              <th>Actions</th>
            )}
          </tr>
        </thead>
        <tbody>
          {data?.map((row, rowIndex) => (
            <tr
              key={rowIndex}
              style={row["active"] === false ? { color: "grey" } : {}}
            >
              {fields.map((field, id) => (
                <td key={id}>
                  {field.boolean !== undefined ? (
                    row[field.fieldName] ? (
                      <img src={true_img} alt="True" style={{ width: "30px", height: "30px" }} />
                    ) : (
                      <img src={false_img} alt="False" style={{ width: "30px", height: "30px" }} />
                    )
                  ) : field.collection ? (
                    <a href={`/${field.collection}/${field.identifier ? (
                      row[field.identifier].id
                    ) : (
                      row["id"]
                    )}`}>
                      {field.identifier ? (
                        row[field.identifier].name
                      ) : (
                        row[field.fieldName]
                      )}
                    </a>
                  ) : field.date ? (
                    formatDate(row[field.fieldName])
                  ) :field.symbol ? (
                    row[field.fieldName] === "javascript" ? (
                        <img src={js_img} alt="javascript" style={{ width: "30px", height: "30px" }} />
                      ): row[field.fieldName] == "external" ?(
                        <img src={os_img} alt="external" style={{ width: "30px", height: "30px" }} />
                      ): row[field.fieldName] == "folder" ? (
                        <img src={folder_img} alt="folder" style={{ width: "35px", height: "35px" }} />
                      ):(
                        row[field.fieldName]
                      )

                      
                  )
                    : (
                      row[field.fieldName]
                    )}
                </td>
              ))}

              {actions.length > 0 && (
                <td>

                  {/* <div className='detail-container-row'> */}
                  {actions.map((props, index) => (
                    <ActionButton
                      key={index}
                      id={row["id"]}
                      actionType={props.actionType}
                      collection={props.collection}
                      postFunction={props.postFunction}
                    />
                  ))}
                  {/* </div> */}
                </td>
              )}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default GenTable;
