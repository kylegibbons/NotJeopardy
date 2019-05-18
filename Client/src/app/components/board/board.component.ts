import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

import { ReplaySubject } from 'rxjs';

import { Game } from  '../../models/game';


@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.scss']
})
export class BoardComponent implements OnInit {
  @Input() mode: string;
  @Input() game: Game;
  @Output() clueClicked = new EventEmitter();
  
  categories: string[] = [];
  board: string[][] = [[],[],[],[],[],[]];

  constructor() {
    if (!this.mode) {
      this.mode = "display";
    }
   }

  ngOnInit() {
    console.log("BOARD MODE:" + this.mode);
  }

  ngOnChanges(changes) {
    console.log(this.game)
    if (this.game) {
    this.game.rounds[this.game.round].categories.forEach((category,index1) => {
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
  }
  }

  clueClick(categoryNumber:string, clueNumber: string) {
    this.clueClicked.emit({
      categoryNumber: categoryNumber,
      clueNumber: clueNumber,
    })
  }

}
