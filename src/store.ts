// store.ts
import { makeAutoObservable } from "mobx"
import QuacksState from "./interfaces/QuacksState"
import Player from "./interfaces/Player"

class Store {
  message = "Hello, Store!"
  state = {
    players: [],
    Round: 0,
    fortune: 0,
    winner: [],
    book: 0,
    bombLimit: 0,
    Awaiting: null,
    debug: false,
  } as QuacksState


  constructor() {
    makeAutoObservable(this);
  }

  updateMessage(newMessage: string) {
    this.message = newMessage;
  }

  updateState(newState: QuacksState) {
    this.state = newState
    console.log("Update state", this.state)
    console.log("Update new State", newState)
  }

  checkState() {
    this.message = JSON.stringify(this.state);
  }
}

export const myStore = new Store();