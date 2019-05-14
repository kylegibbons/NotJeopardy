import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { BoardComponent } from './Components/board/board.component';
import { InterfaceJudgeComponent } from './Components/interface-judge/interface-judge.component';
import { InterfaceDisplayComponent } from './Components/interface-display/interface-display.component';
import { ClueComponent } from './Components/clue/clue.component';

const routes: Routes = [
  { path: '', component: BoardComponent },
  { path: 'test', component: ClueComponent },
  { path: 'judge', component: InterfaceJudgeComponent },
  { path: 'display', component: InterfaceDisplayComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
