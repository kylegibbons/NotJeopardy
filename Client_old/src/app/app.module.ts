import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { BoardComponent } from './components/board/board.component';
import { InterfaceContestantComponent } from './Components/interface-contestant/interface-contestant.component';
import { InterfaceHostComponent } from './Components/interface-host/interface-host.component';
import { InterfaceJudgeComponent } from './Components/interface-judge/interface-judge.component';
import { AppRoutingModule } from './/app-routing.module';

import { GameService } from './services/game.service'
import { WebsocketService } from './services/websocket.service';
import { InterfaceDisplayComponent } from './components/interface-display/interface-display.component'


@NgModule({
  declarations: [
    AppComponent,
    BoardComponent,
    InterfaceContestantComponent,
    InterfaceHostComponent,
    InterfaceJudgeComponent,
    InterfaceDisplayComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [
    GameService,
    WebsocketService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
