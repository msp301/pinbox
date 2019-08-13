import { createAction } from '@ngrx/store';

export const menuOpened = createAction( 'Menu Opened' );
export const menuClosed = createAction( 'Menu Closed' );
export const appLoading = createAction( 'Loading' );

export const loadLabels          = createAction( 'Load Labels' );
export const labelsLoadedSuccess = createAction( 'Load Labels Success', labels => labels );
export const labelsLoadedError   = createAction( 'Load Labels Error' );