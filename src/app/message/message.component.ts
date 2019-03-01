import { Component, Input, OnInit } from '@angular/core';
import { MailboxService } from '../mailbox.service';
import { Message } from '../core/message.model';

@Component({
  selector: 'app-message',
  templateUrl: './message.component.html',
  styleUrls: ['./message.component.scss']
})
export class MessageComponent implements OnInit {
  @Input() id: string;
  author: string;
  content: string;
  date: Date;

  constructor(
    private service: MailboxService,
  ) { }

  ngOnInit() {
    this.service.getMessage( this.id ).subscribe(
      ( message: Message ) => {
        this.author = message.author;
        this.date = message.date;
        this.content = atob( message.content );
      }
    );
  }

}
