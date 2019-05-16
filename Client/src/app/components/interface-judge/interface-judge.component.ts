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

  constructor(private gameService: GameService) { }

  ngOnInit() {

    this.gameService.game$.subscribe(game => {
      this.game = game;
    });

  }
}
