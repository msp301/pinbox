import { Component, Input, EventEmitter, OnChanges, Output } from '@angular/core';
import { MailboxService } from './mailbox.service';
import { Message } from './core/message.model';
import { Label } from './core/label.model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [ MailboxService ]
})

export class AppComponent implements OnChanges {
  title = 'Pinbox';
  public labels: Label[];
  public message: Message;

  @Input() messageId: string;
  @Output() messageOpened = new EventEmitter();

  constructor(
    private mailbox: MailboxService,
  ) {
    mailbox.getLabels().subscribe( value => this.labels = value );
  }

  getMessage() {
    this.messageOpened.emit( this.messageId );
    this.message = null;
  }

  ngOnChanges() {
    this.mailbox.getMessage( this.messageId ).subscribe( value => this.message = value );
  }
}

