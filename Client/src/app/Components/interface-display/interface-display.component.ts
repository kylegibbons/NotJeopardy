import { Component, OnInit } from '@angular/core';
import { GameService } from 'src/app/services/game.service';
import { Game } from 'src/app/models/game';

@Component({
  selector: 'app-interface-display',
  templateUrl: './interface-display.component.html',
  styleUrls: ['./interface-display.component.scss']
})
export class InterfaceDisplayComponent implements OnInit {

  game: Game; 

  constructor(private gameService: GameService) { }

  ngOnInit() {

    this.gameService.game$.subscribe(game => {
      this.game = game;
    });

  }

}
