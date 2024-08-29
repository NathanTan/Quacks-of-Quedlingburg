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
    <h1>Props Message: {message}</h1>
    <SendMoveButton /><LogButton /><NewGameButton />
    {/* {myStore.players.length > 0 && <p>Number of players: {myStore.players.length}</p>} */}
    {/* {myStore.players.length > 0 && Array.from({ length: 4 }, (_, i) => <Board key={i} index={i} />)} */}
    {Array.from({ length: 4 }, (_, i) => <Board key={i} index={i} />)}
    {/* {myStore.state.players.map((player, index) => (
    <div>
      <Board key={index} index={index} />
    </div>
    ))} */}
    </div>)
});


// const ParentComponent: React.FC = () => {
//   return <MyComponent message="Hello, World!" />;
// };


export default MyComponent;
