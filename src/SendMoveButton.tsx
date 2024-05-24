import React, { useState } from 'react';
import styled from 'styled-components';

const StyledButton = styled.button`
  background-color: #4CAF50; /* Green */
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
  transition-duration: 0.4s;
  border-radius: 12px;

  &:hover {
    background-color: #45a049;
  }
`;

const SendMoveButton: React.FC = () => {
  // const [move, setMove] = useState({ direction: 'up' }); // replace with your actual data

  const isButtonDisabled = false;
  const sendMove = async () => {
    console.log("Sending move")

    try {
      const response = await fetch('/move', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({"move": "up"}),
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

  console.log("Rendering SendMoveButton")

  return (
    <StyledButton onClick={sendMove} disabled={isButtonDisabled}>
      Send Move
    </StyledButton>
  );
};

export default SendMoveButton;