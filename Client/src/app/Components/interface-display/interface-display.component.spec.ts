import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InterfaceDisplayComponent } from './interface-display.component';

describe('InterfaceDisplayComponent', () => {
  let component: InterfaceDisplayComponent;
  let fixture: ComponentFixture<InterfaceDisplayComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InterfaceDisplayComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InterfaceDisplayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
