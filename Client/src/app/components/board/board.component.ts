import { Component, OnInit, Input } from '@angular/core';

import { Observable, Subject } from 'rxjs';

import { GameService } from '../../services/game.service';
import { Game } from  '../../models/game';


@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.scss']
})
export class BoardComponent implements OnInit {
  @Input() mode: string;
  
  game:Game;
  round: number = 1;
  
  //mode: string = "scorekeeper";

  constructor(private gameService: GameService) {
    if (!this.mode) {
      this.mode = "display";
    }
   }

  ngOnInit() {
    console.log("BOARD MODE:" + this.mode);

    this.gameService.game$.subscribe(game => {
      console.log(game);
      this.game = game;
    });

    this.gameService.activeClue$.subscribe(clueSelect => {
      console.log('ActiveClue: ');
      console.log(clueSelect);
      //this.game = game;
    });
  }

  clueClick(categoryNumber:string, clueNumber: string) {
    if(this.mode === "judge") {
      this.gameService.SelectClue(categoryNumber, clueNumber);
      console.log(categoryNumber + " : " + clueNumber);
    }
  }

}
