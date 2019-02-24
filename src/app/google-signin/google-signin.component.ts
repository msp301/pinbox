import { Component, ElementRef, AfterViewInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-google-signin',
  templateUrl: './google-signin.component.html',
  styleUrls: ['./google-signin.component.scss']
})
export class GoogleSigninComponent implements AfterViewInit {
  private clientID = '411957897094-h6ofim4lbqrugld28e7dj0jfb60potjv.apps.googleusercontent.com';
  private scopes: string = [
    'email'
  ].join( ' ' );

  public auth2: gapi.auth2.GoogleAuth;
  public signedIn = false;

  constructor(
    private element: ElementRef,
    private http: HttpClient,
  ) { }

  ngAfterViewInit() {
    this.googleInit();
  }

  public googleInit() {
    gapi.load( 'auth2', () => {
      this.auth2 = gapi.auth2.init({
        client_id: this.clientID,
        scope: this.scopes,
      });
    });
  }

  public signinHandler() {
      this.auth2.grantOfflineAccess().then( ( result ) => {
        if ( result.code ) {
          this.signedIn = true;
          const payload = { code: result.code };
          this.http.post( 'localhost:8000/authorize', payload );
        } else {
          this.signedIn = false;
          // Failed to grant access.
          // TODO: Trigger error back to user.
        }
      });
  }
}
