import React from 'react';
import SendMoveButton from './SendMoveButton';
import LogButton from './UpdateStateButton';
import NewGameButton from './NewGameButton';
import Board from './Board';
import { observer } from 'mobx-react';
import { myStore } from './store'; // adjust the path as necessary
import Fortune from './Fortune';
import SendFortuneMoveButton from './SendFortuneMoveButton';
import DrawChipButton from './DrawChipButton';

// import Board from './Board';

interface Props {
  message: string;
}

const MyComponent: React.FC<Props> = observer(({message}) => {
  let playerNumber: number;
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
    <Fortune />

    <DrawChipButton />
    <SendFortuneMoveButton />
    <SendMoveButton />
    <input
        type="text"
        value={playerNumber}
        defaultValue="0"
        onChange={handleInputChange}
        placeholder="0"
      />
    <LogButton />
    <NewGameButton />
    {Array.from({ length: 4 }, (_, i) => <Board key={i} index={i} />)}
    </div>
    
  )
});


export default MyComponent;
