import { Component, Input, EventEmitter, OnChanges, Output } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { MailboxService } from './mailbox.service';
import { Message } from './core/message.model';
import { Label } from './core/label.model';
import { Observable } from 'rxjs';
import { menuOpened, menuClosed } from './app.actions';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [ MailboxService ]
})

export class AppComponent implements OnChanges {
  title = 'Pinbox';
  public labels: Label[];
  public message: Message;

  @Input() messageId: string;
  @Output() messageOpened = new EventEmitter();

  menuOpen$: Observable<boolean>;

  constructor(
    private store: Store<{ menuOpen: boolean }>,
    private mailbox: MailboxService,
  ) {
    this.menuOpen$ = store.pipe( select( 'menuOpen' ) );
    mailbox.getLabels().subscribe( value => this.labels = value );
  }

  menuOpened() { this.store.dispatch( menuOpened() ); }
  menuClosed() { this.store.dispatch( menuClosed() ); }

  getMessage() {
    this.messageOpened.emit( this.messageId );
    this.message = null;
  }

  ngOnChanges() {
    this.mailbox.getMessage( this.messageId ).subscribe( value => this.message = value );
  }
}

