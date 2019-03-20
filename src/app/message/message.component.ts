import { Component, Input } from '@angular/core';
import { MailboxService } from '../mailbox.service';
import { Message } from '../core/message.model';
import { Label } from '../core/label.model';

@Component({
  selector: 'app-message',
  templateUrl: './message.component.html',
  styleUrls: ['./message.component.scss']
})
export class MessageComponent {
  @Input() id: string;
  @Input() title: string;
  @Input() description: string;

  subject: string;
  author: string;
  content: string;
  date: Date;

  newLabels: Label[];

  constructor(
    private service: MailboxService,
  ) {
    service.getLabels().subscribe( value => this.newLabels = value );
  }

  toggle() {
    if ( ! this.content ) {
      this.service.getMessage( this.id ).subscribe(
        ( message: Message ) => {
          this.author = message.author;
          this.date = message.date;
          this.content = atob( message.content );
        }
      );
    }
  }
}
