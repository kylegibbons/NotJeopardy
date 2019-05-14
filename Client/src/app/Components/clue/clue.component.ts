import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-clue',
  templateUrl: './clue.component.html',
  styleUrls: ['./clue.component.scss']
})
export class ClueComponent implements OnInit {
  @Input() category: string;
  @Input() clue: string;

  constructor() { 
    if (!this.category) {
      this.category = "[CATEGORY]"
    }

    if (!this.clue) {
      this.clue = "[Clue Text]"
    }
  }

  ngOnInit() {
  }

}
