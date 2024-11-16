// Modal.tsx
import React, { useEffect } from 'react';
import './modal.css'; // Optional: CSS file for styling the modal

interface ModalProps {
  show: boolean;
  onClose: () => void;
  children: React.ReactNode;
}

const Modal: React.FC<ModalProps> = ({ show, onClose, children }) => {
  useEffect(() => {
    // Function to handle keydown event
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        onClose(); // Call the onClose function when Escape is pressed
      }
    };

    // Add event listener to document when modal is shown
    if (show) {
      document.addEventListener('keydown', handleKeyDown);
    }

    // Cleanup event listener on component unmount or when modal is hidden
    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [show, onClose]);

  if (!show) return null;

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <button className="modal-close" onClick={onClose}>
          &times;
        </button>
        {children}
      </div>
    </div>
  );
};

export default Modal;