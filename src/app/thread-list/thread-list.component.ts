import { Component, OnInit } from '@angular/core';
import { MailboxService } from '../mailbox.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-thread-list',
  templateUrl: './thread-list.component.html',
  styleUrls: ['./thread-list.component.scss']
})
export class ThreadListComponent implements OnInit {
  list: any[] = [];


  constructor(
    private route: ActivatedRoute,
    private service: MailboxService,
  ) { }

  ngOnInit() {
    this.route.paramMap.subscribe( params => {
      const label = params.get( 'label' );

      this.list = [];

      if ( label ) {
        this.getByLabel( label );
      } else {
        this.getAll();
      }
    });
  }

  getByLabel( id: string ) {
    this.service.getMessagesByLabel( id ).subscribe( threads => {
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

  getAll() {
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
