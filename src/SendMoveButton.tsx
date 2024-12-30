import React, { useState } from 'react';
import styled from 'styled-components';
import { myStore } from './store';
import { observer } from 'mobx-react';

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

const postMove = async (move: any) => {
  return await fetch('/move', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(move),
  });
}

const SendMoveButton: React.FC = observer(() => {
  // const [move, setMove] = useState({ direction: 'up' }); // replace with your actual data

  const isButtonDisabled = false;
  const sendMove = async () => {
    console.log("Sending move")
    const status = myStore.state.Status
    console.log("Status from state: " + status)

    if (status === "Open") {
      const move = {"authToken": "game123", "gameId": "game123",
        "move": "1", "type": "Input", "playerId": myStore.activePlayer} 
      
      console.log(`Status is ${status}. Sending move ${JSON.stringify(move)}`)
      const response = await postMove(move)  

      if (!response.ok) {
        throw new Error('HTTP error ' + response.status);
      }

    } else if (status === "closed") {
      console.log("Status is closed, requesting to start game")
      try {
        const response = await fetch('/startGame', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ "move": "StartGame" }),
        });

        if (!response.ok) {
          throw new Error('HTTP error ' + response.status);
        }

        console.log("Sent request to start game")


        // Wait for 3 seconds
        // Loop 3 times, waiting 1 second each time and logging
      //   for (let i = 1; i <= 3; i++) {
      //     await new Promise(resolve => setTimeout(resolve, 1000));
      //     console.log(`Waited ${i} second(s)`);
      //   }

      //   // Make a POST request to /getState
      //   const response2 = await fetch('/getState/game123', { method: 'POST' });

      //   // Parse the response as JSON
      //   const data2 = await response2.json();

      //   // Log the returned value
      //   console.log("Data has arrived")
      //   console.log(data2);

      //   myStore.updateState(data2);
      //   myStore.checkState();
      } catch (error) {
        console.error('Error:', error);
      }
    } else if (status === "New Game") {
      try {
        const startMove = {"authToken": "game123", "gameId": "game123",
          "move": "1", "type": "StartGame"}
        
        console.log(`Status is ${status}. Sending move ${startMove}`)
        const response2 = await postMove(startMove)  
  
        if (!response2.ok) {
          throw new Error('HTTP error ' + response2.status);
        }

        const data = await response2.json();
        console.log("data")
        console.log("Data from send move")
        console.log(data);
        myStore.updateMessage("Updated message " + response2.status);
      } catch (error) {
        console.error('Error:', error);
      }
    }
  };


  return (
    <StyledButton onClick={sendMove} disabled={isButtonDisabled}>
      Initalize Game
    </StyledButton>
  );
});

export default SendMoveButton;