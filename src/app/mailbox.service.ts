import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class MailboxService {

  constructor() { }

  getLabels(): EmailLabel[] {
    return [
        {
            id: 1,
            name: 'Test'
        },
        {
            id: 2,
            name: 'Thing'
        },
        {
            id: 3,
            name: 'Another thing here'
        },
        {
            id: 4,
            name: 'Boo Boo Boo'
        }
    ];
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
