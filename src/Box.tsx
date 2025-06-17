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
  const chipColor = myStore.getPlayerChipColor(playerIndex, index);
  const chipValue = myStore.getPlayerChipNumber(playerIndex, index);

  // Function to get the image source based on chip color
  const getChipImage = (color: string) => {
    
    switch (color) {
      case 'orange':
        return { src: "/static/public/imgs/pumpkin-icon.jpg", alt: "Orange Chip" };
      case 'blue':
        return { src: "/static/public/imgs/blue-chip.png", alt: "Blue Chip" };
      case 'black':
        return { src: "/static/public/imgs/black-chip.jpg", alt: "Black Chip" };
      case 'green':
        return { src: "/static/public/imgs/green-chip.jpg", alt: "Green Chip" };
      case 'red':
        return { src: "/static/public/imgs/red-chip.png", alt: "Red Chip" };
      case 'white':
        return { src: "/static/public/imgs/white-chip.png", alt: "White Chip" };
      case 'yellow':
        return { src: "/static/public/imgs/yellow-chip.png", alt: "Yellow Chip" };
      case 'purple':
        return { src: "/static/public/imgs/purple-chip.png", alt: "Purple Chip" };
      default:
        return null;
    }
  };
  const chipOverlayStyle: CSSProperties = {
    position: 'relative',
    display: 'inline-block',
    width: `${boxSize}px`,
    height: `${boxSize}px`,
  };

  const textOverlayStyle: CSSProperties = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    color: 'black',
    fontWeight: 'bold',
    fontSize: `${boxSize * 0.4}px`,
    textShadow: '1px 1px 2px white',
    zIndex: 10,
  };

  const imageStyle: CSSProperties = {
    width: '100%',
    height: '100%',
    objectFit: 'cover',
    borderRadius: '50%',
  };

  const renderSwitch = (chip: string) => {
    const color = myStore.getPlayerChipColor(playerIndex, index);

    switch(color) {
      case 'orange':
        return (
          <div style={chipOverlayStyle}>
            <img src="/static/public/imgs/pumpkin-icon.jpg" alt="Orange Chip" style={imageStyle} />
            <span style={textOverlayStyle}>{chipValue}</span>
          </div>
        );
      case 'white':
        return (
          <div style={chipOverlayStyle}>
            <img src="/static/public/imgs/white-chip.jpg" alt="White Chip" style={imageStyle} />
            <span style={textOverlayStyle}>{chipValue}</span>
          </div>
        );
      case 'green':
        return (
          <div style={chipOverlayStyle}>
            <img src="/static/public/imgs/green-chip.jpg" alt="Green Chip" style={imageStyle} />
            <span style={textOverlayStyle}>{chipValue}</span>
          </div>
        );
      default:
        return chip;
    }
  }

  const chipImage = getChipImage(chipText);
  return (
    <div key={index} style={boxStyle}>
      {renderSwitch(chipText)}
    </div>
  );
});

export default Box;