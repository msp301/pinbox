import { Component } from '@angular/core';
import { Email, EmailLabel, MailboxService } from './mailbox.service';
import { HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [ MailboxService ]
})

export class AppComponent {
  title = 'Pinbox';
  public labels: Observable<object>;
  public emails: Email[];

  constructor( private mailbox: MailboxService ) {
    mailbox.getLabels().subscribe( value => this.labels = value );
    this.emails = mailbox.getMessages();
  }
}

