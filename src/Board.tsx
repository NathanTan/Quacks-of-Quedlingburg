
import React from 'react';

const Board = () => {
  const boxes = Array.from({ length: 30 }, (_, i) => i);

  const boardStyle: React.CSSProperties = {
    position: 'relative',
    width: '500px',
    height: '500px',
    borderRadius: '50%',
    backgroundColor: 'lightgreen',
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
    left: '50%',
    top: '50%',
    transformOrigin: 'top left',
  };

  return (
    <div style={boardStyle}>
      {boxes.map((box, index) => (
        <div key={index} style={{ ...boxStyle, transform: `rotate(${index * 12}deg) translateY(${index * 10}px)` }}>
        {/* <div key={index} style={{ ...boxStyle, transform: `rotate(${index * 12}deg)` }}> */}
          {box}
        </div>
      ))}
    </div>
  );
};

export default Board;