import React, { useEffect, useRef } from 'react';


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

const Board2: React.FC = () => {
  const boxes = Array.from({ length: 30 }, (_, i) => i);

  const boardStyle: React.CSSProperties = {
    position: 'relative',

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
    transformOrigin: 'center',
  };

  const canvasRef = useRef(null); // Step 1: Create a ref for the canvas


  const boxSize = 50; // size of the box
  const boxSpacing = 5; // space between boxes
  const spiralSpacing = boxSize + boxSpacing; // distance between box centers

  const positions: Position[] = boxes.map((box, index) => {
    const angle = index * Math.sqrt(spiralSpacing); // adjust this to change the tightness of the spiral
    const x = 250 + Math.cos(angle) * spiralSpacing * Math.sqrt(index) + boxSize / 2; // 250 is half the width/height of the board
    const y = 250 + Math.sin(angle) * spiralSpacing * Math.sqrt(index) + boxSize / 2; // adjust this to move the spiral
    return { x, y };
  });

  
  useEffect(() => { // Step 2: Use useEffect to draw on the canvas after the component mounts
    //@ts-ignore
    const canvas = canvasRef.current;
    //@ts-ignore
    const ctx = canvas.getContext('2d');
    if (ctx && canvas) {
    //@ts-ignore
      ctx.clearRect(0, 0, canvas.width, canvas.height); // Clear the canvas
      positions.forEach(position => {
        ctx.beginPath();
        ctx.arc(position.x, position.y, 5, 0, 2 * Math.PI); // Draw a circle for each position
        ctx.fill();
      });
    }
  }, [positions]); // Redraw when positions change


  return (
    <div style={boardStyle}>
      <p>hello</p>
      <canvas ref={canvasRef} width="500" height="500"></canvas> {/* Step 3: Attach the ref to the canvas */}
    </div>
  );
};

export default Board2;


