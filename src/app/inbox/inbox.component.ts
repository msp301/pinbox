import { Component, OnInit } from '@angular/core';
import { MailboxService } from '../mailbox.service';
import { Bundle } from '../core/bundle.model';
import { Thread } from '../core/thread.model';

@Component({
  selector: 'app-inbox',
  templateUrl: './inbox.component.html',
  styleUrls: ['./inbox.component.scss']
})
export class InboxComponent implements OnInit {
  list: any[] = [];

  constructor(
    private service: MailboxService,
  ) { }

  private dateToName( date: Date ): string {
    let name: string;
    const thisYear = new Date().getFullYear();
    const year = date.getFullYear();

    if ( year === thisYear ) {
      name = date.toLocaleDateString( 'en-gb', { month: 'long' } );
    } else if ( thisYear - year <= 1 ) {
      name = date.toLocaleDateString( 'en-gb', { month: 'long', year: 'numeric' } );
    } else {
      name = 'Earlier';
    }

    return name;
  }

  ngOnInit() {
    this.service.getInbox().subscribe( threads => {
      threads.forEach( thread => {

        if ( thread instanceof Thread ) {
          if ( thread.messages.length > 1 ) {
            console.log( `SKIPPING ${thread.subject}` );
          } else {
            thread.messages.forEach( message => {
              this.list.push( {
                id: message.id,
                month: this.dateToName( thread.newestDate ),
                author: thread.authors.join( ', ' ),
                subject: thread.subject,
              });
            });
          }
        } else if ( thread instanceof Bundle ) {
          const messages = [];
          thread.threads.forEach( bundle => {
            bundle.messages.forEach( message => {
              messages.push( {
                id: message.id,
                month: this.dateToName( bundle.newestDate ),
                author: bundle.authors.join( ', ' ),
                subject: bundle.subject,
              });
            });
          });

          this.list.push( {
            id: thread.id,
            month: this.dateToName( thread.date ),
            threads: messages,
          });
        }
      });
    });
  }

}
