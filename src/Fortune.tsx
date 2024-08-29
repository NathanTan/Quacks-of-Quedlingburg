import React from 'react';
import { observer } from 'mobx-react';
import { myStore } from './store'; // adjust the path as necessary


const Fortune: React.FC = observer(({}) => {

    // TODO: Make sure the state doesn't disappear while trying to observe the state
    const fortune = myStore.state.Awaiting?.Description ? myStore.state.Fortune : "No Fortune";

    console.log("Awaiting")
    console.log(myStore.state.Awaiting)
  return (<div>
    <h2>Fortune: {fortune}</h2>
    </div>)
});


export default Fortune;
