import { Component } from '@angular/core';
import { Email, EmailLabel, MailboxService } from './mailbox.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [ MailboxService ]
})

export class AppComponent {
  title = 'Pinbox';
  public labels: EmailLabel[];
  public emails: Email[];

  constructor( private mailbox: MailboxService ) {
    this.labels = mailbox.getLabels();
    this.emails = mailbox.getMessages();
  }
}

