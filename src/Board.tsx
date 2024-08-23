import React from 'react';
import Box from './Box';
import { myStore } from './store';

interface Position {
  x: number;
  y: number;
}

interface LineProps {
  x1: number;
  y1: number;
  x2: number;
  y2: number;
}

const Line: React.FC<LineProps> = ({ x1, y1, x2, y2 }) => {
  const length = Math.sqrt(Math.pow(x2 - x1, 2) + Math.pow(y2 - y1, 2));
  const angle = Math.atan2(y2 - y1, x2 - x1) * 180 / Math.PI;
  const positionStyle = { left: `${x1}px`, top: `${y1}px` };

  return (
    <div>
    <div style={{
      position: 'absolute',
      transform: `rotate(${angle}deg)`,
      width: `${length}px`,
      height: '1px',
      backgroundColor: 'black',
      ...positionStyle
    }} >

    </div>
    <div style={{position: 'absolute', width: '55px', height:'55px', color: 'red', zIndex: 111}}></div>
    </div>
  );
};

interface BoardProps {
  index: number
}

const Board: React.FC<BoardProps> = ({index}) => {
  const boxes = Array.from({ length: 30 }, (_, i) => i);

  const boardStyle: React.CSSProperties = {
    position: 'relative',
    width: '600px',
    height: '600px',
    borderRadius: '50%',
    backgroundColor: 'lightgreen',
    visibility: (myStore.state) ? 'visible' : 'hidden',
  };

  const boxStyle: React.CSSProperties = {
    position: 'absolute',
    width: '50px',
    height: '50px',
    backgroundColor: 'lightblue',
    color: 'black',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    transformOrigin: 'center',
  };

  const boxSize = 50; // size of the box
  const boxSpacing = 5; // space between boxes
  const spiralSpacing = boxSize + boxSpacing; // distance between box centers

  const positions: Position[] = boxes.map((box, index) => {
    const angle = index * Math.sqrt(spiralSpacing); // adjust this to change the tightness of the spiral
    const x = 250 + Math.cos(angle) * spiralSpacing * Math.sqrt(index) + boxSize / 2; // 250 is half the width/height of the board
    const y = 250 + Math.sin(angle) * spiralSpacing * Math.sqrt(index) + boxSize / 2; // adjust this to move the spiral
    return { x, y };
  });

  return (
    <div style={boardStyle}>
      {/* <p>Player: {myStore.state.players[index]?.name}</p> */}
      {<p>Board for player: {myStore.state.players && JSON.stringify(myStore.state.players)}</p>}
      {positions.map(({ x, y }, index2) => (
        <Box key={index2} index={index2} x={x} y={y} boxSize={boxSize} />
      ))}
      {/* {positions.map(({ x, y }, index) => {
        if (index === 0) return null;
        const { x: x1, y: y1 } = positions[index - 1];
        return <Line key={index} x1={x1} y1={y1} x2={x} y2={y} />;
      })} */}
    </div>
  );
};

export default Board;


