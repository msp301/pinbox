import { Injectable } from '@angular/core';
import { Adapter } from './adapter';

export class Label {
  constructor(
    public id: string,
    public name: string,
  ) {}
}

@Injectable({
  providedIn: 'root'
})
export class LabelAdapter implements Adapter<Label> {
  adapt( item: any ): Label {
    return new Label(
      item.id,
      item.name,
    );
  }
}
