import { Injectable } from '@angular/core';
import { Game } from '../models/game';
import { WebsocketService } from './websocket.service';
import { Observable, Subject, Subscription } from 'rxjs';
import { map, filter, retryWhen, delay } from 'rxjs/operators';
import { normalClosureMessage } from 'rxjs-websockets';
import { Message } from '../models/message.model';

@Injectable()
export class GameService {

  public game$: Observable<Game>;
  private game: Subject<Game>;
  private gameData: Game;

  constructor(private socket: WebsocketService) { 
    console.log('Starting Game service');

    this.game = new Subject<Game>();
    this.game$ = this.game.asObservable();

    this.socket.connect();

    const messagesSubscription: Subscription = this.socket.messages$.subscribe(
      (message: string) => {
        console.log('received message:', message)
        const thisMessage: Message = JSON.parse(message);

        switch (thisMessage.messageType) {
          case "GameState":
            console.log("Message: GameState")

            this.gameData = JSON.parse(thisMessage.payload)

            this.game.next(this.gameData);
            break;
          case "SelectClue":
            console.log("Message: SelectClue")
            break;
        }
        
      },
      (error: Error) => {
        const { message } = error
        if (message === normalClosureMessage) {
          console.log('server closed the websocket connection normally')
        } else {
          console.log('socket was disconnected due to error:', message)
        }
      },
      () => {
        // The clean termination only happens in response to the last
        // subscription to the observable being unsubscribed, any
        // other closure is considered an error.
        console.log('the connection was closed in response to the user')
      },
    )


/*
    this.game = Observable.create((obs) => {

      // this.socketSubscription =
      this.socket.messages.pipe(
        retryWhen(errors => errors.pipe(delay(1000)))).pipe(
        /*filter((message: string) => {
          // Print message to the console
          // console.log('Unit:');
          // console.log(message);
          const thisMessage: Message = JSON.parse(message);
          if (thisMessage.MessageType === 'units') {
            return true;
          }
          return false;
        })*/
      /*)
      .subscribe((message: string) => {
        //const thisMessage: Message = JSON.parse(message);
        //obs.next(thisMessage.Payload);
        console.log(message)
        obs.next(message);
        });
      });*/

  }

  public SelectClue(categoryNumber:string, clueNumber: string) {
    this.socket.send(
      {
        messageType: "SelectClue",
        timestamp: null,
        payload: {
          gameID: this.gameData.id,
          round: this.gameData.round,
          categoryNumber: categoryNumber,
          clueNumber: clueNumber,
        }
      }
    )
  }

}
