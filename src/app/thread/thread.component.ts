import { Component, Input, OnInit } from '@angular/core';
import { MessageComponent } from '../message/message.component';
import { Message } from '@angular/compiler/src/i18n/i18n_ast';

@Component({
  selector: 'app-thread',
  templateUrl: './thread.component.html',
  styleUrls: ['./thread.component.scss']
})
export class ThreadComponent extends MessageComponent implements OnInit {
  @Input() authors: string[];
  @Input() messages: Message[];

  ngOnInit() {
    this.toggle();
  }
}
