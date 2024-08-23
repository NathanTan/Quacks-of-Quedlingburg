import React from 'react';
import { myStore } from './store';


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

    myStore.updateState(data);
    myStore.checkState();
  };

  render() {
    return <button onClick={this.handleClick}>New Game Button</button>;
  }
}

export default NewGameButton;