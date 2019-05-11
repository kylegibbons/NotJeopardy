import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InterfaceJudgeComponent } from './interface-judge.component';

describe('InterfaceJudgeComponent', () => {
  let component: InterfaceJudgeComponent;
  let fixture: ComponentFixture<InterfaceJudgeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InterfaceJudgeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InterfaceJudgeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
