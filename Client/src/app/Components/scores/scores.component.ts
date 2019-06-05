import { Component, OnInit, Input } from '@angular/core';
import { Contestant } from 'src/app/models/game';

@Component({
  selector: 'app-scores',
  templateUrl: './scores.component.html',
  styleUrls: ['./scores.component.scss']
})
export class ScoresComponent implements OnInit {
  @Input() contestants: Contestant[];

  constructor() { }

  ngOnInit() {
  }

}
