// store.ts
import { makeAutoObservable } from "mobx";

class Store {
  message = "Hello, Store!";

  constructor() {
    makeAutoObservable(this);
  }

  updateMessage(newMessage: string) {
    this.message = newMessage;
  }
}

export const myStore = new Store();