import React from 'react';
import SendMoveButton from './SendMoveButton';
import LogButton from './UpdateStateButton';
import NewGameButton from './NewGameButton';
import StateOptionsButton from './StateOptionsButton';
import Board from './Board';
import { observer } from 'mobx-react';
import { myStore } from './store'; // adjust the path as necessary
import Fortune from './Fortune';
import SendFortuneMoveButton from './SendFortuneMoveButton';
import DrawChipButton from './DrawChipButton';
import ContinueDrawChipButton from './ContinueDrawChipButton';

// import Board from './Board';

interface Props {
  message: string;
}

const MyComponent: React.FC<Props> = observer(({message}) => {
  let playerNumber: number;
  // playerNumber = 0; // Default player number, can be changed by input
  // const playerName = myStore.state.Players[playerNumber].name;   

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = event.target.value;
    // setPlayerName(newValue);
    myStore.activePlayer = parseInt(newValue, 10); // Update the first player's name as an example
  };


  return (
  
  <div>
    <link rel="manifest" href="/public/manifest.json" />
    {/* <meta http-equiv="refresh" content="30"></meta> */}
    <h1>Game Status: {myStore.state.Status}</h1>
    
    {/* Add State Options display */}
    {myStore.stateOptions.currentState && (
      <div style={{ margin: '10px 0', padding: '10px', backgroundColor: '#e6f7ff', borderRadius: '5px', border: '1px solid #91d5ff' }}>
        <h3>Current Game State: {myStore.stateOptions.currentState}</h3>
        <h4>Possible Next States:</h4>
        <div style={{ display: 'flex', flexWrap: 'wrap', gap: '8px' }}>
          {myStore.stateOptions.possibleTransitions.map((transition: string, index: number) => (
            <div key={index} style={{ 
              padding: '5px 10px', 
              backgroundColor: '#1890ff', 
              color: 'white', 
              borderRadius: '4px',
              fontSize: '14px'
            }}>
              {transition}
            </div>
          ))}
        </div>
      </div>
    )}
    
    <Fortune />
    <h2>Active Player: {JSON.stringify(myStore.state)}</h2>

    <ContinueDrawChipButton />
    <DrawChipButton />
    <SendFortuneMoveButton />
    <SendMoveButton />
    <input
        type="Player Number"
        value={playerNumber}
        defaultValue="0"
        onChange={handleInputChange}
        placeholder="0"      />
    <LogButton />
    <NewGameButton />
    <StateOptionsButton />
    {Array.from({ length: myStore.state.Players.length }, (_, i) => <Board key={i} index={i} />)}
    </div>
    
  )
});


export default MyComponent;
