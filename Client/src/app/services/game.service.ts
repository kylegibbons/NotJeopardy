import { Injectable } from '@angular/core';
import { Game, ClueSelect, ClueDetermination } from '../models/game';
import { WebsocketService } from './websocket.service';
import { Observable, Subject, Subscription, ReplaySubject } from 'rxjs';
import { map, filter, retryWhen, delay } from 'rxjs/operators';
import { normalClosureMessage } from 'rxjs-websockets';
import { Message } from '../models/message.model';

@Injectable()
export class GameService {

  public game$: ReplaySubject<Game> = new ReplaySubject(1);;
  //private game: Subject<Game>;
  private gameData: Game;


  constructor(private socket: WebsocketService) { 
    console.log('Starting Game service');

    //this.game = new Subject<Game>();
    //this.game$ = 

    this.socket.connect();

    //JOIN GAME
    this.socket.send({
      messageType: "JoinGame",
      timestamp: null,
      gameId: "a267dd0b-40cb-4178-ad8c-58d5efa1ff29",
      payload: null
    });

    const messagesSubscription: Subscription = this.socket.messages$.subscribe(
      (message: string) => {
        console.log('received message')
        const thisMessage: Message = JSON.parse(message);

        switch (thisMessage.messageType) {
          case "GameState":
            console.log("Message: GameState")

            this.gameData = thisMessage.payload;

            this.game$.next(this.gameData);
            break;
          /* case "SelectClue":
            this.gameData.activeCategory = this.gameData.rounds[this.gameData.round].categories[thisMessage.payload.categoryNumber];
            this.gameData.activeClue = this.gameData.rounds[this.gameData.round].categories[thisMessage.payload.categoryNumber].clues[thisMessage.payload.clueNumber];
            this.game.next(this.gameData);
            break; */
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
  }

  public SelectClue(categoryNumber: number, clueNumber: number) {
    this.socket.send(
      {
        messageType: "SelectClue",
        gameId: this.gameData.id,
        timestamp: null,
        payload: {
          round: this.gameData.round,
          categoryNumber: categoryNumber,
          clueNumber: clueNumber,
        }
      }
    )
  }

  public EnableBuzzers() {
    this.socket.send(
      {
        messageType: "EnableBuzzers",
        gameId: this.gameData.id,
        timestamp: null,
        payload:  null,
      }
    )
  }

  public ResetBuzzers() {
    this.socket.send(
      {
        messageType: "ResetBuzzers",
        gameId: this.gameData.id,
        timestamp: null,
        payload:  null,
      }
    )
  }

  public SelectContestant(contestant: number) {
    this.socket.send(
      {
        messageType: "SelectContestant",
        gameId: this.gameData.id,
        timestamp: null,
        payload:  {
          contestant: contestant,
        },
      }
    )
  }

  public DetermineClue(correct: boolean, dailyDoubleAmount: number) {
    this.socket.send(
      {
        messageType: "ClueDetermination",
        gameId: this.gameData.id,
        timestamp: null,
        payload:  {
          correct: correct,
          dailyDoubleAmount: dailyDoubleAmount,
        }
      }
    )
  }

}
