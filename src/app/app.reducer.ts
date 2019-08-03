import { createReducer, on } from '@ngrx/store';
import { menuOpened, menuClosed } from './app.actions';

export const menuOpen = true;

export const menuReducer = createReducer(
    menuOpen,
    on( menuOpened, state => true ),
    on( menuClosed, state => false ),
);