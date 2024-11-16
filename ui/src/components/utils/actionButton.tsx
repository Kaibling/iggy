import deleteImage from '../../assets/trash.svg'; 
import playImage from '../../assets/play.svg'; 
import { sendApi } from './api';

export interface ActionButtonProps {
    id: string;
    actionType: string;
    collection: string;
    postFunction: any; //() =>  Promise<void>;
    parentID?: string | null;
  }



export const ActionButton: React.FC<ActionButtonProps> = ({ id, actionType, collection,postFunction,parentID }) => {
    const handleDeleteClick = () => {
      // delete id
        console.log(`${collection} delete clicked for id ${id}`);
        sendApi(collection+ "/" + id,"DELETE","");
        postFunction();

      };
      const handleRunClick = () => {
        console.log(`${collection} run clicked for id ${id}`);
        sendApi(collection+ "/" + id + "/run","POST","");
      };
      const handleChildRemoveClick = () => {
        console.log(`${collection} child clicked for id ${id}`);
        const postData = { "items": [id] };
        sendApi(collection+ "/" + parentID + "/remove-children","POST",postData);
      };
    
    switch (actionType) {
      case 'run':
        return (
            <button onClick={handleRunClick} style={{ border: 'none', background: 'none' }}>
              <img 
              src={playImage} 
              alt="Run" 
              title={'run workflow'}
              style={{ width: '25px', height: '25px' }} />
            </button>
          );
      case 'delete':
        return (
        <button onClick={handleDeleteClick} style={{ border: 'none', background: 'none' }}>
        <img 
        src={deleteImage} 
        alt="Delete" 
        title={'delete workflow'}
        style={{ width: '25px', height: '25px' }} />
      </button>
        );
      case 'remove-child':
        return (
          <button onClick={handleChildRemoveClick} style={{ border: 'none', background: 'none' }}>
          <img 
          src={deleteImage} 
          alt="Remove" 
          title={'remove child'}
          style={{ width: '25px', height: '25px' }} />
        </button>
          );
      default:
        return <span>Unknown action</span>;
    }
  };

