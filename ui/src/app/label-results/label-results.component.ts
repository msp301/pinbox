import { Component, OnInit } from '@angular/core';
import { MailboxService } from '../mailbox.service';
import { ActivatedRoute } from '@angular/router';
import { Thread } from '../core/thread.model';

@Component({
  selector: 'app-label-results',
  templateUrl: './label-results.component.html',
  styleUrls: ['./label-results.component.scss']
})
export class LabelResultsComponent implements OnInit {
  list: Thread[];

  constructor(
    private route: ActivatedRoute,
    private service: MailboxService,
  ) { }

  ngOnInit() {
    this.route.paramMap.subscribe( params => {
      const label = params.get( 'label' );
      this.list = [];
      this.getByLabel( label );
    });
  }

  getByLabel( id: string ) {
    this.service.getMessagesByLabel( id ).subscribe( threads => this.list = threads );
  }
}
