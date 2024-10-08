import Player from './Player';
import Input from './Input';

interface GameState {
  Players: Player[];
  Round: number;
  Fortune: number;
  winner: number[];
  book: number;
  bombLimit: number;
  Awaiting: Input | null;
  FrontEndAwaiting: Input | null;
  debug: boolean;
//   FSM: FSM;
//   fortuneDeck: Fortune[];
//   Stats: Stats;
  Status: string
}

export default GameState