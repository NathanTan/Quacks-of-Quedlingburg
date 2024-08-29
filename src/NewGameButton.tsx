import React from 'react';
import { myStore } from './store';


class NewGameButton extends React.Component {
  handleClick = async () => {
    // Make a POST request to /requestState
    await fetch('/requestState', { method: 'POST' });

    // Wait for 5 seconds
    // await new Promise(resolve => setTimeout(resolve, 5000));
    // Loop 3 times, waiting 1 second each time and logging
    for (let i = 1; i <= 3; i++) {
      await new Promise(resolve => setTimeout(resolve, 1000));
      console.log(`Waited ${i} second(s)`);
    }

    // Make a POST request to /getState
    const response = await fetch('/getState', { method: 'POST' });

    // Parse the response as JSON
    const data = await response.json();

    // Log the returned value
    console.log("Data has arrived")
    console.log(data);

    myStore.updateState(data);
    myStore.checkState();
  };

  render() {
    return <button onClick={this.handleClick}>New Game Button</button>;
  }
}

export default NewGameButton;