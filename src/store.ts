// store.ts
import { makeAutoObservable } from "mobx"
import QuacksState from "./interfaces/QuacksState"
import Player from "./interfaces/Player"

class Store {
  message = "Hello, Store!"
  turnFortune = new Map<number, string>()
  activePlayer = 0
  state = {
    Players: [],
    Round: 0,
    Fortune: 0,
    winner: [],
    book: 0,
    bombLimit: 0,
    Awaiting: null,
    FrontEndAwaiting: null,
    debug: false,
    Status: "New Game"
  } as QuacksState

  buttons = {
    NewGame: true,
    DrawChip: false,
    ContinueDrawChip: false,
    SendMove: false,
    SendFortuneMove: false,
  }


  constructor() {
    makeAutoObservable(this);
  }

  updateMessage(newMessage: string) {
    this.message = newMessage;
  }

  updateState(newState: QuacksState) {
    console.log("1st state", this.state)
    this.state = newState
    console.log("Update new State", newState)
  }

  checkState() {
    this.message = JSON.stringify(this.state);
  }

  updateFortune(round: number, fortuneDescription: string) {

    if (this.state.debug)
        console.log("Updating Fortune for Round", round, "State Round:", this.state.Round, "Round Does Not Yet Have the Fortune:", !this.turnFortune.has(round))

    // Only set the fortune if it hasn't been set yet
    if ((this.state.Round === round && !this.turnFortune.has(round)) ||
        this.turnFortune.get(round) === "No Fortune") {  
      this.turnFortune.set(round, fortuneDescription)
    }
  }

  async update() {
      // Make a POST request to /getState
      const response = await fetch('/getState/game123', { method: 'POST' });

      // Parse the response as JSON
      // const data = await response.json() as QuacksState;
      const data = await response.json()

      // Log the returned value
      console.log("Data has arrived")
      console.log(data);

      const fortuneText = data.Input?.Description ?? "No Fortune";
      console.log("Fortune Text", fortuneText)
      this.updateFortune(data.Round, fortuneText)

      this.state.Status = data.Status
      
    }

  getPlayer(index: number): Player  {  
    if (this.message === "Hello, Store!") {
      return {} as Player
    }

    if (this.state == null) {
      return {} as Player 
    }

    return this.state.Players[index] ?? {} as Player;
  }
  
  getPlayerName(index: number): string {  
    if (this.message === "Hello, Store!") {
      return ""
    }
    return this.state.Players[index]?.Name ?? "";
  }
}

export const myStore = new Store();