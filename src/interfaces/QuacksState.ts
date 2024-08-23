import Player from './Player';
import Input from './Input';

interface GameState {
  players: Player[];
  Round: number;
  fortune: number;
  winner: number[];
  book: number;
  bombLimit: number;
  Awaiting: Input | null;
  debug: boolean;
//   FSM: FSM;
//   fortuneDeck: Fortune[];
//   Stats: Stats;
}

export default GameState