import {
  ActionReducer,
  ActionReducerMap,
  createFeatureSelector,
  createSelector,
  MetaReducer
} from '@ngrx/store';
import { environment } from '../../environments/environment';
import { menuReducer } from '../app.reducer';

export interface State {

}

export const reducers: ActionReducerMap<State> = {
  menuOpen: menuReducer,
};


export const metaReducers: MetaReducer<State>[] = !environment.production ? [] : [];
