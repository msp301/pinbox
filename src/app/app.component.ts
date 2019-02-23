import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'Pinbox';

  labels: EmailLabel[] = [
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
    }
  ];

  email: Email = {
    id: 999,
    epoch: 1550930545657,
    recipient: 'Test <me@test.com>',
    sender: 'Jeff <jeff@news.co.uk>',
    subject: 'Latest news in',
    snippet: 'Not much has happened, it is all old news',
  };
}

interface Email {
  id: number;
  epoch: number;
  recipient: string;
  sender: string;
  subject: string;
  snippet: string;
}

interface EmailLabel {
  id: number;
  name: string;
}
