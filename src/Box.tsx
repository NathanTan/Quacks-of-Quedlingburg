import { observer } from 'mobx-react';
import { myStore } from './store';
import React from 'react';

interface BoxProps {
  playerIndex: number
  index: number
  x: number
  y: number
  boxSize: number
}

const Box: React.FC<BoxProps> = observer(({ playerIndex, index, x, y, boxSize }) => {
  const boxStyle: React.CSSProperties = {
    position: 'absolute',
    width: `${boxSize}px`,
    height: `${boxSize}px`,
    backgroundColor: 'lightblue',
    color: 'black',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    transformOrigin: 'center',
    left: `${x - boxSize / 2}px`,
    top: `${y - boxSize / 2}px`
  };

  return (
    <div key={index} style={boxStyle}>
      {playerIndex},
      {/* {myStore.getPlayersChip(playerIndex, index)}, */}
      {/* {JSON.stringify(myStore.state.Players[playerIndex]?.Board?.Chips[index]?.color) ?? "x"}, */}
      {/* {JSON.stringify(myStore.state.Players[playerIndex]?.Board?.Chips[index]?.value) ?? ""}, */}
      {myStore.getPlayersChip(playerIndex, index)},
    </div>
  );

});




export default Box;