// store.ts
import { makeAutoObservable } from "mobx"
import QuacksState from "./interfaces/QuacksState"
import Player from "./interfaces/Player"

class Store {
  message = "Hello, Store!"
  state = {
    Players: [],
    Round: 0,
    fortune: 0,
    winner: [],
    book: 0,
    bombLimit: 0,
    Awaiting: null,
    debug: false,
    Status: "New Game"
  } as QuacksState


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

  getPlayer(index: number): Player  {  
    if (this.message === "Hello, Store!") {
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