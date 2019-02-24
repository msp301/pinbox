import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import {
  MatButtonModule,
  MatExpansionModule,
  MatInputModule,
  MatListModule,
  MatSlideToggleModule,
  MatSidenavModule,
  MatToolbarModule,
  MatIconModule,
  MatCardModule
} from '@angular/material';
import { MailboxService } from './mailbox.service';
import { GoogleSigninComponent } from './google-signin/google-signin.component';

const appRoutes: Routes = [
];

@NgModule({
  declarations: [
    AppComponent,
    GoogleSigninComponent
  ],
  imports: [
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true } // <-- debugging purposes only
    ),
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MatButtonModule,
    MatCardModule,
    MatExpansionModule,
    MatInputModule,
    MatListModule,
    MatSlideToggleModule,
    MatSidenavModule,
    MatToolbarModule,
    MatIconModule,
  ],
  providers: [
    MailboxService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
