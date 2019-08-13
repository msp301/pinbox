import { Component, Input, EventEmitter, OnChanges, Output, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { MailboxService } from './mailbox.service';
import { Message } from './core/message.model';
import { Label } from './core/label.model';
import { Observable } from 'rxjs';
import { menuOpened, menuClosed, appLoading, loadLabels } from './app.actions';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [ MailboxService ]
})

export class AppComponent implements OnInit {
  title = 'Pinbox';
  public message: Message;

  @Input() messageId: string;
  @Output() messageOpened = new EventEmitter();

  labels$: Observable<Label[]>;
  menuOpen$: Observable<boolean>;

  constructor(
    private store: Store<{ menuOpen: boolean }>,
  ) {
    this.menuOpen$ = store.pipe( select( 'menuOpen' ) );
    this.labels$ = store.pipe( select( 'labels' ) );
  }

  menuOpened() { this.store.dispatch( menuOpened() ); }
  menuClosed() { this.store.dispatch( menuClosed() ); }

  getMessage() {
    this.messageOpened.emit( this.messageId );
    this.message = null;
  }

  ngOnInit() {
    this.store.dispatch( loadLabels() );
  }

  /*ngOnChanges() {
    this.mailbox.getMessage( this.messageId ).subscribe( value => this.message = value );
  }*/
}

