import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

import { QueueingSubject } from 'queueing-subject'
import { Observable, Subscription, SubscriptionLike } from 'rxjs'
import { share, switchMap } from 'rxjs/operators'
import makeWebSocketObservable, {
  GetWebSocketResponses,
  normalClosureMessage,
} from 'rxjs-websockets'

@Injectable()
export class WebsocketService {
  //public messages: Observable<String>;
  //public ConnectionStatus: Observable<Number>;

  // this subject queues as necessary to ensure every message is delivered
  private input$ = new QueueingSubject<string>()

  // create the websocket observable, does *not* open the websocket connection
  private socket$ = makeWebSocketObservable<string>('ws://' + environment.base_url + '/ws')

  // the observable produces a value once the websocket has been opened
  private messages$: Observable<string>;

  private messagesSubscription: Subscription;

  public constructor() { 

    this.messages$ = this.socket$.pipe(
     
      switchMap((getResponses: GetWebSocketResponses<string>) => {
        console.log('websocket opened')
        return getResponses(this.input$)
      }),
      share(),
    )

    const messagesSubscription: Subscription = this.messages$.subscribe(
      (message: string) => {
        console.log('received message:', message)
        // respond to server
        this.input$.next('i got your message')
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

  public closeWebsocket() {
    // this also caused the websocket connection to be closed
    this.messagesSubscription.unsubscribe()
  }

  public connect() {
    /*if (this.messages) {
      return;
    }*/

    
  }

  /*public send(message: string):void {
    // If the websocket is not connected then the QueueingSubject will ensure
    // that messages are queued and delivered when the websocket reconnects.
    // A regular Subject can be used to discard messages sent when the websocket
    // is disconnected.
    this.inputStream.next(message)
  }*/
}
