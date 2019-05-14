import { Component, OnInit } from '@angular/core';

import { Observable, Subject } from 'rxjs';

import { GameService } from '../../services/game.service';
import { Game } from  '../../models/game';


@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.scss']
})
export class BoardComponent implements OnInit {
  game:Game;
  round: number = 1;
  mode: string = "display";
  //mode: string = "scorekeeper";

  constructor(private gameService: GameService) { }

  ngOnInit() {
    /*this.gameService.game.subscribe(game => {
      this.game = game;
    });*/
  }

  clueClick(categoryNumber:string, clueNumber: string) {
    if(this.mode === "scorekeeper") {
      console.log(categoryNumber + " : " + clueNumber);
    }
  }

}
