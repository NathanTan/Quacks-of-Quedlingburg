// import React, { useContext } from 'react';
import SendMoveButton from './SendMoveButton';
import LogButton from './button';
import NewGameButton from './NewGameButton';
// import Board from './board';
import Board from './Board';
import { StoreContext } from './store/StoreConext';
import { useObserver } from 'mobx-react';
// import Board from './Board';

// interface Props {
//   message: string;
// }

// interface AppProps {
//   CounterStore?: any; // replace with the type of your store
// }


// const ParentComponent: React.FC = () => {
//     return <MyComponent message="Hello, World!" />;
// };


// const MyComponent: React.FC<Props> = ({ message }) => {
//   const store = useContext(StoreContext);

//   return useObserver(() => (
//     <div>
//       <link rel="manifest" href="/public/manifest.json" />
//       <h1>{message}</h1>
//       <button onClick={() => store.increment()}>Increment</button>
//       <button onClick={() => store.decrement()}>Decrement</button>
//       <p>Count: {store.count}</p>
//       <SendMoveButton /><LogButton /><NewGameButton />
//       <Board />
//     </div>
//   ));
// };


import React from 'react';
import { inject, observer } from 'mobx-react';

interface AppProps {
  CounterStore?: any; // replace with the type of your store
}

// @inject('CounterStore')
@observer
class App extends React.Component<AppProps> {

  message = "Hello, World!"
  render() {
    const { CounterStore } = this.props;

    return (
      <div>
        <link rel="manifest" href="/public/manifest.json" />
      <h1>{this.message}</h1>
      <button onClick={() => CounterStore.increment()}>Increment</button>
      <button onClick={() => CounterStore.decrement()}>Decrement</button>
      <p>Count: {CounterStore.count}</p>
      <SendMoveButton /><LogButton /><NewGameButton />
      <Board />
      </div>
    );
  }
}

export default App;