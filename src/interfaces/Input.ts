import Chip from './Chip'

interface Input {
    Description: string
    options: string[]
    choice: number
    choice2: Chip[]
    player: number
    code: number
}

export default Input