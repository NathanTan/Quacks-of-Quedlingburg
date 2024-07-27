import React from 'react';
import SendMoveButton from './SendMoveButton';
import LogButton from './button';
import NewGameButton from './NewGameButton';
import Board from './board';
import { observer } from 'mobx-react';
import { myStore } from './store'; // adjust the path as necessary

// import Board from './Board';

interface Props {
  message: string;
}


const MyComponent: React.FC<Props> = observer(({message}) => {
  return (<div>
    <link rel="manifest" href="/public/manifest.json" />
    <h1>Store Message: {myStore.message}</h1>
    <h1>Props Message: {message}</h1>
    <SendMoveButton /><LogButton /><NewGameButton />
    <Board />
    </div>)
});


// const ParentComponent: React.FC = () => {
//   return <MyComponent message="Hello, World!" />;
// };


export default MyComponent;
