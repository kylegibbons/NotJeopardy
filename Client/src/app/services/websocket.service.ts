import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

import { QueueingSubject } from 'queueing-subject'
import { Observable, Subscription, SubscriptionLike } from 'rxjs'
import { share, switchMap } from 'rxjs/operators'
import makeWebSocketObservable, {
  GetWebSocketResponses,
  normalClosureMessage,
} from 'rxjs-websockets'
import { Message } from '../models/message.model';

@Injectable()
export class WebsocketService {
  //public messages: Observable<String>;
  //public ConnectionStatus: Observable<Number>;

  // this subject queues as necessary to ensure every message is delivered
  private input$ = new QueueingSubject<string>()

  // create the websocket observable, does *not* open the websocket connection
  private socket$ = makeWebSocketObservable<string>('ws://' + environment.base_url + '/ws')

  // the observable produces a value once the websocket has been opened
  public messages$: Observable<string>;

  public messagesSubscription: Subscription;

  public constructor() { }

  public closeWebsocket() {
    // this also caused the websocket connection to be closed
    this.messagesSubscription.unsubscribe()
  }

  public connect() {
    if (this.messages$) {
      return;
    }

    this.messages$ = this.socket$.pipe(
     
      switchMap((getResponses: GetWebSocketResponses<string>) => {
        console.log('websocket opened')
        //this.input$.next('test client send')
        return getResponses(this.input$)
      }),
      
      share(),
    )
    
  }

  public send(message: Message) {
    this.input$.next(JSON.stringify(message));
  }
}
