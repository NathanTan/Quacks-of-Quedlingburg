import React from 'react';
import { myStore } from './store';
import { observer } from 'mobx-react';

interface NewGameButtonState {
  isVisible: boolean
}


class NewGameButton extends React.Component<{}, NewGameButtonState> {
  constructor(props: {}) {
    super(props)
    this.state = {
      isVisible: true
    }
  }

  handleClick = async () => {
    if (myStore.state.Players.length == 0) {
      this.setState({ isVisible: false });

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

    }
  };

  getButtonStyle = (): React.CSSProperties => {
    return {
      backgroundColor: 'blue',
      color: 'white',
      padding: '10px',
      borderRadius: '5px',
      border: 'none',
      visibility: this.state.isVisible ? 'visible' : 'hidden',
    };
  };

  render() {
    return <button style={this.getButtonStyle()} onClick={this.handleClick}>New Game Button</button>;
  }
}

export default observer(NewGameButton);