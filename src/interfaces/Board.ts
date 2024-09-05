import Chip from './Chip'

interface Board {
    Chips: Chip[]
    nextPosition: number
    testTubePosition: number
    cherryBombValue: number
}

export default Board