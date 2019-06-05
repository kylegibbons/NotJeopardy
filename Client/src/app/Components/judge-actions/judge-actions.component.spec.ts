import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { JudgeActionsComponent } from './judge-actions.component';

describe('JudgeActionsComponent', () => {
  let component: JudgeActionsComponent;
  let fixture: ComponentFixture<JudgeActionsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ JudgeActionsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(JudgeActionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
