import { Injectable } from '@angular/core';
import { Actions, Effect, createEffect, ofType } from '@ngrx/effects';
import { MailboxService } from '../mailbox.service';
import { of } from 'rxjs';
import { mergeMap, map, catchError, switchMap } from 'rxjs/operators';

@Injectable()
export class LabelEffects {

  constructor(
    private actions$: Actions,
    private mailbox: MailboxService,
  ) {}

  @Effect()
  loadLabels$ = createEffect( () =>
    this.actions$.pipe(
      ofType( 'Load Labels' ),
      switchMap( () => this.mailbox.getLabels().pipe(
          map( labels => ( { type: 'Load Labels Success', payload: labels } ) ),
          catchError( () => ( of( { type: 'Load Labels Error' } ) ) )
        )
      )
    )
  );
}
