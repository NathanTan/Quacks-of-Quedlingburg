import Bag from './Bag';
import Board from './Board';

interface Player {
  Name: string;
  bag: Bag;
  Board: Board;
  isDoneDrawing: boolean;
  hasCompletedTheFortune: boolean;
  hasSpentRubies: boolean;
  rubyCount: number;
  ratToken: number;
  dropplet: number;
  testTubeDropplet: number;
  flask: boolean;
  explosionLimit: number;
  NextPosition: number
  score: number;
  CherryBombValue: number
  chooseVictoryPoints: boolean;
  chooseBuying: boolean;
  buyingPower: number;
}

export default Player