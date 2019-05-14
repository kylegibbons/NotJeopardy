import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InterfaceHostComponent } from './interface-host.component';

describe('InterfaceHostComponent', () => {
  let component: InterfaceHostComponent;
  let fixture: ComponentFixture<InterfaceHostComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InterfaceHostComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InterfaceHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
