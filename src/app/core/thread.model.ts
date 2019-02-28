import { Injectable } from '@angular/core';
import { Adapter } from './adapter';

export class Thread {
  constructor(
    public id: string,
    public subject: string,
    public newestDate: Date,
    public oldestDate: Date,
    public authors: string[],
  ) {}
}

@Injectable({
  providedIn: 'root'
})
export class ThreadAdapter implements Adapter<Thread> {
  adapt( item: any ): Thread {
    return new Thread(
      item.id,
      item.subject,
      new Date( item.newestDate * 1000 ),
      new Date( item.oldestDate * 1000 ),
      item.authors,
    );
  }
}
