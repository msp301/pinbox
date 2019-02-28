import { Component, Input, EventEmitter, OnChanges, Output } from '@angular/core';
import { EmailLabel, MailboxService } from './mailbox.service';
import { ActivatedRoute } from '@angular/router';
import { Thread } from './core/thread.model';
import { Message } from './core/message.model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [ MailboxService ]
})

export class AppComponent implements OnChanges {
  title = 'Pinbox';
  public labels: EmailLabel[];
  public emails: Thread[];
  public message: Message;

  @Input() messageId: string;
  @Output() messageOpened = new EventEmitter();

  constructor(
    private route: ActivatedRoute,
    private mailbox: MailboxService,
  ) {
    mailbox.getLabels().subscribe( value => this.labels = value );
    mailbox.getMessages().subscribe( value => this.emails = value );

    // this.route.params.subscribe( params => this.getMessage( params.id ) );
  }

  getMessage() {
    this.messageOpened.emit( this.messageId );
    this.message = null;
  }

  ngOnChanges() {
    this.mailbox.getMessage( this.messageId ).subscribe( value => this.message = value );
  }
}

