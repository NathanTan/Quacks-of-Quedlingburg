import React from 'react';

interface BoxProps {
  index: number;
  x: number;
  y: number;
  boxSize: number;
}

const Box: React.FC<BoxProps> = ({ index, x, y, boxSize }) => {
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
      {index}
    </div>
  );
};

export default Box;