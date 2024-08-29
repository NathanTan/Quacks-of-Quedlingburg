import React from 'react';
import SendMoveButton from './SendMoveButton';
import LogButton from './button';
import NewGameButton from './NewGameButton';
import Board from './Board';
import { observer } from 'mobx-react';
import { myStore } from './store'; // adjust the path as necessary

// import Board from './Board';

interface Props {
  message: string;
}


const MyComponent: React.FC<Props> = observer(({message}) => {
  return (<div>
    <link rel="manifest" href="/public/manifest.json" />
    <meta http-equiv="refresh" content="30"></meta>
    <h1>Store Message: {myStore.message}</h1>
    <h1>Game Status: {myStore.state.Status}</h1>
    <SendMoveButton /><LogButton /><NewGameButton />
    {Array.from({ length: 4 }, (_, i) => <Board key={i} index={i} />)}
    </div>)
});


export default MyComponent;
