import { Injectable } from '@angular/core';
import { Adapter } from './adapter';
import { Message } from './message.model';

export class Thread {
  constructor(
    public id: string,
    public subject: string,
    public newestDate: Date,
    public oldestDate: Date,
    public authors: string[],
    public messages: Message[],
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
      item.messages,
    );
  }
}
