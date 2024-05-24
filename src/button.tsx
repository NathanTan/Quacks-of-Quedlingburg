import React from 'react';

const LogButton: React.FC = () => {
  const handleClick = () => {
    console.log('Button was clicked');
  };

  return (
    <button onClick={handleClick}>
      Click me
    </button>
  );
};

export default LogButton;