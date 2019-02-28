import { Injectable } from '@angular/core';
import { Adapter } from './adapter';

export class Message {
  constructor(
    public id: string,
    public date: Date,
    public subject: string,
    public author: string,
    public content: string,
  ) {}
}

@Injectable({
  providedIn: 'root'
})
export class MessageAdapter implements Adapter<Message> {
  adapt( item: any ): Message {
    return new Message(
      item.id,
      new Date( item.date * 1000 ),
      item.subject,
      item.author,
      item.content,
    );
  }
}
