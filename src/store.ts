// store.ts
import { makeAutoObservable } from "mobx";
import QuacksState from "./interfaces/QuacksState";

class Store {
  message = "Hello, Store!";
  state = {
  } as QuacksState;

  constructor() {
    makeAutoObservable(this);
  }

  updateMessage(newMessage: string) {
    this.message = newMessage;
  }

  updateState(newState: QuacksState) {
    this.state = newState;
  }

  checkState() {
    this.message = JSON.stringify(this.state);
  }
}

export const myStore = new Store();