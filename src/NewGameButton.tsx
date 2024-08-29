import React from 'react';
import { myStore } from './store';
import { observer } from 'mobx-react';

class NewGameButton extends React.Component {
  handleClick = async () => {
    // Make a POST request to /requestState
    await fetch('/requestState', { method: 'POST' });

    // Wait for 3 seconds
    // Loop 3 times, waiting 1 second each time and logging
    for (let i = 1; i <= 3; i++) {
      await new Promise(resolve => setTimeout(resolve, 1000));
      console.log(`Waited ${i} second(s)`);
    }

    // Make a POST request to /getState
    const response = await fetch('/getState/game123', { method: 'POST' });

    // Parse the response as JSON
    const data = await response.json();

    // Log the returned value
    console.log("Data has arrived")
    console.log(data);

    myStore.updateState(data);
    myStore.checkState();
  };

  render() {
    console.log("zzzz state", myStore.state)
    const style: React.CSSProperties = {
      backgroundColor: 'blue',
      color: 'white',
      padding: '10px',
      borderRadius: '5px',
      border: 'none',
      visibility: (myStore.state.Players.length === 0) ? 'visible' : 'hidden',
    }
    return <button style={style} onClick={this.handleClick}>New Game Button</button>;
  }
}

export default observer(NewGameButton);