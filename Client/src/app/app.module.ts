import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { BoardComponent } from './components/board/board.component';
import { InterfaceContestantComponent } from './Components/interface-contestant/interface-contestant.component';
import { InterfaceHostComponent } from './Components/interface-host/interface-host.component';
import { InterfaceJudgeComponent } from './Components/interface-judge/interface-judge.component';
import { AppRoutingModule } from './/app-routing.module';

import { GameService } from './services/game.service'


@NgModule({
  declarations: [
    AppComponent,
    BoardComponent,
    InterfaceContestantComponent,
    InterfaceHostComponent,
    InterfaceJudgeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [
    GameService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
