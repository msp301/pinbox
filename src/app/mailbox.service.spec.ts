import { TestBed } from '@angular/core/testing';

import { MailboxService } from './mailbox.service';

describe('MailboxService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: MailboxService = TestBed.get(MailboxService);
    expect(service).toBeTruthy();
  });
});
