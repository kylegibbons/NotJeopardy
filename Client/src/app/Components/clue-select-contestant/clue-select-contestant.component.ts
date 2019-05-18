import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { Contestant } from 'src/app/models/game';

@Component({
  selector: 'app-clue-select-contestant',
  templateUrl: './clue-select-contestant.component.html',
  styleUrls: ['./clue-select-contestant.component.scss']
})
export class ClueSelectContestantComponent implements OnInit {
  @Input() contestants: Contestant[];
  @Output() select = new EventEmitter();

  constructor() { }

  ngOnInit() {
    console.log(this.contestants);
  }

  contestantClicked(index: number) {
    this.select.emit(index);
    console.log("Clicked contestant " + index);
  }

}
