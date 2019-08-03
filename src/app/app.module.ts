import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import {
  MatButtonModule,
  MatExpansionModule,
  MatInputModule,
  MatListModule,
  MatSlideToggleModule,
  MatSidenavModule,
  MatToolbarModule,
  MatIconModule,
  MatCardModule,
  MatMenuModule
} from '@angular/material';
import { QuillModule } from 'ngx-quill';
import { MailboxService } from './mailbox.service';
import { MessageComponent } from './message/message.component';
import { ThreadListComponent } from './thread-list/thread-list.component';
import { InboxComponent } from './inbox/inbox.component';
import { LabelResultsComponent } from './label-results/label-results.component';
import { ThreadComponent } from './thread/thread.component';
import { ServiceWorkerModule } from '@angular/service-worker';
import { environment } from '../environments/environment';
import { ComposerComponent } from './composer/composer.component';

const appRoutes: Routes = [
  { path: '', redirectTo: 'inbox', pathMatch: 'full' },
  { path: 'compose', component: ComposerComponent, outlet: 'right' },
  { path: 'inbox', component: InboxComponent },
  { path: 'messages/:label', component: LabelResultsComponent },
];

@NgModule({
  declarations: [
    AppComponent,
    MessageComponent,
    ThreadListComponent,
    InboxComponent,
    LabelResultsComponent,
    ThreadComponent,
    ComposerComponent
  ],
  imports: [
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true } // <-- debugging purposes only
    ),
    BrowserModule,
    BrowserAnimationsModule,
    FontAwesomeModule,
    HttpClientModule,
    MatButtonModule,
    MatCardModule,
    MatExpansionModule,
    MatInputModule,
    MatListModule,
    MatMenuModule,
    MatSlideToggleModule,
    MatSidenavModule,
    MatToolbarModule,
    MatIconModule,
    ServiceWorkerModule.register('ngsw-worker.js', { enabled: environment.production }),
    QuillModule.forRoot()
  ],
  providers: [
    MailboxService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
