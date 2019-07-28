import { createReducer, on } from '@ngrx/store';
import { menuIconClicked } from './app.actions';

export const menuOpen = true;

export const menuReducer = createReducer(
    menuOpen,
    on( menuIconClicked, state => !state ),
);