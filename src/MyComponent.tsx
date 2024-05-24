import React from 'react';
import SendMoveButton from './SendMoveButton';
import LogButton from './button';

interface Props {
  message: string;
}

const ParentComponent: React.FC = () => {
    return <MyComponent message="Hello, World!" />;
};

const MyComponent: React.FC<Props> = ({ message }) => {
  return (<div><link rel="manifest" href="/public/manifest.json" /><h1>{message}</h1><SendMoveButton /><LogButton /></div>);
};

export default MyComponent;
