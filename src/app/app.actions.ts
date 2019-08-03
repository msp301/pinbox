import { createAction } from '@ngrx/store';

export const menuOpened = createAction( 'Menu Opened' );
export const menuClosed = createAction( 'Menu Closed' );