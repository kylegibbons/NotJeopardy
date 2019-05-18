import { Component, OnInit, Output, EventEmitter, Input } from '@angular/core';

@Component({
  selector: 'app-clue-determination',
  templateUrl: './clue-determination.component.html',
  styleUrls: ['./clue-determination.component.scss']
})
export class ClueDeterminationComponent implements OnInit {
  @Input() dailyDouble: boolean;
  @Output() correct = new EventEmitter();
  @Output() wrong = new EventEmitter();

  constructor() { }

  ngOnInit() {
  }

  correctClicked() {
    this.correct.emit();
    console.log("correct clicked")
  }

  wrongClicked() {
    this.wrong.emit();
    console.log("wrong clicked")
  }



}
