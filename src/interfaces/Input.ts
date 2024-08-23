import Chip from './Chip'

interface Input {
    description: string
    options: string[]
    choice: number
    choice2: Chip[]
    player: number
    code: number
}

export default Input