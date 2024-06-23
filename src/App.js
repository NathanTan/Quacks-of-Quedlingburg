// // import React, { useContext } from 'react';
// import SendMoveButton from './SendMoveButton';
// import LogButton from './button';
// import NewGameButton from './NewGameButton';
// // import Board from './board';
// import Board from './Board';
// import { StoreContext } from './store/StoreConext';
// import { useObserver } from 'mobx-react';
// // import Board from './Board';



// import React from 'react';
// import { inject, observer } from 'mobx-react';
// import Board2 from './Board2';

// interface AppProps {
//   CounterStore?: AppProps.CounterStore; // replace with the type of your store
// }

// // @inject('CounterStore')
// class App extends React.Component<AppProps> {

//   message = "Hello, World!"
//   render() {
//     const { CounterStore } = this.props;

//     return (
//       <div>
//         <link rel="manifest" href="/public/manifest.json" />
//       <h1>{this.message}</h1>
//       <button onClick={() => CounterStore.increment()}>Increment</button>
//       <button onClick={() => CounterStore.decrement()}>Decrement</button>
//       <p>Count: {CounterStore.count}</p>
//       <SendMoveButton /><LogButton /><NewGameButton />
//       <Board />
//       <Board2 />
//       </div>
//     );
//   }
// }

// export default App;