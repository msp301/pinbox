import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ThreadListComponent } from './thread-list.component';

describe('ThreadListComponent', () => {
  let component: ThreadListComponent;
  let fixture: ComponentFixture<ThreadListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ThreadListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ThreadListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
