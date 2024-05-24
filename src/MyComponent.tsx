import React from 'react';
import SendMoveButton from './sendMoveButton';

interface Props {
  message: string;
}

const ParentComponent: React.FC = () => {
    return <MyComponent message="Hello, World!" />;
};

const MyComponent: React.FC<Props> = ({ message }) => {
  return (<div><h1>{message}</h1><SendMoveButton /></div>);
};

export default MyComponent;
