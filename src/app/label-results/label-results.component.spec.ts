import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LabelResultsComponent } from './label-results.component';

describe('LabelResultsComponent', () => {
  let component: LabelResultsComponent;
  let fixture: ComponentFixture<LabelResultsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LabelResultsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LabelResultsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
