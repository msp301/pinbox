import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { Observable } from 'rxjs';

import { LabelEffects } from './label.effects';

describe('LabelEffects', () => {
  let actions$: Observable<any>;
  let effects: LabelEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        LabelEffects,
        provideMockActions(() => actions$)
      ]
    });

    effects = TestBed.get<LabelEffects>(LabelEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
