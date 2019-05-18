import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ClueSelectContestantComponent } from './clue-select-contestant.component';

describe('ClueSelectContestantComponent', () => {
  let component: ClueSelectContestantComponent;
  let fixture: ComponentFixture<ClueSelectContestantComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ClueSelectContestantComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClueSelectContestantComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
