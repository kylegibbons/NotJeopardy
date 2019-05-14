import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { InterfaceJudgeComponent } from './Components/interface-judge/interface-judge.component';
import { InterfaceHostComponent } from './Components/interface-host/interface-host.component';
import { InterfaceContestantComponent } from './Components/interface-contestant/interface-contestant.component';
import { BoardComponent } from './Components/board/board.component';

import { GameService } from './services/game.service';
import { WebsocketService } from './services/websocket.service';
import { InterfaceDisplayComponent } from './Components/interface-display/interface-display.component';
import { ClueComponent } from './Components/clue/clue.component';

@NgModule({
  declarations: [
    AppComponent,
    InterfaceJudgeComponent,
    InterfaceHostComponent,
    InterfaceContestantComponent,
    BoardComponent,
    InterfaceDisplayComponent,
    ClueComponent
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
