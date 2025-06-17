import React, { useState } from 'react';
import { observer } from 'mobx-react';
import { myStore } from './store';

const StateOptionsButton: React.FC = observer(() => {
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const fetchStateOptions = async () => {
    try {
      setIsLoading(true);
      
      // Use the store method to fetch state options
      await myStore.fetchStateOptions();
    } catch (error) {
      console.error("Failed to fetch state options:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const getButtonStyle = (): React.CSSProperties => {
    return {
      backgroundColor: 'green',
      color: 'white',
      padding: '10px',
      borderRadius: '5px',
      border: 'none',
      marginLeft: '10px'
    };
  };

  return (
    <div>
      <button 
        style={getButtonStyle()} 
        onClick={fetchStateOptions}
        disabled={isLoading}
      >
        {isLoading ? 'Loading...' : 'Get State Options'}      </button>
      
      {myStore.stateOptions.currentState && (
        <div style={{ marginTop: '10px', padding: '10px', backgroundColor: '#f5f5f5', borderRadius: '5px' }}>
          <h4>Current State: {myStore.stateOptions.currentState}</h4>
          <h4>Possible Transitions:</h4>
          <ul>
            {myStore.stateOptions.possibleTransitions.map((transition: string, index: number) => (
              <li key={index}>{transition}</li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
});

export default StateOptionsButton;
