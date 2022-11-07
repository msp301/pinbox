import { Component, OnInit } from '@angular/core';
import { MailboxService } from '../mailbox.service';

@Component({
  selector: 'app-inbox',
  templateUrl: './inbox.component.html',
  styleUrls: ['./inbox.component.scss']
})
export class InboxComponent implements OnInit {
  list: any[] = [];

  constructor(
    private service: MailboxService,
  ) { }

  ngOnInit() {
    this.service.getInbox().subscribe( threads => this.list = threads );
  }
}
