import React from 'react';
import { myStore } from './store';

const LogButton: React.FC = () => {
  const handleClick = async () => {
    await myStore.update()
          // myStore.updateState(data);
    
    console.log("Updated store")
  };

  return (
    <button onClick={handleClick}>
      Update State
    </button>
  );
};

export default LogButton;