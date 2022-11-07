import { Injectable } from '@angular/core';
import { Adapter } from './adapter';
import { Thread, ThreadAdapter } from './thread.model';

export class Bundle {
  constructor(
    public id: string,
    public date: Date,
    public threads: Thread[],
  ) {}
}

@Injectable({
  providedIn: 'root'
})
export class BundleAdapter implements Adapter<Bundle> {
  adapt( item: any ): Bundle {
    return new Bundle(
      item.id,
      new Date( item.date * 1000 ),
      item.threads.map( thread => new ThreadAdapter().adapt( thread ) ),
    );
  }
}
