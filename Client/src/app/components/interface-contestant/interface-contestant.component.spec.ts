import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InterfaceContestantComponent } from './interface-contestant.component';

describe('InterfaceContestantComponent', () => {
  let component: InterfaceContestantComponent;
  let fixture: ComponentFixture<InterfaceContestantComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InterfaceContestantComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InterfaceContestantComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
