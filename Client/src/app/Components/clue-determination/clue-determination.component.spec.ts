import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ClueDeterminationComponent } from './clue-determination.component';

describe('ClueDeterminationComponent', () => {
  let component: ClueDeterminationComponent;
  let fixture: ComponentFixture<ClueDeterminationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ClueDeterminationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClueDeterminationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
