import { Component, OnInit } from '@angular/core';
import { MailboxService } from '../mailbox.service';
import { Thread } from '../core/thread.model';
import { Message } from '../core/message.model';

@Component({
  selector: 'app-thread-list',
  templateUrl: './thread-list.component.html',
  styleUrls: ['./thread-list.component.scss']
})
export class ThreadListComponent implements OnInit {
  list: any[] = [];

  constructor(
    private service: MailboxService,
  ) { }

  ngOnInit() {
    this.service.getMessages().subscribe( threads => {
      threads.forEach( thread => {
        if ( thread.messages.length > 1 ) {
          console.log( `SKIPPING ${thread.subject}` );
        } else {
          thread.messages.forEach( message => {
            this.list.push( {
              id: message.id,
              author: thread.authors.join( ', ' ),
              subject: thread.subject,
            });
          });
        }
      });
    });
  }

}
