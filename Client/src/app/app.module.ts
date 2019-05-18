import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatButtonModule} from '@angular/material';


import { InterfaceJudgeComponent } from './Components/interface-judge/interface-judge.component';
import { InterfaceHostComponent } from './Components/interface-host/interface-host.component';
import { InterfaceContestantComponent } from './Components/interface-contestant/interface-contestant.component';
import { BoardComponent } from './Components/board/board.component';

import { GameService } from './services/game.service';
import { WebsocketService } from './services/websocket.service';
import { InterfaceDisplayComponent } from './Components/interface-display/interface-display.component';
import { ClueComponent } from './Components/clue/clue.component';
import { ClueDeterminationComponent } from './Components/clue-determination/clue-determination.component';
import { ClueEnableBuzzerComponent } from './Components/clue-enable-buzzer/clue-enable-buzzer.component';
import { ClueSelectContestantComponent } from './Components/clue-select-contestant/clue-select-contestant.component';

@NgModule({
  declarations: [
    AppComponent,
    InterfaceJudgeComponent,
    InterfaceHostComponent,
    InterfaceContestantComponent,
    BoardComponent,
    InterfaceDisplayComponent,
    ClueComponent,
    ClueDeterminationComponent,
    ClueEnableBuzzerComponent,
    ClueSelectContestantComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatButtonModule
  ],
  providers: [
    GameService,
    WebsocketService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
