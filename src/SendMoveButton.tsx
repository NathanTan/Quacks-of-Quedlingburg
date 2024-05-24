import React, { useState } from 'react';

const SendMoveButton: React.FC = () => {
  const [move, setMove] = useState({ direction: 'up' }); // replace with your actual data

  const sendMove = async () => {
    try {
      const response = await fetch('/move', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(move),
      });

      if (!response.ok) {
        throw new Error('HTTP error ' + response.status);
      }

      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <button onClick={sendMove}>
      Send Move
    </button>
  );
};

export default SendMoveButton;