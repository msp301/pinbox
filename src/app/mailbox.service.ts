import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { map } from 'rxjs/operators';

import { Thread, ThreadAdapter } from './core/thread.model';
import { Message, MessageAdapter } from './core/message.model';

@Injectable({
  providedIn: 'root'
})
export class MailboxService {

  constructor(
    private http: HttpClient,
    private adapterThread: ThreadAdapter,
    private adapterMessage: MessageAdapter,
  ) { }

  getLabels(): Observable<any> {
    return this.http.get( '/api/labels' );
  }

  getMessage(id: string): Observable<Message> {
    return this.http.get( '/api/messages/' + id ).pipe(
      map( ( data: any ) => data.map( item => this.adapterMessage.adapt(item) ) ),
    );
  }

  getMessages(): Observable<Thread[]> {
    return this.http.get( '/api/messages' ).pipe(
      map( ( data: any[] ) => data.map( item => this.adapterThread.adapt( item ) ) ),
    );
  }
}

export interface EmailLabel {
    id: number;
    name: string;
}

/*export interface Thread {
  id: number;
  authors: string[];
  subject: string;
  newestDate: number;
  oldestDate: number;
  messages: Email[];
}*/

export interface Email {
    id: number;
    epoch: number;
    recipient: string;
    sender: string;
    subject: string;
    snippet: string;
}

export interface EmailContent {
  id: string;
  author: string;
  content: string;
}
