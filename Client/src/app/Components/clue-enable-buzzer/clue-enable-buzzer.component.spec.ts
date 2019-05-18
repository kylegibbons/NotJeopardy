import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ClueEnableBuzzerComponent } from './clue-enable-buzzer.component';

describe('ClueEnableBuzzerComponent', () => {
  let component: ClueEnableBuzzerComponent;
  let fixture: ComponentFixture<ClueEnableBuzzerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ClueEnableBuzzerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClueEnableBuzzerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
