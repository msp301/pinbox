import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { Label, LabelAdapter } from './core/label.model';
import { Message, MessageAdapter } from './core/message.model';
import { Thread, ThreadAdapter } from './core/thread.model';
import { BundleAdapter, Bundle } from './core/bundle.model';

@Injectable({
  providedIn: 'root'
})
export class MailboxService {

  constructor(
    private http: HttpClient,
    private adapterBundle: BundleAdapter,
    private adapterLabel: LabelAdapter,
    private adapterThread: ThreadAdapter,
    private adapterMessage: MessageAdapter,
  ) { }

  getInbox(): Observable<( Thread | Bundle )[]> {
    return this.http.get( '/api/inbox' ).pipe(
      map( ( data: any[] ) => {
        return data.map(item => {
          if (item.type === 'thread') {
            return this.adapterThread.adapt(item);
          } else {
            return this.adapterBundle.adapt(item);
          }
        });
      }),
    );
  }

  getLabels(): Observable<Label[]> {
    return this.http.get( '/api/labels' ).pipe(
      map( ( data: any[] ) => data.map( item => this.adapterLabel.adapt( item ) ) ),
    );
  }

  getMessage(id: string): Observable<Message> {
    return this.http.get( '/api/messages/' +  encodeURIComponent( id ) ).pipe(
      map( ( data: any ) => this.adapterMessage.adapt( data ) ),
    );
  }

  getMessages(): Observable<Thread[]> {
    return this.http.get( '/api/messages' ).pipe(
      map( ( data: any[] ) => data.map( item => this.adapterThread.adapt( item ) ) ),
    );
  }

  getMessagesByLabel(id: string): Observable<Thread[]> {
    return this.http.get( '/api/messages?label=' + id ).pipe(
      map( ( data: any[] ) => data.map( item => this.adapterThread.adapt( item ) ) ),
    );
  }
}
