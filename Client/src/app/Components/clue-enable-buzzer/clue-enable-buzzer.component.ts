import { Component, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-clue-enable-buzzer',
  templateUrl: './clue-enable-buzzer.component.html',
  styleUrls: ['./clue-enable-buzzer.component.scss']
})
export class ClueEnableBuzzerComponent implements OnInit {
  @Output() enable = new EventEmitter();
  @Output() reset = new EventEmitter();

  constructor() { }

  ngOnInit() {
  }

  enableClicked() {
    this.enable.emit();
    console.log("enable clicked")
  }

  resetClicked() {
    this.reset.emit();
    console.log("reset clicked")
  }



}
