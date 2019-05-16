import { Component, OnInit, Input } from '@angular/core';

import { Observable, Subject } from 'rxjs';

import { GameService } from '../../services/game.service';
import { Game } from  '../../models/game';
import { forEach } from '@angular/router/src/utils/collection';
import { ReturnStatement } from '@angular/compiler';


@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.scss']
})
export class BoardComponent implements OnInit {
  @Input() mode: string;
  
  game:Game;
  categories: string[] = [];
  board: string[][] = [[],[],[],[],[],[]];
  
  //mode: string = "scorekeeper";

  constructor(private gameService: GameService) {
    if (!this.mode) {
      this.mode = "display";
    }
   }

  ngOnInit() {
    console.log("BOARD MODE:" + this.mode);

    this.gameService.game$.subscribe(game => {
      this.game = game;

      game.rounds[game.round].categories.forEach((category,index1) => {
        this.categories[index1] = category.name;
        
        category.clues.forEach((clue, index2) => {
          if (clue.answered) {
            this.board[index1][index2] = ""
            return;
          }
          this.board[index1][index2] = "$" + (200 * (index2 + 1));
        })
      })

      console.log("BOARD:")
      console.log(this.board)

    });
  }

  clueClick(categoryNumber:string, clueNumber: string) {
    if(this.mode === "judge") {
      this.gameService.SelectClue(categoryNumber, clueNumber);
      console.log(categoryNumber + " : " + clueNumber);
    }
  }

}
