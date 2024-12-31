import React, { useState } from 'react';
import styled from 'styled-components';
import { myStore } from './store';
import { observer } from 'mobx-react';

const StyledButton = styled.button`
  background-color:rgb(118, 175, 120); /* Green */
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
  return await fetch('/move', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(move),
  });
}

const SendFortuneMoveButton: React.FC = observer(() => {
  // const [move, setMove] = useState({ direction: 'up' }); // replace with your actual data

  const isButtonDisabled = false;
  myStore.buttons.DrawChip = true
  const sendMove = async () => {
    console.log("Sending fortune move")
    const status = myStore.state.Status
    console.log("Status from state: " + status)

    if (status === "fortune_input") {
      const move = {"authToken": "game123", "gameId": "game123",
        "move": "1", "type": "Input", "playerId": myStore.activePlayer}
      
      console.log(`Status is ${status}. Sending move ${JSON.stringify(move)}`)
      const response = await postMove(move)  

      if (!response.ok) {
        throw new Error('HTTP error ' + response.status);
      }
    } else {
      console.log(`Status '${status}'is not fortune, cannot send fortune move`)
    }

    myStore.buttons.DrawChip = true
    myStore.update()
  };


  return (
    <StyledButton onClick={sendMove} disabled={isButtonDisabled}>
      Send Fortune Move
    </StyledButton>
  );
});

export default SendFortuneMoveButton;