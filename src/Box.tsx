import React, { CSSProperties } from 'react';
import { observer } from 'mobx-react-lite';
import { myStore } from './store';  

interface BoxProps {
  x: number;
  y: number;
  boxSize: number;
  playerIndex: number;
  index: number;
}

const Box: React.FC<BoxProps> = observer(({ x, y, boxSize, playerIndex, index }) => {
  const boxStyle: CSSProperties = {
    position: 'absolute',
    left: `${x - boxSize / 2}px`,
    top: `${y - boxSize / 2}px`
  };

  const chipText = myStore.getPlayersChip(playerIndex, index);

  return (
    <div key={index} style={boxStyle}>
      {chipText.includes("orange") ? (
        <img src="/static/public/imgs/pumpkin-icon.jpg" alt="Pumpkin" />
      ) : (
        chipText
      )}
    </div>
  );
});

export default Box;