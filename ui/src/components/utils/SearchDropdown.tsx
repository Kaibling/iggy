// DropdownWithSearch.tsx
import React, { useState } from 'react';
import './SearchDropdown.css';
import { Option } from './types';




interface SearchDropdownProps {
  options: Option[] | null;
  onSelect: (selectedOption: Option) => void;
  placeholder: string,
}

const SearchDropdown: React.FC<SearchDropdownProps> = ({ options, onSelect, placeholder = "Search..." }) => {
  const [searchQuery, setSearchQuery] = useState('');
  const [isOpen, setIsOpen] = useState(false);


  const filteredOptions = options
    ? options.filter(option =>
      option.name.toLowerCase().includes(searchQuery.toLowerCase())
    )
    : [];

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchQuery(e.target.value);
  };

  const handleOptionClick = (option: {id:string,name:string;}) => {
    onSelect(option);
    setSearchQuery(option.name);
    setIsOpen(false);
  };

  const toggleDropdown = () => {
    setIsOpen(!isOpen);
  };

  return (
    <div className="dropdown-with-search">
      <input
        type="text"
        value={searchQuery}
        onChange={handleInputChange}
        onClick={toggleDropdown}
        placeholder={placeholder}
      />
      {isOpen && (
        <ul className="dropdown-list">
          {filteredOptions.length > 0 ? (
            filteredOptions.map((option, index) => (
              <li key={option.id} onClick={() => handleOptionClick(option)}>
                {option.name}
              </li>
            ))
          ) : (
            <li className="no-options">No options found</li>
          )}
        </ul>
      )}
    </div>
  );
};

export default SearchDropdown;