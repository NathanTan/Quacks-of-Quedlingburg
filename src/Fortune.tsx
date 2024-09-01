import React from 'react';
import { observer } from 'mobx-react';
import { myStore } from './store'; // adjust the path as necessary


const Fortune: React.FC = observer(({}) => {

  console.log("Round: " + myStore.state.Round);
    const fortune = myStore.turnFortune.get(myStore.state.Round) ?? "No Fortune";

  return (<div>
    <h2>Fortune: {fortune}</h2>
    </div>)
});


export default Fortune;
