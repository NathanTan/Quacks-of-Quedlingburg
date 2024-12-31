import React, { useState } from 'react';
import styled from 'styled-components';
import { myStore } from './store';
import { observer } from 'mobx-react';

const StyledButton = styled.button`
  background-color:rgb(36, 83, 252); /* Green */
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

const postMove = async (move: any) => {
  return await fetch('/continueDrawChip', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(move),
  });
}

const ContinueDrawChipButton: React.FC = observer(() => {

  const isButtonDisabled = false;
    myStore.buttons.DrawChip = true
    myStore.buttons.ContinueDrawChip = false

  const sendMove = async () => {
    console.log("Sending fortune move")
    const status = myStore.state.Status
    console.log("Status from state: " + status)

      const move = {"authToken": "game123", "gameId": "game123",
        "move": "1", "type": "ContinueDrawInput", "playerId": myStore.activePlayer}
      
      console.log(`Status is ${status}. Sending move ${JSON.stringify(move)}`)
      const response = await postMove(move)  

      if (!response.ok) {
        throw new Error('HTTP error ' + response.status);
      }

      myStore.update()
  };


  return (
    <StyledButton onClick={sendMove} disabled={isButtonDisabled}>
      Continue Draw Chip
    </StyledButton>
  );
});

export default ContinueDrawChipButton;