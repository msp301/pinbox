import { Component, Input, OnChanges } from '@angular/core';
import { Thread } from '../core/thread.model';
import { Bundle } from '../core/bundle.model';

interface IBundleListItem {
  id: string;
  month: string;
  threads: Thread[];
}

interface IThreadListItem {
  id: string;
  month: string;
  author: string;
  subject: string;
}

@Component({
  selector: 'app-thread-list',
  templateUrl: './thread-list.component.html',
  styleUrls: ['./thread-list.component.scss']
})
export class ThreadListComponent implements OnChanges {
  @Input() threads: Thread[];
  list: ( IBundleListItem | IThreadListItem )[];

  constructor() { }

  ngOnChanges() {
    this.list = this.threads.map( thread => {
      if ( thread instanceof Bundle ) {
        return {
          id: thread.id,
          month: this.dateToName( thread.date ),
          threads: thread.threads,
        };
      } else {
        if ( thread.messages.length > 1 ) {
          console.log( `SKIPPING ${thread.subject}` );
        } else {
          return {
            id: thread.messages[0].id,
            month: this.dateToName( thread.newestDate ),
            author: thread.authors.join( ', ' ),
            subject: thread.subject,
          };
        }
      }
    });
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
}
