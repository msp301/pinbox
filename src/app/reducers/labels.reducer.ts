import { createReducer, on } from '@ngrx/store';
import { labelsLoadedSuccess } from '../app.actions';

export const labelsReducer = createReducer(
    [],
    on( labelsLoadedSuccess, ( _state, action ) => action.payload ),
);