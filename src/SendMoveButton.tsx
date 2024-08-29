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


const SendMoveButton: React.FC = observer(() => {
  // const [move, setMove] = useState({ direction: 'up' }); // replace with your actual data

  const isButtonDisabled = false;
  const sendMove = async () => {
    console.log("Sending move")
    const status = myStore.state.Status

    if (status === "closed") {
      console.log("Status is closed, requesting to start game")
      try {
        const response = await fetch('/move', {
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
        for (let i = 1; i <= 3; i++) {
          await new Promise(resolve => setTimeout(resolve, 1000));
          console.log(`Waited ${i} second(s)`);
        }

        // Make a POST request to /getState
        const response2 = await fetch('/getState/game123', { method: 'POST' });

        // Parse the response as JSON
        const data2 = await response2.json();

        // Log the returned value
        console.log("Data has arrived")
        console.log(data2);

        myStore.updateState(data2);
        myStore.checkState();
      } catch (error) {
        console.error('Error:', error);
      }
    } else {
      try {
        const response = await fetch('/move', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ "move": "up" }),
        });

        if (!response.ok) {
          throw new Error('HTTP error ' + response.status);
        }

        const data = await response.json();
        console.log("data")
        console.log("Data from send move")
        console.log(data);
        myStore.updateMessage("Updated message " + response.status);
      } catch (error) {
        console.error('Error:', error);
      }
    }
  };

  console.log("Rendering SendMoveButton")

  return (
    <StyledButton onClick={sendMove} disabled={isButtonDisabled}>
      Send Move
    </StyledButton>
  );
});

export default SendMoveButton;