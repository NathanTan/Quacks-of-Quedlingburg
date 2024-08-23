import Bag from './Bag';
import Board from './Board';

interface Player {
  name: string;
  bag: Bag;
  board: Board;
  isDoneDrawing: boolean;
  hasCompletedTheFortune: boolean;
  hasSpentRubies: boolean;
  rubyCount: number;
  ratToken: number;
  dropplet: number;
  testTubeDropplet: number;
  flask: boolean;
  explosionLimit: number;
  score: number;
  chooseVictoryPoints: boolean;
  chooseBuying: boolean;
  buyingPower: number;
}

export default Player