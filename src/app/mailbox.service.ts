import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MailboxService {

  constructor( private http: HttpClient ) { }

  getLabels(): Observable<any> {
    return this.http.get( '/api/labels' );
  }

  getMessages(): Email[] {
    return [
        {
            id: 999,
            epoch: 1550930545657,
            recipient: 'Test <me@test.com>',
            sender: 'Jeff <jeff@news.co.uk>',
            subject: 'Latest news in',
            snippet: 'Not much has happened, it is all old news',
        },
        {
            id: 999,
            epoch: 1550930545657,
            recipient: 'Test <me@test.com>',
            sender: 'Jeff <jeff@news.co.uk>',
            subject: 'Latest news in',
            snippet: 'Not much has happened, it is all old news',
        }
    ];
  }

  getMessage() {}
}

export interface EmailLabel {
    id: number;
    name: string;
}

export interface Email {
    id: number;
    epoch: number;
    recipient: string;
    sender: string;
    subject: string;
    snippet: string;
}
