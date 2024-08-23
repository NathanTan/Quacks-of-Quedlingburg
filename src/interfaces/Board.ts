import Chip from './Chip'

interface Board {
    chips: Chip[]
    nextPosition: number
    testTubePosition: number
    cherryBombValue: number
}

export default Board