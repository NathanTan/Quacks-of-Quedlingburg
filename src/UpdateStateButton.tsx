import React from 'react';
import { myStore } from './store';

const LogButton: React.FC = () => {
  const handleClick = () => {
    myStore.update()
    console.log("Updated store")
  };

  return (
    <button onClick={handleClick}>
      Update State
    </button>
  );
};

export default LogButton;