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

  getMessages(): Observable<any> {
    return this.http.get( '/api/messages' );
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
