import { Component, OnInit, Input } from '@angular/core';
import { MailboxService } from '../mailbox.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-thread-list',
  templateUrl: './thread-list.component.html',
  styleUrls: ['./thread-list.component.scss']
})
export class ThreadListComponent implements OnInit {
  @Input() threads: any[];
  list: any[] = [];

  constructor(
    private route: ActivatedRoute,
    private service: MailboxService,
  ) { }

  ngOnInit() {
    // TODO
    // Split this out into multiple components ???
    // Common behaviour is to display a thread list only.
    // Source of thread list differs.
    if ( this.threads ) {
      this.list = this.threads;
    } else {
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
  }

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

  getByLabel( id: string ) {
    this.service.getMessagesByLabel( id ).subscribe( threads => {
      threads.forEach( thread => {
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
              month: this.dateToName( thread.newestDate ),
              author: thread.authors.join( ', ' ),
              subject: thread.subject,
            });
          });
        }
      });
    });
  }
}
