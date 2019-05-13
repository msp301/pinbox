import { Component, Input, OnInit } from '@angular/core';
import { MessageComponent } from '../message/message.component';
import { Message } from '@angular/compiler/src/i18n/i18n_ast';

@Component({
  selector: 'app-thread',
  templateUrl: './thread.component.html',
  styleUrls: ['./thread.component.scss']
})
export class ThreadComponent extends MessageComponent {
  @Input() authors: string[];
  @Input() messages: Message[];

  // TODO: Consider whether this should be a subclass of a Message
  toggle() {}
}
