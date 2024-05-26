// store.ts
import { createContext } from 'react';
import CounterStore from './CounterStore';

export const StoreContext = createContext(CounterStore);