import { Component, OnInit } from '@angular/core';
import { ClueSelect, Game } from 'src/app/models/game';
import { GameService } from 'src/app/services/game.service';

@Component({
  selector: 'app-interface-judge',
  templateUrl: './interface-judge.component.html',
  styleUrls: ['./interface-judge.component.scss']
})
export class InterfaceJudgeComponent implements OnInit {

  game: Game;

  constructor(public gameService: GameService) { }

  ngOnInit() {

    this.gameService.game$.subscribe(game => {
      this.game = game;
    });

  }

  clueClicked(info) {
    console.log("Clue Clicked:");
    console.log(info);

    this.gameService.SelectClue(info.categoryNumber, info.clueNumber);
  }

  enableBuzzers() {
    this.gameService.EnableBuzzers();
  }

  resetBuzzers() {
    this.gameService.ResetBuzzers();
  }

  selectContestant(contestant: number) {
    console.log("Selecting contestant " + contestant);
    this.gameService.SelectContestant(contestant);
  }

  determinedCorrect() {
    console.log("determined correct")
    this.gameService.DetermineClue(true,0);
  }

  determinedWrong() {
    console.log("determined wrong")
    this.gameService.DetermineClue(false,0);
  }
}
