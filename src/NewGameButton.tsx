import React from 'react';
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

class NewGameButton extends React.Component {

  handleClick = async () => {
    // Make a POST request to /requestState
    await fetch('/requestState', { method: 'POST' });

    // Wait for 5 seconds
    await new Promise(resolve => setTimeout(resolve, 5000));

    // Make a POST request to /getState
    const response = await fetch('/getState', { method: 'POST' });

    // Parse the response as JSON
    const data = await response.json();

    // Log the returned value
    console.log(data);
  };

  isButtonDisabled = false;

  render() {
    return (
      <StyledButton onClick={this.handleClick} disabled={this.isButtonDisabled}>
        Create New Game
      </StyledButton>
    );
  }
}

export default NewGameButton;